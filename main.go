package main

import (
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/go-nats-streaming"
	"io"
	"net/http"
	"time"
)

var sc stan.Conn
var sessionManager *SessionManager

func main() {
	router := gin.Default()

	nc, _ := stan.Connect("test-cluster", "publisher")
	sc = nc
	sessionManager = NewSessionManager()

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
		Data string `json:"data" binding:"required"`
		To   string `json:"to" binding:"required"`
	}

	c.BindJSON(&event)
	sc.Publish(event.To, []byte(event.Data))
	c.Status(http.StatusAccepted)
}

func Stream(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	userId := c.DefaultQuery("userId", "")
	ses := sessionManager.StartSession(userId)
	pinger := time.NewTicker(5 * time.Second)

	defer func() {
		sessionManager.EndSession(ses)
		pinger.Stop()
	}()

	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-ses.Messages:
			c.SSEvent(msg.Subject, fmt.Sprintf("%s", msg.Data))
		case <-pinger.C:
			c.SSEvent("ping", time.Now().String())
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
