package server

import (
	"fmt"

	"github.com/SofaSubs/server/controller"
	"github.com/SofaSubs/server/db"
	"github.com/SofaSubs/server/repository"
	"github.com/SofaSubs/server/routes"
	"github.com/labstack/echo/v4"
)

var subsController *controller.SubsController
var reproducerController *controller.ReproducerController
var err error

func Start() {
	e := echo.New()

	routes.GetSubsApiRoutes(e, subsController)
	routes.Reproducer(e, reproducerController)

	e.Logger.Fatal(e.Start(":8000"))
}

func init() {
	db, _ := db.Init()
	subsRepository := repository.NewSubsRepository(db)

	subsController = controller.NewSubsController(subsRepository)
	reproducerController, err = controller.NewReproducerController()

	if err != nil {
		fmt.Println(err)
	}
}
