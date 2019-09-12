package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*Begin starts the blog webserver on a specified port*/
func Begin(port string, email func(c *gin.Context) (string, string, string, string)) {

	router := gin.Default()
	router.Use(CORSMiddleware())
	posts := router.Group("/blog")
	{
		posts.GET("/posts", getPosts)
		posts.GET("/post/:id", getPost)

		posts.GET("/search/tag/:id", getByTag)
		posts.GET("/search/string/:id", getByString)

		posts.GET("/number", getPostCount)

		posts.PUT("/", putPost)
		posts.DELETE("/:id", deletePost)
	}

	if email != nil {
		// See http/email.go for email sending stuff
		router.POST("/email", sendEmailFromPost(email))
		fmt.Println("Using Email Server")
	}

	writer := router.Group("/author")
	writer.Use(authRequired())
	{
		writer.Static("/", getWritePath())
	}
	fmt.Println("Serving Author out of: ", getWritePath())

	router.Run(port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
