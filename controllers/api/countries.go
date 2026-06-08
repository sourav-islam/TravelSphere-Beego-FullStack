package apicontrollers

import (
	"TravelSphere/services"
	"TravelSphere/utils"

	"github.com/beego/beego/v2/server/web"
)

type CountriesAPIController struct {
	web.Controller
}

// GET /api/countries?search=...&region=...
func (c *CountriesAPIController) List() {
	search := c.GetString("search")
	region := c.GetString("region")

	countries, err := services.SearchCountries(search, region)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = utils.ErrorResponse("Failed to fetch countries")
		c.ServeJSON()
		return
	}

	c.Data["json"] = utils.SuccessResponse(countries)
	c.ServeJSON()
}

// GET /api/countries/:slug
func (c *CountriesAPIController) Detail() {
	slug := c.Ctx.Input.Param(":slug")
	country, err := services.GetCountryBySlug(slug)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = utils.ErrorResponse("Country not found")
		c.ServeJSON()
		return
	}
	c.Data["json"] = utils.SuccessResponse(country)
	c.ServeJSON()
}
