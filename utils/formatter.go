package utils

import (
	"fmt"
	"strings"
)

func FormatPopulation(pop int64) string {
	if pop >= 1_000_000_000 {
		return fmt.Sprintf("%.1fB", float64(pop)/1_000_000_000)
	} else if pop >= 1_000_000 {
		return fmt.Sprintf("%.1fM", float64(pop)/1_000_000)
	} else if pop >= 1_000 {
		return fmt.Sprintf("%.1fK", float64(pop)/1_000)
	}
	return fmt.Sprintf("%d", pop)
}

func FormatLanguages(langs map[string]string) string {
	if len(langs) == 0 {
		return "N/A"
	}
	var list []string
	for _, v := range langs {
		list = append(list, v)
	}
	return strings.Join(list, ", ")
}

func FormatCurrencies(currencies map[string]interface{}) string {
	return "Various"
}

func CountryNameToSlug(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "'", "")
	name = strings.ReplaceAll(name, ",", "")
	name = strings.ReplaceAll(name, ".", "")
	return name
}

func SlugToSearch(slug string) string {
	return strings.ReplaceAll(slug, "-", " ")
}
