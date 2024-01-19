package controller

import (
	"net/http"
	"strconv"

	"github.com/flouressaint/schedule-service/internal/entity"
	"github.com/flouressaint/schedule-service/internal/service"
	"github.com/labstack/echo/v4"
)

type disciplineRoutes struct {
	disciplineService service.Discipline
}

func newDisciplineRoutes(g *echo.Group, disciplineService service.Discipline, middleware ...echo.MiddlewareFunc) {
	r := &disciplineRoutes{
		disciplineService: disciplineService,
	}

	g.Use(middleware...)

	g.POST("", r.CreateDiscipline)
	g.GET("", r.GetDisciplines)
	g.DELETE("/:id", r.DeleteDiscipline)
	g.PUT("/:id", r.UpdateDiscipline)
}

type disciplineInput struct {
	Name string `json:"name" validate:"required"`
}

// @Summary Create discipline
// @Description Create discipline
// @Tags discipline
// @Accept json
// @Produce json
// @Param input body disciplineInput true "input"
// @Success 201 {object} controller.disciplineRoutes.Creatediscipline.response
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/discipline [post]
func (r *disciplineRoutes) CreateDiscipline(c echo.Context) error {
	var input disciplineInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}
	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	data := entity.Discipline{
		Name: input.Name,
	}

	id, err := r.disciplineService.CreateDiscipline(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	type response struct {
		Id uint `json:"id"`
	}

	return c.JSON(http.StatusCreated, response{
		Id: id,
	})
}

// @Summary Get disciplines
// @Description Get disciplines
// @Tags discipline
// @Accept json
// @Produce json
// @Success 200 {array} entity.Discipline
// @Failure 401 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/discipline [get]
func (r *disciplineRoutes) GetDisciplines(c echo.Context) error {
	disciplines, err := r.disciplineService.GetDisciplines()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, disciplines)
}

// @Summary Delete discipline
// @Description Delete discipline
// @Tags discipline
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} controller.disciplineRoutes.Deletediscipline.response
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/discipline/{id} [delete]
func (r *disciplineRoutes) DeleteDiscipline(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = r.disciplineService.DeleteDiscipline(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	type response struct {
		Status string `json:"status"`
	}

	return c.JSON(http.StatusOK, response{
		Status: "ok",
	})
}

// @Summary Update discipline
// @Description Update discipline
// @Tags discipline
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param input body disciplineInput true "input"
// @Success 200 {object} controller.disciplineRoutes.Updatediscipline.response
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/discipline/{id} [put]
func (r *disciplineRoutes) UpdateDiscipline(c echo.Context) error {
	var input disciplineInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}
	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	data := entity.Discipline{
		Name: input.Name,
	}

	err = r.disciplineService.UpdateDiscipline(uint(id), data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	type response struct {
		Status string `json:"status"`
	}

	return c.JSON(http.StatusOK, response{
		Status: "ok",
	})
}
