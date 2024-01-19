package controller

import (
	_ "github.com/flouressaint/schedule-service/docs"
	"github.com/flouressaint/schedule-service/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(handler *echo.Echo, services *service.Services) {
	handler.GET("/health", func(c echo.Context) error { return c.NoContent(200) })
	handler.GET("/swagger/*", echoSwagger.WrapHandler)

	authMiddleware := &AuthMiddleware{services.Auth}
	api := handler.Group("/api", authMiddleware.UserIdentity)
	{
		// admin
		newAuditoriumRoutes(api.Group("/auditorium"), services.Auditorium, authMiddleware.AdminIdentity)
		newDisciplineRoutes(api.Group("/discipline"), services.Discipline, authMiddleware.AdminIdentity)
		newLessonRoutes(api.Group("/lesson"), services.Lesson, authMiddleware.AdminIdentity)

		// teacher
		newHometaskRoutes(api.Group("/hometask"), services.Hometask, authMiddleware.TeacherIdentity)

		// student and others
		newLessonRoutesForStudent(api.Group("/lesson"), services.Lesson)
	}

}
