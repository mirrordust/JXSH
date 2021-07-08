package web

import "github.com/gin-gonic/gin"

func Server() *gin.Engine {
	r := gin.Default()
	r.Use(Gate())

	v0 := r.Group("/api/v0")
	v0.Use(Authentication())

	v0.GET("/posts", RetrievePosts)
	v0.GET("/posts/:id", RetrievePost)
	v0.POST("/posts", CreatePost)
	v0.PATCH("/posts/:id", UpdatePost)
	v0.DELETE("/posts/:id", DeletePost)

	v0.GET("/tags", RetrieveTags)
	v0.GET("/tags/:id", RetrieveTag)
	v0.POST("/tags", CreateTag)
	v0.PATCH("/tags/:id", UpdateTag)
	v0.DELETE("/tags/:id", DeleteTag)

	v0.GET("/users", RetrieveUsers)
	v0.GET("/users/:id", RetrieveUser)
	v0.POST("/users", CreateUser)
	v0.PATCH("/users/:id", UpdateUser)
	v0.DELETE("/users/:id", DeleteUser)

	return r
}
