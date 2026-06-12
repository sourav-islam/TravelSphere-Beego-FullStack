package services

import (
	"TravelSphere/models"
	"TravelSphere/utils"
	"fmt"
	"strings"
	"sync"

	"github.com/beego/beego/v2/server/web"
)

var (
	countryCache     []models.CountryDTO
	countryCacheLock sync.RWMutex
	cacheLoaded      bool
)

func GetAllCountries() ([]models.CountryDTO, error) {
	countryCacheLock.RLock()
	if cacheLoaded {
		defer countryCacheLock.RUnlock()
		return countryCache, nil
	}
	countryCacheLock.RUnlock()

	// Read API configuration
	baseURL := web.AppConfig.DefaultString("RESTCOUNTRIES_API_BASE_URL", "https://api.restcountries.com/countries/v5")
	apiKey := web.AppConfig.DefaultString("RESTCOUNTRIES_API_KEY", "")

	// Build URL with v5 query parameters
	url := fmt.Sprintf("%s?limit=100&response_fields=names.common,capitals,population,region,subregion,flag.url_png,flag.description,languages,currencies,codes.alpha_2,codes.alpha_3", baseURL)

	// Fetch data with API key authentication
	var apiResponse models.APIResponse
	err := utils.GetJSONWithAuth(url, apiKey, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch countries: %w", err)
	}

	// Transform raw API response to DTOs
	var dtos []models.CountryDTO
	for _, c := range apiResponse.Data.Objects {
		dtos = append(dtos, transformCountry(c))
	}

	// Cache the results
	countryCacheLock.Lock()
	countryCache = dtos
	cacheLoaded = true
	countryCacheLock.Unlock()

	return dtos, nil
}

func SearchCountries(search, region string) ([]models.CountryDTO, error) {
	all, err := GetAllCountries()
	if err != nil {
		return nil, err
	}

	search = strings.ToLower(strings.TrimSpace(search))
	region = strings.ToLower(strings.TrimSpace(region))

	var filtered []models.CountryDTO
	for _, c := range all {
		matchSearch := search == "" ||
			strings.Contains(strings.ToLower(c.Name), search) ||
			strings.Contains(strings.ToLower(c.Capital), search)
		matchRegion := region == "" || region == "all regions" ||
			strings.Contains(strings.ToLower(c.Region), region) ||
			strings.Contains(strings.ToLower(c.Subregion), region)

		if matchSearch && matchRegion {
			filtered = append(filtered, c)
		}
	}
	return filtered, nil
}

func GetCountryBySlug(slug string) (*models.CountryDTO, error) {
	all, err := GetAllCountries()
	if err != nil {
		return nil, err
	}
	for _, c := range all {
		if c.Slug == slug {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("country not found: %s", slug)
}

func GetFeaturedCountries() ([]models.CountryDTO, error) {
	all, err := GetAllCountries()
	if err != nil {
		return nil, err
	}
	featured := []string{"united states", "france", "japan", "australia", "brazil", "bangladesh"}
	var result []models.CountryDTO
	for _, name := range featured {
		for _, c := range all {
			if strings.ToLower(c.Name) == name {
				result = append(result, c)
				break
			}
		}
	}
	return result, nil
}

func transformCountry(c models.Country) models.CountryDTO {
	// Extract capital name from capitals array
	capital := "N/A"
	if len(c.Capitals) > 0 {
		capital = c.Capitals[0].Name
	}

	// Format currencies from array
	var currencyParts []string
	for _, cur := range c.Currencies {
		if cur.Name != "" {
			currencyParts = append(currencyParts, fmt.Sprintf("%s (%s)", cur.Name, cur.Code))
		}
	}
	currencies := strings.Join(currencyParts, ", ")
	if currencies == "" {
		currencies = "N/A"
	}

	// Format languages from array
	var languageParts []string
	for _, lang := range c.Languages {
		if lang.Name != "" {
			languageParts = append(languageParts, lang.Name)
		}
	}
	languages := strings.Join(languageParts, ", ")
	if languages == "" {
		languages = "N/A"
	}

	return models.CountryDTO{
		Slug:       utils.CountryNameToSlug(c.Names.Common),
		Name:       c.Names.Common,
		Capital:    capital,
		Population: utils.FormatPopulation(c.Population),
		Region:     c.Region,
		Subregion:  c.Subregion,
		FlagURL:    c.Flag.URLPng,
		FlagAlt:    c.Flag.Description,
		Languages:  languages,
		Currencies: currencies,
		Cca2:       c.Codes.Alpha2,
	}
}
