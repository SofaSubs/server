package server

import (
	"github.com/SofaSubs/server/controller"
	"github.com/SofaSubs/server/routes"
	"github.com/labstack/echo/v4"
)

var subsController *controller.SubsController

func Start() {
	e := echo.New()

	routes.GetUserApiRoutes(e, subsController)

	// echo server 9000 de başlatıldı.
	e.Logger.Fatal(e.Start(":8000"))
}

func init() {
	subsController = controller.NewSubsController()
}
