package controller

import (
	"net/http"
	"strconv"

	"github.com/flouressaint/schedule-service/internal/entity"
	"github.com/flouressaint/schedule-service/internal/service"
	"github.com/labstack/echo/v4"
)

type auditoriumRoutes struct {
	auditoriumService service.Auditorium
}

func newAuditoriumRoutes(g *echo.Group, auditoriumService service.Auditorium) {
	r := &auditoriumRoutes{
		auditoriumService: auditoriumService,
	}

	g.POST("", r.CreateAuditorium)
	g.GET("", r.GetAuditoriums)
	g.DELETE("/:id", r.DeleteAuditorium)
	g.PUT("/:id", r.UpdateAuditorium)
}

type auditoriumInput struct {
	Name string `json:"name" validate:"required"`
}

// @Summary Create auditorium
// @Description Create auditorium
// @Tags auditoriums
// @Accept json
// @Produce json
// @Param input body auditoriumInput true "input"
// @Success 201 {object} controller.auditoriumRoutes.CreateAuditorium.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/auditorium [post]
func (r *auditoriumRoutes) CreateAuditorium(c echo.Context) error {
	var input auditoriumInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}
	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	data := entity.Auditorium{
		Name: input.Name,
	}

	id, err := r.auditoriumService.CreateAuditorium(data)
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

// @Summary Get auditoriums
// @Description Get auditoriums
// @Tags auditoriums
// @Accept json
// @Produce json
// @Success 200 {array} entity.Auditorium
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/auditorium [get]
func (r *auditoriumRoutes) GetAuditoriums(c echo.Context) error {
	auditoriums, err := r.auditoriumService.GetAuditoriums()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, auditoriums)
}

// @Summary Delete auditorium
// @Description Delete auditorium
// @Tags auditoriums
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} controller.auditoriumRoutes.Deleteauditorium.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/auditorium/{id} [delete]
func (r *auditoriumRoutes) DeleteAuditorium(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = r.auditoriumService.DeleteAuditorium(uint(id))
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

// @Summary Update auditorium
// @Description Update auditorium
// @Tags auditoriums
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param input body auditoriumInput true "input"
// @Success 200 {object} controller.auditoriumRoutes.Updateauditorium.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/auditorium/{id} [put]
func (r *auditoriumRoutes) UpdateAuditorium(c echo.Context) error {
	var input auditoriumInput
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

	data := entity.Auditorium{
		Name: input.Name,
	}

	err = r.auditoriumService.UpdateAuditorium(uint(id), data)
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
