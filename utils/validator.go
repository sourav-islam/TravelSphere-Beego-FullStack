package utils

import "strings"

var allowedStatuses = map[string]bool{
	"Planned": true,
	"Visited": true,
}

func ValidateWishlistStatus(status string) bool {
	return allowedStatuses[status]
}

func ValidateCountryName(name string) bool {
	return strings.TrimSpace(name) != ""
}

func ValidateSlug(slug string) bool {
	return strings.TrimSpace(slug) != "" && len(slug) <= 100
}
