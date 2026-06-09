package services

import (
	"TravelSphere/models"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// In-memory store per user
var (
	wishlistStore = make(map[string][]models.WishlistItem) // userID -> items
	wishlistMu    sync.RWMutex
)

func GetWishlist(userID string) []models.WishlistItem {
	wishlistMu.RLock()
	defer wishlistMu.RUnlock()
	items := wishlistStore[userID]
	if items == nil {
		return []models.WishlistItem{}
	}
	return items
}

func CreateWishlistItem(userID string, req models.CreateWishlistRequest) (*models.WishlistItem, error) {
	if req.CountryName == "" {
		return nil, fmt.Errorf("country_name is required")
	}
	if req.Status == "" {
		req.Status = "Planned"
	}
	if req.Status != "Planned" && req.Status != "Visited" {
		return nil, fmt.Errorf("invalid status: must be Planned or Visited")
	}

	item := models.WishlistItem{
		ID:          uuid.New().String(),
		CountryName: req.CountryName,
		Note:        req.Note,
		Status:      req.Status,
		CreatedAt:   time.Now(),
		UserID:      userID,
	}

	wishlistMu.Lock()
	wishlistStore[userID] = append(wishlistStore[userID], item)
	wishlistMu.Unlock()

	return &item, nil
}

func UpdateWishlistItem(userID, id string, req models.UpdateWishlistRequest) (*models.WishlistItem, error) {
	if req.Status != "Planned" && req.Status != "Visited" {
		return nil, fmt.Errorf("invalid status: must be Planned or Visited")
	}

	wishlistMu.Lock()
	defer wishlistMu.Unlock()

	items := wishlistStore[userID]
	for i, item := range items {
		if item.ID == id {
			items[i].Note = req.Note
			items[i].Status = req.Status
			wishlistStore[userID] = items
			return &items[i], nil
		}
	}
	return nil, fmt.Errorf("wishlist item not found")
}

func DeleteWishlistItem(userID, id string) error {
	wishlistMu.Lock()
	defer wishlistMu.Unlock()

	items := wishlistStore[userID]
	for i, item := range items {
		if item.ID == id {
			wishlistStore[userID] = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("wishlist item not found")
}
