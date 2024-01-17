package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/flouressaint/schedule-service/internal/entity"
	"github.com/flouressaint/schedule-service/internal/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type lessonRoutes struct {
	lessonService service.Lesson
}

func newLessonRoutes(g *echo.Group, lessonService service.Lesson) {
	r := &lessonRoutes{
		lessonService: lessonService,
	}

	g.POST("", r.CreateLesson)
	g.GET("", r.GetLessons)
	g.DELETE("/:id", r.DeleteLesson)
	g.PUT("/:id", r.UpdateLesson)
}

type lessonInput struct {
	Date     time.Time `json:"date" validate:"required" gorm:"type:timestamp"`
	Duration int       `json:"duration" validate:"required" gorm:"type:int;not null"`

	DisciplineID  uint      `json:"discipline_id" validate:"required" gorm:"type:int;not null"`
	TeacherUserID uuid.UUID `json:"teacher_user_id" validate:"required" gorm:"type:uuid"`
	StudyGroupID  uint      `json:"study_group_id" validate:"required" gorm:"type:int;not null"`
	AuditoriumID  uint      `json:"auditorium_id" validate:"required" gorm:"type:int;not null"`
	HometaskID    uint      `json:"hometask_id" validate:"required" gorm:"type:int;not null"`
}

// @Summary Create lesson
// @Description Create lesson
// @Tags lessons
// @Accept json
// @Produce json
// @Param input body lessonInput true "input"
// @Success 201 {object} controller.lessonRoutes.CreateLesson.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/admin/lesson [post]
func (r *lessonRoutes) CreateLesson(c echo.Context) error {
	var input lessonInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}
	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	data := entity.Lesson{
		Date:          input.Date,
		Duration:      input.Duration,
		DisciplineID:  input.DisciplineID,
		TeacherUserID: input.TeacherUserID,
		StudyGroupID:  input.StudyGroupID,
		AuditoriumID:  input.AuditoriumID,
		HometaskID:    input.HometaskID,
	}

	id, err := r.lessonService.CreateLesson(data)
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

// @Summary Get lessons
// @Description Get lessons
// @Tags lessons
// @Accept json
// @Produce json
// @Success 200 {array} entity.Lesson
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/admin/lesson [get]
func (r *lessonRoutes) GetLessons(c echo.Context) error {
	lessons, err := r.lessonService.GetLessons()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, lessons)
}

// @Summary Delete lesson
// @Description Delete lesson
// @Tags lessons
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} controller.lessonRoutes.DeleteLesson.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/admin/lesson/{id} [delete]
func (r *lessonRoutes) DeleteLesson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = r.lessonService.DeleteLesson(uint(id))
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

// @Summary Update lesson
// @Description Update lesson
// @Tags lessons
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param input body lessonInput true "input"
// @Success 200 {object} controller.lessonRoutes.UpdateLesson.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/admin/lesson/{id} [put]
func (r *lessonRoutes) UpdateLesson(c echo.Context) error {
	var input lessonInput
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

	data := entity.Lesson{
		Date:          input.Date,
		Duration:      input.Duration,
		DisciplineID:  input.DisciplineID,
		TeacherUserID: input.TeacherUserID,
		StudyGroupID:  input.StudyGroupID,
		AuditoriumID:  input.AuditoriumID,
		HometaskID:    input.HometaskID,
	}

	err = r.lessonService.UpdateLesson(uint(id), data)
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
