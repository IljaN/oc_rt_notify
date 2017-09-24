package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io"
	"ocevent/router"
	"ocevent/service"
)

type Controller struct {
	Beacon *service.RandomNumbers
	Router *router.Named
}

func NewController(r *router.Named) *Controller {
	return &Controller{Beacon: service.NewBeacon(), Router: r}
}


func (ctrl *Controller) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "test.tmpl", gin.H{
		"eventSourceHost": ctrl.Router.GetUrlForRoute("event_stream", c),
	})
}

func (ctrl *Controller) PublishEvent(c *gin.Context) {
	var event struct {
		Type string `json:"type" binding:"required"`
		To   string `json:"to" binding:"required"`
	}

	c.BindJSON(&event)

	//sc.Publish("events:"+event.To, []byte(event.Type))
	c.Status(http.StatusAccepted)
}


func (ctrl *Controller) EventStream(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-ctrl.Beacon.Numbers:
			c.SSEvent("count", msg)
		}
		return true
	})
}



