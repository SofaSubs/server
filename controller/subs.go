package controller

import (
	"net/http"

	"github.com/SofaSubs/server/model"
	"github.com/SofaSubs/server/reader"
	"github.com/SofaSubs/server/repository"
	"github.com/labstack/echo/v4"
)

type SubsController struct {
	subsRepository repository.SubsRepository
}

func NewSubsController(subsRepository repository.SubsRepository) *SubsController {
	return &SubsController{
		subsRepository: subsRepository,
	}
}

func (subsController *SubsController) GetSubs(c echo.Context) error {

	subs, err := reader.GetSubs()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, subs)
}

func (subsController *SubsController) AddSub(c echo.Context) error {

	var sub model.Sub
	if err := c.Bind(&sub); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if _, err := subsController.subsRepository.AddSub(sub); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "sub added")
}

func (subsController *SubsController) FindSubs(c echo.Context) error {
	if subs, err := subsController.subsRepository.FindSubs(0, 10); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		return c.JSON(http.StatusOK, subs)
	}
}
