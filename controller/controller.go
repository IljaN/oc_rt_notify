package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/go-nats"
	"io"
	"net/http"
	"ocevent/router"
	"strings"
)

type Controller struct {
	Router *router.Named
	Nats   *nats.Conn
}

func NewController(r *router.Named, n *nats.Conn) *Controller {
	return &Controller{Router: r, Nats: n}
}

func (ctrl *Controller) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "test.tmpl", gin.H{
		"eventSourceHost": ctrl.Router.GetUrlForRoute("event_stream", c),
	})
}

func (ctrl *Controller) Publish(c *gin.Context) {
	var event struct {
		Type string `json:"type" binding:"required"`
		To   string `json:"to" binding:"required"`
	}

	c.BindJSON(&event)
	ctrl.Nats.Publish(event.To, []byte(event.Type))
	c.Status(http.StatusAccepted)
}

func (ctrl *Controller) Stream(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	messages := make(chan *nats.Msg, 15)
	topics := strings.Split(c.DefaultQuery("topics", ""), ";")

	for _, t := range topics {
		ctrl.Nats.Subscribe(t, func(msg *nats.Msg) {
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
