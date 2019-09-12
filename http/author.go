package http

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jacobrec/jebs/blog"
	"github.com/jacobrec/jebs/sql"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func authRequired() func(c *gin.Context) {
	return func(c *gin.Context) {

		if strings.Index(c.GetHeader("Authorization"), "Basic") == 0 {
			userPass := strings.SplitAfter(c.GetHeader("Authorization"), "Basic ")[1]

			// decode username and password
			bytes, _ := base64.StdEncoding.DecodeString(userPass)

			decoded := strings.Split(string(bytes), ":")

			if os.Getenv("BLOG_USER") == decoded[0] && os.Getenv("BLOG_PASSWORD") == decoded[1] {
				c.Next()
				return
			}
		}

		c.Writer.Header().Add("WWW-Authenticate", "Basic")
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Next()
	}

}

func deletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	sql.DeletePost(id)

}
func putPost(c *gin.Context) {
	var post blog.Post
	c.BindJSON(&post)
	fmt.Println(post.ID)
	if post.ID == -1 {
		sql.AddPost(post)
	} else {
		sql.EditPost(post)
	}
	fmt.Println(post)
}
