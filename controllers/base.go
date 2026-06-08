package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
	CurrentUser string
	IsLoggedIn  bool
}

func (c *BaseController) Prepare() {
	// Session-based auth
	user := c.GetSession("username")
	if user != nil {
		c.CurrentUser = user.(string)
		c.IsLoggedIn = true
	}
	c.Data["CurrentUser"] = c.CurrentUser
	c.Data["IsLoggedIn"] = c.IsLoggedIn
	c.Data["ActiveMenu"] = ""
}

func (c *BaseController) RequireLogin() bool {
	if !c.IsLoggedIn {
		c.Redirect("/", 302)
		return false
	}
	return true
}

func (c *BaseController) GetUserID() string {
	if c.IsLoggedIn {
		return c.CurrentUser
	}
	// Demo: use session ID as anonymous user
	sid := c.Ctx.GetCookie("beegosessionID")
	if sid == "" {
		return "anonymous"
	}
	return sid
}
