package routes

import (
	"geolink-go/api/controller"
	"geolink-go/infrastructure"
)

type GeoDataRoute struct {
	Controller controller.GeoDataController
	Handler    infrastructure.GinRouter
}

func NewGeoDataRoute(geoDataController controller.GeoDataController, router infrastructure.GinRouter) GeoDataRoute {
	return GeoDataRoute{
		Controller: geoDataController,
		Handler:    router,
	}
}

func (r *GeoDataRoute) Setup() {
	configuration := r.Handler.Gin.Group("/geo")
	{
		configuration.GET("/", r.Controller.GetIpGeoData)
	}
}
