package controller

import (
	"net/http"

	"github.com/SofaSubs/server/reader"
	"github.com/labstack/echo/v4"
)

type SubsController struct {
}

func NewSubsController() *SubsController {
	return &SubsController{}
}

func (subsController *SubsController) GetSubs(c echo.Context) error {

	subs, err := reader.GetSubs()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, subs)
}
