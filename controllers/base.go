package controllers

import (
	"TravelSphere/utils"

	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

func (c *BaseController) ServeJSONResponse(resp utils.JSONResponse) {
	c.Data["json"] = resp
	c.ServeJSON()
}
