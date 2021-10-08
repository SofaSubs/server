package controller

import (
	"net/http"

	"github.com/SofaSubs/server/internal/reproducers"
	"github.com/labstack/echo/v4"
)

type ReproducerController struct {
	reproducer reproducers.Reproducer
}

func NewReproducerController() (*ReproducerController, error) {
	//TODO: use strategy pattern
	reproducer, err := reproducers.NewVlcReproducer()
	if err != nil {
		return nil, err
	}

	return &ReproducerController{
		reproducer: reproducer,
	}, nil
}

func (rController *ReproducerController) Play(c echo.Context) error {

	if err := rController.reproducer.Play(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "play")
}

func (rController *ReproducerController) Stop(c echo.Context) error {

	if err := rController.reproducer.Stop(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "stop")
}
