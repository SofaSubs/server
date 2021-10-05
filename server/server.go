package server

import (
	"github.com/SofaSubs/server/controller"
	"github.com/SofaSubs/server/routes"
	"github.com/labstack/echo/v4"
)

var subsController *controller.SubsController

func Start() {
	e := echo.New()

	routes.GetSubsApiRoutes(e, subsController)

	e.Logger.Fatal(e.Start(":8000"))
}

func init() {
	subsController = controller.NewSubsController()
}
