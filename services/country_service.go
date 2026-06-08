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

	var raw []models.Country
	baseURL := web.AppConfig.DefaultString("restcountries_api_BASE_URL", "https://restcountries.com/v3.1/all")
	url := fmt.Sprintf("%s?fields=name,capital,population,region,subregion,flags,languages,currencies,cca2,cca3", baseURL)
	err := utils.GetJSON(url, &raw)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch countries: %w", err)
	}

	var dtos []models.CountryDTO
	for _, c := range raw {
		dtos = append(dtos, transformCountry(c))
	}

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
	capital := "N/A"
	if len(c.Capital) > 0 {
		capital = c.Capital[0]
	}

	var currencyParts []string
	for code, cur := range c.Currencies {
		if cur.Name != "" {
			currencyParts = append(currencyParts, fmt.Sprintf("%s (%s)", cur.Name, code))
		}
	}
	currencies := strings.Join(currencyParts, ", ")
	if currencies == "" {
		currencies = "N/A"
	}

	return models.CountryDTO{
		Slug:       utils.CountryNameToSlug(c.Name.Common),
		Name:       c.Name.Common,
		Capital:    capital,
		Population: utils.FormatPopulation(c.Population),
		Region:     c.Region,
		Subregion:  c.Subregion,
		FlagURL:    c.Flags.Png,
		FlagAlt:    c.Flags.Alt,
		Languages:  utils.FormatLanguages(c.Languages),
		Currencies: currencies,
		Cca2:       c.Cca2,
	}
}
