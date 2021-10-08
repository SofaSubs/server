package routes

import (
	"github.com/SofaSubs/server/controller"
	"github.com/labstack/echo/v4"
)

func Reproducer(e *echo.Echo, userController *controller.ReproducerController) {
	v1 := e.Group("/api/v1")
	{
		v1.GET("/play", userController.Play)
		v1.GET("/stop", userController.Stop)
	}
}
