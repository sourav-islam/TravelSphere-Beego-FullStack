package controllers

import (
	"TravelSphere/services"
	"log"
)

type HomeController struct {
	BaseController
}

// Get handles standard browser request rendering for the home layout
func (c *HomeController) Prepare() {
	c.BaseController.Prepare()
	c.Data["ActiveMenu"] = "home"
}

func (c *HomeController) Get() {
	featured, err := services.GetFeaturedCountries()
	if err != nil {
		log.Printf("Error getting featured countries: %v", err)
		featured = nil
	}

	attractions := services.GetPopularAttractions()

	c.Data["FeaturedCountries"] = featured
	c.Data["PopularAttractions"] = attractions
	c.Layout = "layout/main.tpl"
	c.TplName = "home.tpl"
}

// AJAX: destination search suggestions
func (c *HomeController) Search() {
	query := c.GetString("q")
	if query == "" {
		c.Data["json"] = []interface{}{}
		c.ServeJSON()
		return
	}
	countries, err := services.SearchCountries(query, "")
	if err != nil {
		c.Data["json"] = []interface{}{}
		c.ServeJSON()
		return
	}
	limit := 8
	if len(countries) < limit {
		limit = len(countries)
	}
	c.Data["json"] = countries[:limit]
	c.ServeJSON()
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
