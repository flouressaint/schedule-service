package controller

import (
	"net/http"
	"strconv"

	"github.com/flouressaint/schedule-service/internal/entity"
	"github.com/flouressaint/schedule-service/internal/service"
	"github.com/labstack/echo/v4"
)

type hometaskRoutes struct {
	hometaskService service.Hometask
}

func newHometaskRoutes(g *echo.Group, hometaskService service.Hometask) {
	r := &hometaskRoutes{
		hometaskService: hometaskService,
	}

	g.POST("", r.CreateHometask)
	g.GET("", r.GetHometasks)
	g.DELETE("/:id", r.DeleteHometask)
	g.PUT("/:id", r.UpdateHometask)
}

type hometaskInput struct {
	Description string `json:"description" validate:"required" gorm:"type:varchar(255);not null"`
	Attachment  string `json:"attachment" gorm:"type:varchar(255)"`
}

// @Summary Create hometask
// @Description Create hometask
// @Tags hometasks
// @Accept json
// @Produce json
// @Param input body hometaskInput true "input"
// @Success 201 {object} controller.hometaskRoutes.CreateHometask.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/teacher/hometask [post]
func (r *hometaskRoutes) CreateHometask(c echo.Context) error {
	var input hometaskInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	data := entity.Hometask{
		Description: input.Description,
		Attachment:  input.Attachment,
	}

	id, err := r.hometaskService.CreateHometask(data)
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

// @Summary Get hometasks
// @Description Get hometasks
// @Tags hometasks
// @Accept json
// @Produce json
// @Success 200 {array} entity.Hometask
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/teacher/hometask [get]
func (r *hometaskRoutes) GetHometasks(c echo.Context) error {
	hometasks, err := r.hometaskService.GetHometasks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, hometasks)
}

// @Summary Delete hometask
// @Description Delete hometask
// @Tags hometasks
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} controller.hometaskRoutes.DeleteHometask.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/teacher/hometask/{id} [delete]
func (r *hometaskRoutes) DeleteHometask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = r.hometaskService.DeleteHometask(uint(id))
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

// @Summary Update hometask
// @Description Update hometask
// @Tags hometasks
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param input body hometaskInput true "input"
// @Success 200 {object} controller.hometaskRoutes.UpdateHometask.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/teacher/hometask/{id} [put]
func (r *hometaskRoutes) UpdateHometask(c echo.Context) error {
	var input hometaskInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
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

	data := entity.Hometask{
		Description: input.Description,
		Attachment:  input.Attachment,
	}

	err = r.hometaskService.UpdateHometask(uint(id), data)
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
