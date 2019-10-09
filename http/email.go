package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go"
	"os"
	"time"
)

func sendEmailFromPost(email func(*gin.Context) (string, string, string, string)) func(*gin.Context) {
	inner := func(c *gin.Context) {
		fmt.Println("Starting join")
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")

		mg := mailgun.NewMailgun(
			os.Getenv("MG_DOMAIN"),
			os.Getenv("MG_KEY"))

		sender, subject, body, recipient := email(c)
		m := mg.NewMessage(
			sender,
			subject,
			body,
			recipient,
		)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		_, id, err := mg.Send(ctx, m)
		fmt.Println(id, err)

		loc := ""
		fmt.Println(c.Request.Header)
		if len(c.Request.Header["Origin"]) > 0 {
			loc = c.Request.Header["Origin"][0]
		} else if len(c.Request.Header["Referer"]) > 0 {
			loc = c.Request.Header["Referer"][0]
		}
		c.Redirect(302, loc)
	}
	return inner
}
