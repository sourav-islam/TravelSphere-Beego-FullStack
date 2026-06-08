package models

import "time"

type WishlistItem struct {
	ID          string    `json:"id"`
	CountryName string    `json:"country_name"`
	Note        string    `json:"note"`
	Status      string    `json:"status"` // "Planned" or "Visited"
	CreatedAt   time.Time `json:"created_at"`
	UserID      string    `json:"user_id"`
}

type CreateWishlistRequest struct {
	CountryName string `json:"country_name"`
	Note        string `json:"note"`
	Status      string `json:"status"`
}

type UpdateWishlistRequest struct {
	Note   string `json:"note"`
	Status string `json:"status"`
}

type DashboardSummary struct {
	Total   int `json:"total"`
	Planned int `json:"planned"`
	Visited int `json:"visited"`
}
