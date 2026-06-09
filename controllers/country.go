package controllers

import (
	"TravelSphere/models"
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
	c.Layout = "layout/main.tpl"
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

	// Use capital name directly — no lat/lon lookup needed
	var attractions []models.AttractionDTO
	if country.Capital != "" && country.Capital != "N/A" {
		attrs, aErr := services.GetAttractionsByCapital(country.Capital)
		if aErr != nil {
			log.Printf("Attractions error for %s (%s): %v", country.Name, country.Capital, aErr)
			// Non-fatal — page still renders, attractions section shows fallback
		} else {
			attractions = attrs
		}
	}

	c.Data["Country"] = country
	c.Data["Attractions"] = attractions
	c.Layout = "layout/main.tpl"
	c.TplName = "destination.tpl"
}
