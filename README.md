## API
    all of these are prepended by blog
    posts.GET("/post/:id", getPost) // this gets the post by the actual ID
    posts.GET("/search/tag/:id", getByTag)
    posts.GET("/search/string/:id", getByString)
    posts.GET("/posts", getPosts) // this supports the query string offset, and number
    //     offset: specifies the number of recent blog posts to skip. Default 0
    //     number: upper limit on the number of blog posts to return. Default 5
    // Eg) to get the most recent 5 posts
    // GET http://this-blog-server.com:8049/blog/posts
    // to get the next 10 most recent posts
    // GET http://this-blog-server.com:8049/blog/posts?offset=5&number=10

    posts.GET("/number, getPostCount) // number of posts


    posts.PUT("/", putPost)
    posts.DELETE("/:id", deletePost)

A post is a json object with the following structure

    type Post struct {
        ID        int      `json:"id"`
        Title     string   `json:"title"`
        Author    string   `json:"author"`
        Post      string   `json:"post"`
        Timestamp uint64   `json:"timestamp"`
        Tags      []string `json:"tags"`
    }

## Emails
Send emails to http://this-blog-server.com:8049/email

    posts.POST("/email, sendEmail) // sends an email

This uses mailgun to send an email. You are required to provide a function to convert the post request to an email

## Setup
1. Install go
2. Get files with `go get github.com/jacobrec/jebs`
3. Install mysql server
4. In mysql create a database and a user and grant that user all privaliges on that database
1. Make the required files described below
2. run with -setup from the src/github.com/jacobrec/spearserver directory
3. run normally from the src/github.com/jacobrec/spearserver directory to start webserver



## Required Files

```
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jacobrec/jebs"
)

func makeEmailFunction(c *gin.Context) (string, string, string, string) {
	var req EmailRequest
	c.BindJSON(&req)
	fmt.Println(req.toString())

	sender := "Sender Name Here <mailgun@mg.example.ca>"
	subject := "Hello, you have 1 new join request"
	body := req.toString()
	recipient := "SPEAR <spear@email.ca>"
    return sender, subject, body, recipient
}

type EmailRequest struct {
	Email    string
	Name     string
	Msg      string
}

func (e EmailRequest) toString() string {
	return "A new member wishes to join spear. Their name is [" + e.Name + "], their email is [" + e.Email + "].\n" + "Their message for you is:\n\n" + e.Msg + "\n"
}


func main() {
    jebs.start_server(makeEmailFunction)
}
```

Config File:
```
export MG_DOMAIN='mg.example.com'
export MG_KEY='MAILGUN-KEY-GOES-HERE'
export PORT='8049'
export GIN_MODE='release'
export DB_USER='user'
export DB_PASSWORD='password'
export DB_DATABASE='database'
export BLOG_USER='author'
export BLOG_PASSWORD='author_password'
```

Run With:
```
go build
source config.env
./foldername # this part can be run like ./foldername -setup # for the first run
```
