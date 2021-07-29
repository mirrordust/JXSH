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
	"github.com/mirrordust/w/m/util"
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
	if err = repo.FindAll(&posts, cond); err != nil {
		log.Printf("[RetrievePosts] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func RetrievePost(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	var post repo.Post
	if err = repo.FindOne(&post, "id = ? AND user_id = ?", id, userId); err != nil {
		log.Printf("[RetrievePost] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	p := repo.Post{}
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	// if there is no viewPath when creating, generate a UUID for the post
	if p.ViewPath == "" {
		p.ViewPath = util.GenerateUUID()
	}

	p.UserId = userId

	if err := checkPostInsert(&p); err != nil {
		if _, ok := err.(*repo.DBError); ok {
			log.Printf("[CreatePost] check fields error: %v\n", err)
			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		}
		return
	}

	if err := repo.Create(&p); err != nil {
		log.Printf("[CreatePost] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, p)
}

func UpdatePost(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	p := repo.Post{}
	if err = c.ShouldBindBodyWith(&p, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	pMap := make(map[string]interface{})
	if err = c.ShouldBindBodyWith(&pMap, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}
	delete(pMap, "id")
	delete(pMap, "userId")
	delete(pMap, "user_id")

	p.UserId = userId

	fields := p.SelectFields(pMap)
	if err = checkPostInsert(&p, fields...); err != nil {
		if _, ok := err.(*repo.DBError); ok {
			log.Printf("[UpdatePost] check fields error: %v\n", err)
			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		}
		return
	}

	model := repo.Post{}
	model.ID = id
	if err = repo.UpdateOne(&model, fields, p, "user_id = ?", userId); err != nil {
		log.Printf("[UpdatePost] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, p)
}

func DeletePost(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	model := repo.Post{}
	model.ID = id
	if err = repo.Delete(&model, "id = ? AND user_id = ?", id, userId); err != nil {
		log.Printf("[DeletePost] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func RetrieveTags(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	cond := repo.Condition{
		Query:  "user_id = ?",
		Args:   []interface{}{userId},
		Orders: []interface{}{"name asc"},
	}

	var tags []repo.Tag
	if err := repo.FindAll(&tags, cond); err != nil {
		log.Printf("[RetrieveTags] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tags)
}

func RetrieveTag(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	var tag repo.Tag
	if err = repo.FindOne(&tag, "id = ? AND user_id = ?", id, userId); err != nil {
		log.Printf("[RetrieveTag] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tag)
}

func CreateTag(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	t := repo.Tag{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	// if there is no viewPath when creating, generate a UUID for the tag
	if t.ViewPath == "" {
		t.ViewPath = util.GenerateUUID()
	}

	t.UserId = userId

	if err := checkTagInsert(&t); err != nil {
		if _, ok := err.(*repo.DBError); ok {
			log.Printf("[CreateTag] check fields error: %v\n", err)
			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		}
		return
	}

	if err := repo.Create(&t); err != nil {
		log.Printf("[CreateTag] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, t)
}

func UpdateTag(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	t := repo.Tag{}
	if err = c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	tMap := make(map[string]interface{})
	if err = c.ShouldBindBodyWith(&tMap, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}
	delete(tMap, "id")
	delete(tMap, "code") // tag code is not permitted to update
	delete(tMap, "userId")
	delete(tMap, "user_id")

	t.UserId = userId

	fields := t.SelectFields(tMap)
	if err = checkTagInsert(&t, fields...); err != nil {
		if _, ok := err.(*repo.DBError); ok {
			log.Printf("[UpdateTag] check fields error: %v\n", err)
			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		}
		return
	}

	model := repo.Tag{}
	model.ID = id
	if err = repo.UpdateOne(&model, fields, t, "user_id = ?", t.UserId); err != nil {
		log.Printf("[UpdateTag] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, t)
}

func DeleteTag(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	model := repo.Tag{}
	model.ID = id
	if err = repo.Delete(&model, "id = ? AND user_id = ?", id, userId); err != nil {
		log.Printf("[DeleteTag] %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

// ******************************
// handler utilities

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

func checkPostInsert(p *repo.Post, fields ...string) error {
	checks := map[string]func() error{
		"Title": func() error {
			if p.Title == "" {
				return errors.New("empty title")
			}
			return nil
		},
		"ViewPath": func() error {
			if p.ViewPath == "" {
				return errors.New("empty viewPath")
			}
			return nil
		},
		"Tags": func() error {
			existingTags, err := allTags(p.UserId)
			if err != nil {
				return err
			}
			if existingTags|p.Tags != existingTags {
				return errors.New("illegal post tags")
			}
			return nil
		},
	}

	if len(fields) == 0 {
		for _, v := range checks {
			if err := v(); err != nil {
				return err
			}
		}
	} else {
		for _, fieldName := range fields {
			if _, present := checks[fieldName]; present {
				if err := checks[fieldName](); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func checkTagInsert(t *repo.Tag, fields ...string) error {
	checks := map[string]func() error{
		"Name": func() error {
			if t.Name == "" {
				return errors.New("empty name")
			}
			return nil
		},
		"ViewPath": func() error {
			if t.ViewPath == "" {
				return errors.New("empty viewPath")
			}
			return nil
		},
		"Code": func() error {
			tt, err := allTags(t.UserId)
			if err != nil {
				return err
			}
			validBits := ^tt
			k, available := 0, false
			for validBits != 0 {
				if validBits&1 == 1 {
					available = true
					break
				} else {
					k++
					validBits >>= 1
				}
			}
			if !available {
				return errors.New("exceed maximum tags number")
			}
			validTag := uint64(1) << k
			// ** here modify the input Tag model
			t.Code = validTag
			return nil
		},
	}

	if len(fields) == 0 {
		for _, v := range checks {
			if err := v(); err != nil {
				return err
			}
		}
	} else {
		for _, fieldName := range fields {
			if _, present := checks[fieldName]; present {
				if err := checks[fieldName](); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func allTags(userId string) (uint64, error) {
	if userId == "" {
		return 0, errors.New("empty userId")
	}
	var tags []repo.Tag
	var cond = repo.Condition{
		Query: "user_id = ?",
		Args:  []interface{}{userId},
	}
	if err := repo.FindAll(&tags, cond); err != nil {
		return 0, err
	}
	if len(tags) == 0 {
		return 0, nil
	}
	tc := tags[0].Code
	for i := 1; i < len(tags); i++ {
		tc |= tags[i].Code
	}
	return tc, nil
}
