package controllers

import (
	"TravelSphere/services"
	"log"
)

type CountryController struct {
	BaseController
}

func (c *CountryController) Prepare() {
	c.BaseController.Prepare()
	c.Data["ActiveMenu"] = "countries"
}

// SSR: /countries
func (c *CountryController) List() {
	countries, err := services.GetAllCountries()
	if err != nil {
		log.Printf("Error loading countries: %v", err)
		c.Data["Error"] = "Failed to load countries. Please try again."
		countries = nil
	}
	c.Data["Countries"] = countries
	c.TplName = "countries.tpl"
}

// SSR: /countries/:slug
func (c *CountryController) Detail() {
	slug := c.Ctx.Input.Param(":slug")

	country, err := services.GetCountryBySlug(slug)
	if err != nil {
		c.Data["Error"] = "Country not found"
		c.TplName = "404.tpl"
		c.Ctx.Output.SetStatus(404)
		return
	}

	var attractions interface{}
	lat, lon, found := services.GetCapitalCoords(country.Capital)
	if found {
		attrs, aErr := services.GetAttractionsByCountry(lat, lon)
		if aErr != nil {
			log.Printf("Attractions error for %s: %v", country.Name, aErr)
		} else {
			attractions = attrs
		}
	}

	c.Data["Country"] = country
	c.Data["Attractions"] = attractions
	c.TplName = "destination.tpl"
}
