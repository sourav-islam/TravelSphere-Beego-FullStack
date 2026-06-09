package routers

import (
	apicontrollers "TravelSphere/controllers/api"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

func InitAPIRoutes() {
	// Countries API
	web.Router("/api/countries", &apicontrollers.CountriesAPIController{}, "get:List")
	web.Router("/api/countries/:slug", &apicontrollers.CountriesAPIController{}, "get:Detail")

	// Wishlist API
	web.Router("/api/wishlist", &apicontrollers.WishlistAPIController{}, "get:List;post:Create")
	web.Router("/api/wishlist/:id", &apicontrollers.WishlistAPIController{}, "put:Update;delete:Delete")

	// // Dashboard API
	web.Router("/api/dashboard/summary", &apicontrollers.DashboardAPIController{}, "get:Summary")

	fmt.Println("API routes initialized")
}
