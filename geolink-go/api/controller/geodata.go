package controller

import (
	"geolink-go/api/service"
	"geolink-go/api/structs"
	"geolink-go/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GeoDataController struct {
	geoDataService service.GeoDataService
}

func NewGeoDataController(geoDataService service.GeoDataService) GeoDataController {
	return GeoDataController{
		geoDataService: geoDataService,
	}
}

func (c *GeoDataController) GetIpGeoData(ctx *gin.Context) {
	var request structs.GetGeoDataRequest

	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		util.HandleCommonValidationError(ctx, err)
		return
	}

	response, err := c.geoDataService.GetIpGeoData(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
