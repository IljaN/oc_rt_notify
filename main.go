package main

import (
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/go-nats-streaming"
	"io"
	"net/http"
	"strings"
)

var sc stan.Conn

func main() {
	router := gin.Default()
	nc, _ := stan.Connect("test-cluster", "foo123")
	sc = nc

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

func Publish(c *gin.Context) {
	var event struct {
		Type string `json:"type" binding:"required"`
		To   string `json:"to" binding:"required"`
	}

	c.BindJSON(&event)
	sc.Publish(event.To, []byte(event.Type))
	c.Status(http.StatusAccepted)
}

func Stream(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	messages := make(chan *stan.Msg, 15)
	topics := strings.Split(c.DefaultQuery("topics", ""), ";")

	for _, t := range topics {
		sc.Subscribe(t, func(msg *stan.Msg) {
			messages <- msg

		}, stan.DeliverAllAvailable())
	}

	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-messages:
			c.SSEvent(msg.Subject, fmt.Sprintf("%s", msg.Data))
		}
		return true
	})
}

func Index(c *gin.Context) {
	url := location.Get(c)
	url.Path = "/events"
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"eventSourceHost": url.String(),
	})
}
