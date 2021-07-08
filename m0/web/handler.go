package web

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	uuid "github.com/satori/go.uuid"

	"github.com/mirrordust/splendour/m0/repo"
	"github.com/mirrordust/splendour/m0/util"
)

// ********** middlewares **********

func Gate() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		if strings.ToUpper(c.Request.Method) != "GET" {
			err := authorization(c)
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
	}
}

func authorization(c *gin.Context) error {
	return nil
}

// ********** post handlers **********

// RetrievePosts handle uri /posts/?status=normal&tag=1&order=id,desc;title,asc&page=1&pageSize=10
func RetrievePosts(c *gin.Context) {
	scope := c.Query("scope")
	c1, err := scopeCondition(scope)
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	tag := c.Query("tag")
	c2, err := tagCondition(tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	orders := c.DefaultQuery("order", "id,desc")
	c3, err := orderCondition(orders)
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	c4, err := paginationCondition(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	cond := mergeCondition(c1, c2, c3, c4)
	var posts []repo.Post
	if err = repo.FindAll(&posts, cond); err != nil {
		log.Printf("[RetrievePosts] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func RetrievePost(c *gin.Context) {
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	var post repo.Post
	if err = repo.FindOne(&post, id); err != nil {
		log.Printf("[RetrievePost] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	p := repo.Post{}
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Printf("[CreatePost] request error: %v\n", err)
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	// if there is no viewPath when creating, generate an UUID for the post
	if p.ViewPath == "" {
		p.ViewPath = generateUUID()
	}

	if err := checkPostInsert(&p); err != nil {
		log.Printf("[CreatePost] check fields error: %v\n", err)
		if _, ok := err.(*dBError); ok {
			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		}
		return
	}

	if err := repo.Create(&p); err != nil {
		log.Printf("[CreatePost] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, p)
}

func UpdatePost(c *gin.Context) {
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	p := repo.Post{}
	if err := c.ShouldBindBodyWith(&p, binding.JSON); err != nil {
		log.Printf("[UpdatePost] request error: %v\n", err)
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	pMap := make(map[string]interface{})
	if err := c.ShouldBindBodyWith(&pMap, binding.JSON); err != nil {
		log.Printf("[UpdatePost] request error: %v\n", err)
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}
	delete(pMap, "id")

	fields := p.Select(pMap)
	if err := checkPostInsert(&p, fields...); err != nil {
		log.Printf("[UpdatePost] check fields error: %v\n", err)
		if _, ok := err.(*dBError); ok {
			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		}
		return
	}

	model := repo.Post{}
	model.ID = id
	if err := repo.UpdateOne(&model, fields, p); err != nil {
		log.Printf("[UpdatePost] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, p)
}

func DeletePost(c *gin.Context) {
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	model := repo.Post{}
	model.ID = id
	if err := repo.Delete(&model); err != nil {
		log.Printf("[DeletePost] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

// ********** tag handlers **********

func RetrieveTags(c *gin.Context) {
	cond := repo.Condition{
		Query:  "1=1",
		Orders: []interface{}{"name asc"},
	}

	var tags []repo.Tag
	if err := repo.FindAll(&tags, cond); err != nil {
		log.Printf("[RetrieveTags] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tags)
}

func RetrieveTag(c *gin.Context) {
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	var tag repo.Tag
	if err = repo.FindOne(&tag, id); err != nil {
		log.Printf("[RetrieveTag] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tag)
}

func CreateTag(c *gin.Context) {
	t := repo.Tag{}
	if err := c.ShouldBindJSON(&t); err != nil {
		log.Printf("[CreateTag] request error: %v\n", err)
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	// if there is no viewPath when creating, generate an UUID for the tag
	if t.ViewPath == "" {
		t.ViewPath = generateUUID()
	}

	if err := checkTagInsert(&t); err != nil {
		log.Printf("[CreateTag] check fields error: %v\n", err)
		if _, ok := err.(*dBError); ok {
			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		}
		return
	}

	if err := repo.Create(&t); err != nil {
		log.Printf("[CreateTag] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, t)
}

func UpdateTag(c *gin.Context) {
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	t := repo.Tag{}
	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		log.Printf("[UpdateTag] request error: %v\n", err)
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	tMap := make(map[string]interface{})
	if err := c.ShouldBindBodyWith(&tMap, binding.JSON); err != nil {
		log.Printf("[UpdateTag] request error: %v\n", err)
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}
	delete(tMap, "id")
	delete(tMap, "code") // tag code is not permitted to update

	fields := t.Select(tMap)
	if err := checkTagInsert(&t, fields...); err != nil {
		log.Printf("[UpdateTag] check fields error: %v\n", err)
		if _, ok := err.(*dBError); ok {
			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		}
		return
	}

	model := repo.Tag{}
	model.ID = id
	if err := repo.UpdateOne(&model, fields, t); err != nil {
		log.Printf("[UpdateTag] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, t)
}

func DeleteTag(c *gin.Context) {
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	model := repo.Tag{}
	model.ID = id
	if err := repo.Delete(&model); err != nil {
		log.Printf("[DeleteTag] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

// ********** user handlers **********

func RetrieveUsers(c *gin.Context) {
	cond := repo.Condition{
		Query:  "1=1",
		Orders: []interface{}{"name asc"},
	}

	var users []repo.User
	if err := repo.FindAll(&users, cond); err != nil {
		log.Printf("[RetrieveUsers] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	for i := 0; i < len(users); i++ {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, users)
}

func RetrieveUser(c *gin.Context) {
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	var user repo.User
	if err = repo.FindOne(&user, id); err != nil {
		log.Printf("[RetrieveUser] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	user.Password = ""

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	u := repo.User{}
	if err := c.ShouldBindJSON(&u); err != nil {
		log.Printf("[CreateUser] request error: %v\n", err)
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	u.Password = util.Encrypt(u.Password)

	if err := checkUserInsert(&u); err != nil {
		log.Printf("[CreateUser] check fields error: %v\n", err)
		if _, ok := err.(*dBError); ok {
			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		}
		return
	}

	if err := repo.Create(&u); err != nil {
		log.Printf("[CreateUser] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	u.Password = ""
	c.JSON(http.StatusCreated, u)
}

func UpdateUser(c *gin.Context) {
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	u := repo.User{}
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		log.Printf("[UpdateUser] request error: %v\n", err)
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	uMap := make(map[string]interface{})
	if err := c.ShouldBindBodyWith(&uMap, binding.JSON); err != nil {
		log.Printf("[UpdateUser] request error: %v\n", err)
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}
	delete(uMap, "id")

	fields := u.Select(uMap)
	if err := checkUserInsert(&u, fields...); err != nil {
		log.Printf("[UpdateUser] check fields error: %v\n", err)
		if _, ok := err.(*dBError); ok {
			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		}
		return
	}

	model := repo.User{}
	model.ID = id
	if err := repo.UpdateOne(&model, fields, u); err != nil {
		log.Printf("[UpdateUser] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, u)
}

func DeleteUser(c *gin.Context) {
	id, err := checkID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
		return
	}

	model := repo.User{}
	model.ID = id
	if err := repo.Delete(&model); err != nil {
		log.Printf("[DeleteUser] DB error: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

// ********** utilities **********

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
			existingTags, err := allTags()
			if err != nil {
				return newDBError(err.Error())
			}
			if existingTags|p.Tags != existingTags {
				return errors.New("illegal tags")
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
			tt, err := allTags()
			if err != nil {
				return newDBError(err.Error())
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
			t.Code = validTag // modify input Tag model
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

func checkUserInsert(u *repo.User, fields ...string) error {
	checks := map[string]func() error{
		"Name": func() error {
			if u.Name == "" {
				return errors.New("empty name")
			}
			return nil
		},
		"Password": func() error {
			if u.Password == "" {
				return errors.New("empty password")
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

func allTags() (uint64, error) {
	var tags []repo.Tag
	if err := repo.FindAll(&tags, repo.Condition{}); err != nil {
		return 0, err
	}
	tc := tags[0].Code
	for i := 1; i < len(tags); i++ {
		tc |= tags[i].Code
	}
	return tc, nil
}

func checkID(id string) (uint64, error) {
	if id == "" {
		return 0, errors.New("empty id string")
	}
	nid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, errors.New("id not number")
	}
	return nid, nil
}

// ******************************

func mergeCondition(conds ...repo.Condition) repo.Condition {
	var query = ""
	var args, orders []interface{}
	var offset, limit int
	for _, c := range conds {
		if c.Query != "" {
			if query == "" {
				query = c.Query
			} else {
				query = query + " AND " + c.Query
			}
		}

		args = append(args, c.Args...)

		orders = append(orders, c.Orders...)

		if c.Offset != 0 {
			offset = c.Offset
		}
		if c.Limit != 0 {
			limit = c.Limit
		}
	}
	return repo.Condition{
		Query:  query,
		Args:   args,
		Orders: orders,
		Offset: offset,
		Limit:  limit,
	}
}

func scopeCondition(scope string) (c repo.Condition, e error) {
	switch strings.ToLower(scope) {
	case "":
		fallthrough
	case "normal":
		c = repo.Condition{
			Query: "status & ? = ?",
			Args:  []interface{}{repo.PUBLISHED, repo.PUBLISHED},
		}
		e = nil
	case "all":
		c = repo.Condition{}
		e = nil
	default:
		c = repo.Condition{}
		e = errors.New("scope param error")
	}
	return
}

func tagCondition(tag string) (c repo.Condition, e error) {
	if tag == "" {
		c = repo.Condition{}
		e = nil
		return
	}

	tid, err := strconv.ParseUint(tag, 10, 64)
	if err != nil {
		c = repo.Condition{}
		e = errors.New("tag id not number")
		return
	}

	c = repo.Condition{
		Query: "tags & ? = ?",
		Args:  []interface{}{tid, tid},
	}
	e = nil
	return
}

func orderCondition(orders string) (c repo.Condition, e error) {
	var ods []interface{}
	for _, o := range strings.Split(orders, ";") {
		ods = append(ods, strings.ReplaceAll(o, ",", " "))
	}
	c = repo.Condition{
		Orders: ods,
	}
	e = nil
	return
}

func paginationCondition(page, pageSize string) (c repo.Condition, e error) {
	p, err := strconv.Atoi(page)
	if err != nil {
		c = repo.Condition{}
		e = errors.New("page is not number")
		return
	}

	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		c = repo.Condition{}
		e = errors.New("pageSize is not number")
		return
	}

	offset, limit := paginate(p, ps)
	c = repo.Condition{
		Offset: offset,
		Limit:  limit,
	}
	e = nil
	return
}

func paginate(page, pageSize int) (offset, limit int) {
	if page <= 0 {
		page = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	return (page - 1) * pageSize, pageSize
}

func generateUUID() string {
	return uuid.NewV4().String()
}
