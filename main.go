package main

import (
	stdjson "encoding/json"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
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
	router.GET("/system/status", Status)

	router.Run()

}

func Publish(c *gin.Context) {
	var event struct {
		To   string                 `json:"to" binding:"required"`
		Data map[string]interface{} `json:"data"`
	}

	c.BindJSON(&event)
	data, _ := json.Marshal(event.Data)

	sc.Publish(event.To, []byte(data))
	c.Status(http.StatusAccepted)
}

func Stream(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "http://localhost:8000")

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
			resp := new(map[string]interface{})
			err := stdjson.Unmarshal(msg.Data, resp)

			if err == nil {
				c.SSEvent(msg.Subject, resp)
			}

		case <-pinger.C:
			c.SSEvent("ping", time.Now().String())
		}
		return true
	})

}

func Status(c *gin.Context) {
	c.Status(200)
}

func Index(c *gin.Context) {
	url := location.Get(c)
	url.Path = "/events"
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"eventSourceHost": url.String(),
	})
}
