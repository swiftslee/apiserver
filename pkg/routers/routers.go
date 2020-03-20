package routers

import (
	"github.com/emicklei/go-restful/v3"

	"github.com/yuswift/apiserver/pkg/controller/user"
)

func routes(ws *restful.WebService) {
	ws.
		Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/users").To(user.Controller.List))
	ws.Route(ws.GET("/users/{name}").To(user.Controller.Get))
	ws.Route(ws.POST("/users").To(user.Controller.Create))
	ws.Route(ws.DELETE("/users/{name}").To(user.Controller.Delete))
	ws.Route(ws.PUT("/users/{name}").To(user.Controller.Update))

}

func Run() {
	ws := new(restful.WebService)
	routes(ws)
	restful.DefaultContainer.Add(ws)
}
