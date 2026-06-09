package services

import "TravelSphere/models"

func GetDashboardSummary(userID string) models.DashboardSummary {
	items := GetWishlist(userID)
	summary := models.DashboardSummary{Total: len(items)}
	for _, item := range items {
		if item.Status == "Planned" {
			summary.Planned++
		} else if item.Status == "Visited" {
			summary.Visited++
		}
	}
	return summary
}
