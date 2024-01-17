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

	// newUserRoutes(handler.Group("/user"), services.User)

	authMiddleware := &AuthMiddleware{services.Auth}
	api := handler.Group("/api", authMiddleware.UserIdentity)
	{
		admin := api.Group("/admin", authMiddleware.isAdmin)
		{
			newAuditoriumRoutes(admin.Group("/auditorium"), services.Auditorium)
			newDisciplineRoutes(admin.Group("/discipline"), services.Discipline)
			newLessonRoutes(admin.Group("/lesson"), services.Lesson)
		}
		teacher := api.Group("/teacher", authMiddleware.isTeacher)
		{
			newHometaskRoutes(teacher.Group("/hometask"), services.Hometask)
		}

		// newTodoRoutes(api.Group("/todo"), services.Todo)
	}

}
