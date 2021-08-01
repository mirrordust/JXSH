package web

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/mirrordust/w/m/repo"
)

const MaxPageSize = 100

// RetrievePosts handles url /posts/?scope=normal&tag=1&order=id,desc;title,asc&page=1&pageSize=10
func RetrievePosts(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	var query = "user_id = ?"
	var args = []interface{}{userId}

	scope := strings.ToLower(c.Query("scope"))
	if scope == "" || scope == "normal" {
		query += " AND status & ? = ?"
		args = append(args, repo.PUBLISHED, repo.PUBLISHED)
	} else if scope != "all" {
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid scope"))
		return
	}

	tag := c.Query("tag")
	if tag != "" {
		tagId, err := strconv.ParseUint(tag, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, NewErrorResponse("tag is not number"))
			return
		}
		query += " AND tags & ? = ?"
		args = append(args, tagId, tagId)
	}

	orderQuery := c.DefaultQuery("order", "id,desc")
	orders := orderCond(orderQuery)

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	offset, limit, err := paginationCond(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	cond := repo.Condition{
		Query:  query,
		Args:   args,
		Orders: orders,
		Offset: offset,
		Limit:  limit,
	}

	var posts []repo.Post
	if _, err = repo.FindAll(&posts, cond); err != nil {
		log.Printf("[RetrievePosts] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func RetrievePost(c *gin.Context) {
	var model repo.Post
	_getOne(c, &model)
}

func CreatePost(c *gin.Context) {
	var model repo.Post
	_create(c, &model)
}

func UpdatePost(c *gin.Context) {
	var model repo.Post
	_update(c, &model)
}

func DeletePost(c *gin.Context) {
	var model repo.Post
	_delete(c, &model)
}

func RetrieveTags(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	cond := repo.Condition{
		Query:  "user_id = ?",
		Args:   []interface{}{userId},
		Orders: []interface{}{"name asc"},
	}

	var tags []repo.Tag
	if _, err := repo.FindAll(&tags, cond); err != nil {
		log.Printf("[RetrieveTags] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tags)
}

func RetrieveTag(c *gin.Context) {
	var model repo.Tag
	_getOne(c, &model)
}

func CreateTag(c *gin.Context) {
	var model repo.Tag
	_create(c, &model)
}

func UpdateTag(c *gin.Context) {
	var model repo.Tag
	_update(c, &model)
}

func DeleteTag(c *gin.Context) {
	var model repo.Tag
	_delete(c, &model)
}

// ******************************
// handler utilities

func _getOne(c *gin.Context, model interface{}) {
	userId := c.MustGet("userId").(string)
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	rows, err := repo.FindOne(model, "id = ? AND user_id = ?", id, userId)
	if err != nil {
		log.Printf("[FindOne] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if rows == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, model)
	}
}

func _create(c *gin.Context, model interface{}) {
	userId := c.MustGet("userId").(string)
	err := c.ShouldBindJSON(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	m, ok := model.(repo.Resource)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	m.SetUserId(userId)
	m.CheckAndSetViewPath()

	if err = m.ValidateFields(); err != nil {
		if _, ok = err.(*repo.DBError); ok {
			log.Printf("[Create] check fields error: %v\n", err)
			c.Status(http.StatusInternalServerError)
			return
		} else {
			c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
			return
		}
	}

	if _, err = repo.Create(model); err != nil {
		log.Printf("[Create] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, model)
}

func _update(c *gin.Context, model interface{}) {
	userId := c.MustGet("userId").(string)
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	if err = c.ShouldBindBodyWith(model, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}
	map_ := make(map[string]interface{})
	if err = c.ShouldBindBodyWith(&map_, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	m, ok := model.(repo.Resource)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	m.SetId(id)
	m.SetUserId(userId)
	delete(map_, "id")
	delete(map_, "userId")
	delete(map_, "user_id")
	if _, ok = model.(repo.Tag); ok {
		delete(map_, "code") // Tag Code is not permitted to update
	}
	fields := m.SelectFields(map_)
	if err = m.ValidateFields(fields...); err != nil {
		if _, ok = err.(*repo.DBError); ok {
			log.Printf("[Update] check fields error: %v\n", err)
			c.Status(http.StatusInternalServerError)
			return
		} else {
			c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
			return
		}
	}

	rows, err := repo.UpdateOne(model, fields, model, "user_id = ?", userId)
	if err != nil {
		log.Printf("[Update] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if rows == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusCreated)
	}
}

func _delete(c *gin.Context, model interface{}) {
	userId := c.MustGet("userId").(string)
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	m, ok := model.(repo.Resource)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	m.SetId(id)

	rows, err := repo.Delete(model, "user_id = ?", userId)
	if err != nil {
		log.Printf("[Delete] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if rows == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func orderCond(orders string) []interface{} {
	var ods []interface{}
	for _, o := range strings.Split(orders, ";") {
		ods = append(ods, strings.ReplaceAll(o, ",", " "))
	}
	return ods
}

func paginationCond(page, pageSize string) (offset, limit int, err error) {
	p, err := strconv.Atoi(page)
	if err != nil || p <= 0 {
		return 0, 0, errors.New("invalid page")
	}
	ps, err := strconv.Atoi(pageSize)
	if err != nil || ps <= 0 {
		return 0, 0, errors.New("invalid pageSize")
	}
	if ps > MaxPageSize {
		return 0, 0, errors.New(fmt.Sprintf("the maximum allowed pageSize is %d", MaxPageSize))
	}
	return (p - 1) * ps, ps, nil
}

func checkID(id string) (uint64, error) {
	if id == "" {
		return 0, errors.New("empty id string")
	}
	nid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, errors.New("id is not number")
	}
	return nid, nil
}
