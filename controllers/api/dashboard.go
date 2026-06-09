package apicontrollers

import (
	"TravelSphere/services"
	"TravelSphere/utils"

	"github.com/beego/beego/v2/server/web"
)

type DashboardAPIController struct {
	web.Controller
}

// GET /api/dashboard/summary
func (c *DashboardAPIController) Summary() {
	user := c.GetSession("username")
	userID := "anonymous"
	if user != nil {
		userID = user.(string)
	}
	summary := services.GetDashboardSummary(userID)
	c.Data["json"] = utils.SuccessResponse(summary)
	c.ServeJSON()
}
