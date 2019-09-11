package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go"
	"time"
    "os"
)


func sendEmailFromPost(email func(*gin.Context) (string, string, string, string)) func(*gin.Context){
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

            c.String(200, "Okay")
        }
        return inner
}
