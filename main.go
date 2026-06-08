package main

import (
	"TravelSphere/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	// Explicit asset directories and configs map
	web.BConfig.WebConfig.ViewsPath = "views"
	web.BConfig.WebConfig.StaticDir["/static"] = "static"

	// Initialize path route handlers
	routers.InitSSRRoutes()
	routers.InitAPIRoutes()

	// Run application instance (Reads conf/app.conf natively)
	web.Run()
}
