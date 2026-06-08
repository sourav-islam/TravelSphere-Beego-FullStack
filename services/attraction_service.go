package services

import (
	"TravelSphere/models"
	"TravelSphere/utils"
	"fmt"
	"os"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

func GetAttractionsByCountry(lat, lon float64) ([]models.AttractionDTO, error) {
	apiKey := web.AppConfig.DefaultString("OPENTRIPMAP_API_KEY", "")
	if apiKey == "" {
		apiKey = os.Getenv("OPENTRIPMAP_API_KEY")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("OPENTRIPMAP_API_KEY not set")
	}

	baseURL := web.AppConfig.DefaultString("OPENTRIPMAP_BASE_URL", "https://api.opentripmap.com/0.1/en/places/radius")
	url := fmt.Sprintf(
		"%s?radius=50000&lon=%f&lat=%f&kinds=interesting_places&limit=10&apikey=%s",
		baseURL, lon, lat, apiKey,
	)

	var response models.AttractionResponse
	err := utils.GetJSON(url, &response)
	if err != nil {
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

func GetCapitalCoords(capital string) (float64, float64, bool) {
	coords := map[string][2]float64{
		"Paris":            {48.8566, 2.3522},
		"Tokyo":            {35.6762, 139.6503},
		"Washington":       {38.9072, -77.0369},
		"London":           {51.5074, -0.1278},
		"Canberra":         {-35.2809, 149.1300},
		"Dhaka":            {23.8103, 90.4125},
		"New Delhi":        {28.6139, 77.2090},
		"Beijing":          {39.9042, 116.4074},
		"Berlin":           {52.5200, 13.4050},
		"Brasília":         {-15.7975, -47.8919},
		"Buenos Aires":     {-34.6037, -58.3816},
		"Cairo":            {30.0444, 31.2357},
		"Rome":             {41.9028, 12.4964},
		"Madrid":           {40.4168, -3.7038},
		"Moscow":           {55.7558, 37.6173},
		"Tirana":           {41.3317, 19.8319},
		"Kabul":            {34.5553, 69.2075},
		"Algiers":          {36.7372, 3.0865},
		"Andorra la Vella": {42.5063, 1.5218},
		"Luanda":           {-8.8368, 13.2343},
		"Saint John's":     {17.1274, -61.8468},
		"Oranjestad":       {12.5246, -70.0265},
		"Yerevan":          {40.1792, 44.4991},
		"Vienna":           {48.2082, 16.3738},
		"Baku":             {40.4093, 49.8671},
		"Nassau":           {25.0343, -77.3963},
		"Manama":           {26.2285, 50.5860},
		"Seoul":            {37.5665, 126.9780},
		"Bangkok":          {13.7563, 100.5018},
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
