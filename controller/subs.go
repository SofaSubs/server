package controller

import (
	"fmt"
	"net/http"

	"github.com/SofaSubs/server/mock"
	"github.com/labstack/echo/v4"
)

type SubsController struct {
}

func NewSubsController() *SubsController {
	return &SubsController{}
}

func (subsController *SubsController) GetSubs(c echo.Context) error {

	subs, err := mock.GetSubs()

	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, subs)
}
