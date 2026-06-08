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

func (c *HomeController) Login() {
	username := c.GetString("username")
	if username == "" {
		username = "guest"
	}
	c.SetSession("username", username)
	c.Redirect("/", 302)
}

func (c *HomeController) Logout() {
	c.DestroySession()
	c.Redirect("/", 302)
}
