package controllers

import (
	"TravelSphere/services"
)

type DashboardController struct {
	BaseController
}

func (c *DashboardController) Prepare() {
	c.BaseController.Prepare()
	c.Data["ActiveMenu"] = "dashboard"
}

// SSR: /dashboard (protected)
func (c *DashboardController) Get() {
	if !c.RequireLogin() {
		return
	}
	userID := c.GetUserID()
	summary := services.GetDashboardSummary(userID)
	items := services.GetWishlist(userID)

	c.Data["Summary"] = summary
	c.Data["WishlistItems"] = items
	c.TplName = "dashboard.tpl"
}
