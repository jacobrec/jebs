package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

/*Begin starts the blog webserver on a specified port*/
func Begin(port string, email func(c *gin.Context) (string, string, string, string)) {

	os.Mkdir(STORAGE, os.ModeDir|0777)
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

	imgs := router.Group("/images")
	{
		imgs.GET(VIEW_ENDPOINT, viewHandler)
		imgs.POST(HANDLE_ENDPOINT, uploadHandler)
		imgs.DELETE(DELETE_ENDPOINT, deleteHandler)
	}

	writer := router.Group("/author")
	{
		router.LoadHTMLFiles(getWritePath()+"/index.html", getWritePath()+"/posts.html", getWritePath()+"/images.html")
		loc := os.Getenv("LOCATION") + ":" + os.Getenv("PORT")
		writer.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"location": loc,
			})
		})
		writer.GET("/posts.html", func(c *gin.Context) {
			c.HTML(http.StatusOK, "posts.html", gin.H{
				"location": loc,
			})
		})
		writer.GET("/images.html", func(c *gin.Context) {
			c.HTML(http.StatusOK, "images.html", gin.H{
				"location": loc,
				"files":    "[\"" + strings.Join(listFiles(), "\", \"") + "\"]",
			})
		})
		writer.Use(authRequired())
		writer.Static("/static", getWritePath())
	}
	fmt.Println("Serving Author out of: ", getWritePath())

	router.Run(port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
