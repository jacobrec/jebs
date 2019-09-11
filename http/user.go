package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jacobrec/jebs/sql"
)

func getPosts(c *gin.Context) {
	index, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	number, _ := strconv.Atoi(c.DefaultQuery("number", "5"))

	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, sql.GetPosts(index, number))

}

func getPostCount(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, sql.GetPostCount())
}

func getPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))


	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, sql.GetPostByID(id))

}

func getByString(c *gin.Context) {
	searchStr := c.Params.ByName("id")

	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, sql.GetPostsBySearch(searchStr))
}

func getByTag(c *gin.Context) {
	tag := c.Params.ByName("id")

	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, sql.GetPostsByTag(tag))
}
