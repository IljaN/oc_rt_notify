package main

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/go-nats"
	"ocevent/controller"
	nr "ocevent/router"
)

func main() {
	router, named := configure()
	n, _ := nats.Connect("nats://localhost:4222")
	ctrl := controller.NewController(named, n)
	named.Routes(
		nr.C("GET", "/", "index", ctrl.Index),
		nr.C("POST", "/events", "create_event", ctrl.Publish),
		nr.C("GET", "/events", "event_stream", ctrl.Stream),
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
