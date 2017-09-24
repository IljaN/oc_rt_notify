package main

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	nr "ocevent/router"
	"ocevent/controller"
	//"github.com/nats-io/go-nats"
)

func main() {
	//queue, _ := nats.Connect("nats://localhost:4222")
	router, named := configure()
	ctrl := controller.NewController(named)
	named.Routes(
		nr.C("GET", "/", "index", ctrl.Index),
		nr.C("POST", "/events", "create_event", ctrl.PublishEvent),
		nr.C("GET", "/notifications", "event_stream", ctrl.EventStream),
	)

	router.Run()

}

func configure() (*gin.Engine, *nr.Named) {
	r := gin.Default()
	r.Use(location.New(location.Config{
		Scheme: "http",
		Host:   "localhost:8080",
	}))

	r.LoadHTMLGlob("templates/*")


	return r, nr.CreateNamedRouter(r)
}
