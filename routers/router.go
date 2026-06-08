package routers

import (
	"TravelSphere/controllers"

	"github.com/beego/beego/v2/server/web"
)

func InitSSRRoutes() {

	// SSR routes
	web.Router("/", &controllers.HomeController{}, "get:Get")
	web.Router("/login", &controllers.HomeController{}, "post:Login")
	web.Router("/logout", &controllers.HomeController{}, "get:Logout")
	web.Router("/search", &controllers.HomeController{}, "get:Search")
	web.Router("/countries", &controllers.CountryController{}, "get:List")
	web.Router("/countries/:slug", &controllers.CountryController{}, "get:Detail")
	// web.Router("/wishlist", &controllers.WishlistController{}, "get:Get")
	// web.Router("/dashboard", &controllers.DashboardController{}, "get:Get")
}

func InitAPIRoutes() {
	// API routes
}
