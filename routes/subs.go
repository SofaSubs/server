package routes

import (
	"github.com/SofaSubs/server/controller"
	"github.com/labstack/echo/v4"
)

func GetSubsApiRoutes(e *echo.Echo, userController *controller.SubsController) {
	v1 := e.Group("/api/v1")
	{
		v1.GET("/subs", userController.GetSubs)
	}
}
