package controllers

type HomeController struct {
	BaseController
}

// Get handles standard browser request rendering for the home layout
func (c *HomeController) Get() {
	c.Layout = "layout/main.tpl"
	c.TplName = "home.tpl"
	c.Data["ActiveMenu"] = "home"
}
