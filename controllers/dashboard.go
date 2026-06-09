package controllers

import "TravelSphere/services"

type DashboardController struct {
	BaseController
}

func (c *DashboardController) Prepare() {
	c.BaseController.Prepare()
	c.Data["ActiveMenu"] = "dashboard"
}

func (c *DashboardController) Get() {
	if !c.RequireLogin() {
		return
	}
	userID := c.GetUserID()
	c.Data["Summary"] = services.GetDashboardSummary(userID)
	c.Data["WishlistItems"] = services.GetWishlist(userID)
	c.Layout = "layout/main.tpl"
	c.TplName = "dashboard.tpl"
}
