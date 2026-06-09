package apicontrollers

import (
	"TravelSphere/models"
	"TravelSphere/services"
	"TravelSphere/utils"
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type WishlistAPIController struct {
	web.Controller
}

func (c *WishlistAPIController) getUserID() string {
	user := c.GetSession("username")
	if user != nil {
		return user.(string)
	}
	sid := c.Ctx.GetCookie("beegosessionID")
	if sid == "" {
		return "anonymous"
	}
	return sid
}

// GET /api/wishlist
func (c *WishlistAPIController) List() {
	userID := c.getUserID()
	items := services.GetWishlist(userID)
	c.Data["json"] = utils.SuccessResponse(items)
	c.ServeJSON()
}

// POST /api/wishlist
func (c *WishlistAPIController) Create() {
	var req models.CreateWishlistRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = utils.ValidationErrorResponse("Invalid JSON payload")
		c.ServeJSON()
		return
	}

	if !utils.ValidateCountryName(req.CountryName) {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = utils.ValidationErrorResponse("country_name is required")
		c.ServeJSON()
		return
	}
	if req.Status != "" && !utils.ValidateWishlistStatus(req.Status) {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = utils.ValidationErrorResponse("status must be Planned or Visited")
		c.ServeJSON()
		return
	}

	userID := c.getUserID()
	item, err := services.CreateWishlistItem(userID, req)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = utils.ErrorResponse(err.Error())
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = utils.SuccessMessageResponse("Added to wishlist", item)
	c.ServeJSON()
}

// PUT /api/wishlist/:id
func (c *WishlistAPIController) Update() {
	id := c.Ctx.Input.Param(":id")
	var req models.UpdateWishlistRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = utils.ValidationErrorResponse("Invalid JSON payload")
		c.ServeJSON()
		return
	}

	if !utils.ValidateWishlistStatus(req.Status) {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = utils.ValidationErrorResponse("status must be Planned or Visited")
		c.ServeJSON()
		return
	}

	userID := c.getUserID()
	item, err := services.UpdateWishlistItem(userID, id, req)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = utils.ErrorResponse(err.Error())
		c.ServeJSON()
		return
	}

	c.Data["json"] = utils.SuccessMessageResponse("Updated", item)
	c.ServeJSON()
}

// DELETE /api/wishlist/:id
func (c *WishlistAPIController) Delete() {
	id := c.Ctx.Input.Param(":id")
	userID := c.getUserID()

	if err := services.DeleteWishlistItem(userID, id); err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = utils.ErrorResponse(err.Error())
		c.ServeJSON()
		return
	}

	c.Data["json"] = utils.SuccessMessageResponse("Deleted successfully", nil)
	c.ServeJSON()
}
