package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/go-nats"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	notifications := make(chan int, 20)
	go func() {
		for {
			notifications <- rand.Int()
			time.Sleep(1 * time.Second)

		}

	}()
	sc, _ := nats.Connect("nats://localhost:4222")

	router := gin.Default()
	router.POST("/events", func(c *gin.Context) {
		var event struct {
			Type string `json:"type" binding:"required"`
			To   string `json:"to" binding:"required"`
		}

		c.BindJSON(&event)
		sc.Publish("events:"+event.To, []byte(event.Type))
		c.Status(http.StatusAccepted)
	})

	router.GET("/notifications", func(c *gin.Context) {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		c.Stream(func(w io.Writer) bool {
			select {
			case msg := <-notifications:
				c.SSEvent("count", msg)
			}
			return true
		})
	})

	router.GET("/", func(c *gin.Context) {
		html, _ := ioutil.ReadFile("test.html")
		c.Writer.Write([]byte(html))
		c.Status(http.StatusOK)
	})

	router.Run(":8080")
}
