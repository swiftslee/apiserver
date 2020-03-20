package user

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful/v3"

	"github.com/yuswift/apiserver/pkg/models"
)

type controller struct{}

var Controller = controller{}

func (c *controller) Get(req *restful.Request, res *restful.Response) {
	ret, err := models.UserModel.Get(req.PathParameter("name"))
	if err != nil {
		res.WriteErrorString(http.StatusNotFound, "")
	}else{
		res.WriteEntity(ret)
	}
}

func (c *controller) List(req *restful.Request, res *restful.Response) {
	ret, err := models.UserModel.List("")
	if err != nil {
		res.WriteErrorString(http.StatusNotFound, fmt.Sprintf("%s", err))
	}else {
		res.WriteEntity(ret)
	}
}

func (c *controller) Create(req *restful.Request, res *restful.Response) {
	user := &models.User{}
	if err := req.ReadEntity(user); err != nil {
		res.WriteErrorString(http.StatusBadRequest, "invalid request body")
		return
	}

	ret, err := models.UserModel.Create(user)
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, fmt.Sprintf("%s", err))

	}else{
		res.WriteEntity(ret)
	}
}

func (c *controller) Delete(req *restful.Request, res *restful.Response) {
	err := models.UserModel.Delete(req.PathParameter("name"))
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}else{
		res.WriteEntity("")
	}
}

func (c *controller) Update(req *restful.Request, res *restful.Response) {
	user := &models.User{}
	if err := req.ReadEntity(user); err != nil {
		res.WriteErrorString(http.StatusBadRequest, "invalid request body")
		return
	}

	name := req.PathParameter("name")
	err := models.UserModel.Update(name, user)
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, "")

	}else {
		res.WriteEntity(user)
	}

}
