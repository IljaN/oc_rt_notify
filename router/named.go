package router

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/gin-contrib/location"
	"net/url"
)



type Routes = map[Name]NamedRoute
type Name = string
type NamedRoute struct {
	Name Name
	HandlerFunc gin.HandlerFunc
	gin.RouteInfo

}


func C(method string, path string, name string, handler gin.HandlerFunc) NamedRoute {
	return NamedRoute{
		Name: name,
		HandlerFunc: handler,
		RouteInfo: gin.RouteInfo{
			Method: method,
			Path:path,
			Handler: "",
			},
	}


}

type Named struct {
	routes Routes
	engine *gin.Engine
}


func CreateNamedRouter(e *gin.Engine) *Named {
	return &Named{
		engine: e,
		routes: Routes{},
	}
}

func (n *Named) handle(r *NamedRoute, handler gin.HandlerFunc) gin.IRoutes {
	return n.engine.Handle(r.Method, r.Path, handler)
}




func (n *Named) Routes(routes ...NamedRoute) {
	for _, r := range routes {
		n.addRoute(r)
		n.engine.Handle(r.Method, r.Path, r.HandlerFunc)
	}

}

func (n *Named) addRoute(r NamedRoute) {
	n.routes[r.Name] = r
}

func (n *Named) GetRoute(name string) NamedRoute {
	if _,ok := n.routes[name]; !ok {
		panic(fmt.Sprintf("Unknown route, %v", name))
	}

	return n.routes[name]
}

func (n *Named) GetUrlForRoute(name string, c *gin.Context) string {
	r := n.GetRoute(name)
	l := location.Get(c)
	u := url.URL{Scheme: l.Scheme, Host: l.Host, Path: r.Path}
	return u.String()

}




