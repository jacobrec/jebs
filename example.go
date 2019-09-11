package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jacobrec/spearserver"
)

func makeEmailFunction(c *gin.Context) (string, string, string, string) {
	var req EmailRequest
	c.BindJSON(&req)
	fmt.Println(req.toString())

	sender := "Jacob Reckhard <mailgun@mg.reckhard.ca>"
	subject := "Hello, you have 1 new join request"
	body := req.toString()
	recipient := "SPEAR <spearua@ualberta.ca>"
    return sender, subject, body, recipient
}

type EmailRequest struct {
	TeamId   string
	TeamName string
	Email    string
	Name     string
	Msg      string
}

func (e EmailRequest) toString() string {
	return "A new member wishes to join spear. Their name is [" + e.Name + "], their email is [" + e.Email + "].\n" +
		"They wish to join the [" + e.TeamName + "] team. Their message for you is:\n\n" + e.Msg + "\n"
}


func main() {
    spearserver.start_server(makeEmailFunction)
}
