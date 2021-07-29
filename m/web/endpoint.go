package web

import "github.com/gin-gonic/gin"

func Server() *gin.Engine {
	r := gin.Default()
	r.Use(Gate())

	g := r.Group("/api/v0", Auth())

	g.GET("/posts", RetrievePosts)
	g.GET("/posts/:id", RetrievePost)
	g.POST("/posts", CreatePost)
	g.PATCH("/posts/:id", UpdatePost)
	g.DELETE("/posts/:id", DeletePost)

	g.GET("/tags", RetrieveTags)
	g.GET("/tags/:id", RetrieveTag)
	g.POST("/tags", CreateTag)
	g.PATCH("/tags/:id", UpdateTag)
	g.DELETE("/tags/:id", DeleteTag)

	return r
}
