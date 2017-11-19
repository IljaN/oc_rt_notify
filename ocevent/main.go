package main

import (
	stdjson "encoding/json"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"github.com/nats-io/go-nats-streaming"
	"io"
	"net/http"
	"sync"
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

	router.POST("/events", Authenticate(Publishing), publishEvent)
	router.GET("/events", Authenticate(Subscribing), subscribe)
	router.GET("/system/status", getSystemStatus)
	router.OPTIONS("/events", preFlight)

	router.Run()
}

func publishEvent(c *gin.Context) {
	var event struct {
		To   []string               `json:"to" binding:"required"`
		Data map[string]interface{} `json:"data"`
	}

	c.BindJSON(&event)
	data, _ := json.Marshal(event.Data)

	var wg sync.WaitGroup
	wg.Add(len(event.To))

	for i := range event.To {
		go func(i int) {
			defer wg.Done()
			sc.Publish(event.To[i], []byte(data))
		}(i)
	}

	wg.Wait()
	c.Status(http.StatusAccepted)
}

func subscribe(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, authorization")

	ses := sessionManager.StartSession(c.MustGet("username").(string))
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
			c.SSEvent("comment", time.Now().String())
		}
		return true
	})

}

func getSystemStatus(c *gin.Context) {
	c.Status(200)
}

func preFlight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, authorization")
	c.Status(http.StatusOK)
}
