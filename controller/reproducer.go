package controller

import (
	"net/http"
	"strconv"

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

func (rController *ReproducerController) Pause(c echo.Context) error {

	if err := rController.reproducer.Pause(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "pause")
}

func (rController *ReproducerController) Stop(c echo.Context) error {

	if err := rController.reproducer.Stop(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "stop")
}

func (rController *ReproducerController) Time(c echo.Context) error {

	if time, err := rController.reproducer.GetTime(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		return c.JSON(http.StatusOK, strconv.FormatUint(uint64(time), 10))
	}
}
