package controllers

import (
	"TravelSphere/services"
)

type WishlistController struct {
	BaseController
}

func (c *WishlistController) Prepare() {
	c.BaseController.Prepare()
	c.Data["ActiveMenu"] = "wishlist"
}

// SSR: /wishlist (protected)
func (c *WishlistController) Get() {
	if !c.RequireLogin() {
		return
	}
	userID := c.GetUserID()
	items := services.GetWishlist(userID)
	c.Data["WishlistItems"] = items
	c.Layout = "layout/main.tpl"
	c.TplName = "wishlist.tpl"
}
