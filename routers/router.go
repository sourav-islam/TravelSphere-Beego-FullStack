package routers

import (
	"TravelSphere/controllers"
	"fmt"
	"log"
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func InitSSRRoutes() {
	// Logging filter
	web.InsertFilter("/*", web.BeforeRouter, func(ctx *context.Context) {
		start := time.Now()
		ctx.Input.SetData("startTime", start)
		log.Printf("[REQUEST] %s %s", ctx.Input.Method(), ctx.Input.URI())
	})

	web.InsertFilter("/*", web.AfterExec, func(ctx *context.Context) {
		start, ok := ctx.Input.GetData("startTime").(time.Time)
		if ok {
			log.Printf("[RESPONSE] %s %s — %v", ctx.Input.Method(), ctx.Input.URI(), time.Since(start))
		}
	}, web.WithReturnOnOutput(false))
	// SSR routes
	web.Router("/", &controllers.HomeController{}, "get:Get")
	web.Router("/login", &controllers.HomeController{}, "post:Login")
	web.Router("/logout", &controllers.HomeController{}, "get:Logout")
	web.Router("/search", &controllers.HomeController{}, "get:Search")
	web.Router("/countries", &controllers.CountryController{}, "get:List")
	web.Router("/countries/:slug", &controllers.CountryController{}, "get:Detail")
	web.Router("/wishlist", &controllers.WishlistController{}, "get:Get")
	// web.Router("/dashboard", &controllers.DashboardController{}, "get:Get")
	fmt.Println("SSR routes initialized")
}

func authFilter(ctx *context.Context) {
	user := ctx.Input.Session("username")
	if user == nil {
		ctx.Redirect(302, "/")
	}
}
