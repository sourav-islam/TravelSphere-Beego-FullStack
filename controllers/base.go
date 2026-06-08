package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

func (c *BaseController) Prepare() {
	// Provide safe defaults for templates so comparisons do not fail on nil values.
	c.Data["ActiveMenu"] = ""
	c.Data["IsLoggedIn"] = false
}
