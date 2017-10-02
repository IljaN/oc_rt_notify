package main

import (
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/go-nats"
	"io"
	"net/http"
	"strings"
)

var nc *nats.Conn

func main() {
	router := gin.Default()
	nc, _ = nats.Connect("nats://localhost:4222")

	router.Use(location.New(location.Config{
		Scheme: "http",
		Host:   "localhost:8080",
	}))

	router.LoadHTMLGlob("templates/*")
	router.GET("/", Index)
	router.POST("/events", Publish)
	router.GET("/events", Stream)
	router.Run()

}

func Index(c *gin.Context) {
	url := location.Get(c)
	url.Path = "/events"
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"eventSourceHost": url.String(),
	})
}

func Publish(c *gin.Context) {
	var event struct {
		Type string `json:"type" binding:"required"`
		To   string `json:"to" binding:"required"`
	}

	c.BindJSON(&event)
	nc.Publish(event.To, []byte(event.Type))
	c.Status(http.StatusAccepted)
}

func Stream(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	messages := make(chan *nats.Msg, 15)
	topics := strings.Split(c.DefaultQuery("topics", ""), ";")

	for _, t := range topics {
		nc.Subscribe(t, func(msg *nats.Msg) {
			messages <- msg

		})
	}

	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-messages:
			c.SSEvent(msg.Subject, fmt.Sprintf("%s", msg.Data))
		}
		return true
	})
}
