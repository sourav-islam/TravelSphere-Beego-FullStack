package services

import (
	"TravelSphere/models"
	"TravelSphere/utils"
	"fmt"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// GetAttractionsByCapital looks up coordinates for the capital name,
// then fetches nearby attractions — no hardcoded lat/lon needed.
func GetAttractionsByCapital(capital string) ([]models.AttractionDTO, error) {

	apiKey, _ := beego.AppConfig.String("OPENTRIPMAP_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("OPENTRIPMAP_API_KEY not set")
	}

	lat, lon, err := geocodeCapital(capital, apiKey)
	if err != nil {
		return nil, fmt.Errorf("geocode failed for %s: %w", capital, err)
	}

	return fetchAttractions(lat, lon, apiKey)
}

// geocodeCapital uses OpenTripMap geoname endpoint to resolve a city to coordinates.
func geocodeCapital(capital, apiKey string) (float64, float64, error) {
	url := fmt.Sprintf(
		"https://api.opentripmap.com/0.1/en/places/geoname?name=%s&apikey=%s",
		strings.ReplaceAll(capital, " ", "+"),
		apiKey,
	)

	var result struct {
		Lat    float64 `json:"lat"`
		Lon    float64 `json:"lon"`
		Name   string  `json:"name"`
		Status string  `json:"status"`
	}

	if err := utils.GetJSON(url, &result); err != nil {
		return 0, 0, err
	}

	if result.Status == "NOT_FOUND" || (result.Lat == 0 && result.Lon == 0) {
		return 0, 0, fmt.Errorf("city not found: %s", capital)
	}

	return result.Lat, result.Lon, nil
}

// fetchAttractions calls OpenTripMap radius endpoint with resolved coordinates.
func fetchAttractions(lat, lon float64, apiKey string) ([]models.AttractionDTO, error) {
	url := fmt.Sprintf(
		"https://api.opentripmap.com/0.1/en/places/radius?radius=50000&lon=%f&lat=%f&kinds=interesting_places&limit=10&apikey=%s",
		lon, lat, apiKey,
	)

	var response models.AttractionResponse
	if err := utils.GetJSON(url, &response); err != nil {
		return nil, fmt.Errorf("failed to fetch attractions: %w", err)
	}

	var dtos []models.AttractionDTO
	for _, f := range response.Features {
		name := strings.TrimSpace(f.Properties.Name)
		if name == "" {
			continue
		}
		dtos = append(dtos, models.AttractionDTO{
			Name:  name,
			Kinds: f.Properties.Kinds,
			XID:   f.Properties.XID,
		})
	}
	return dtos, nil
}

// GetCapitalCoords kept for any legacy use but no longer needed by Detail().
func GetCapitalCoords(capital string) (float64, float64, bool) {
	coords := map[string][2]float64{
		"Paris":     {48.8566, 2.3522},
		"Tokyo":     {35.6762, 139.6503},
		"London":    {51.5074, -0.1278},
		"Dhaka":     {23.8103, 90.4125},
		"New Delhi": {28.6139, 77.2090},
		"Beijing":   {39.9042, 116.4074},
		"Canberra":  {-35.2809, 149.1300},
	}
	if c, ok := coords[capital]; ok {
		return c[0], c[1], true
	}
	return 0, 0, false
}

func GetPopularAttractions() []models.AttractionDTO {
	return []models.AttractionDTO{
		{Name: "Eiffel Tower", Kinds: "architecture,historic"},
		{Name: "Grand Canyon", Kinds: "natural"},
		{Name: "Sydney Opera House", Kinds: "architecture,theatre"},
		{Name: "Colosseum", Kinds: "historic,architecture"},
	}
}
