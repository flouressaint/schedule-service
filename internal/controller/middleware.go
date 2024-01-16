package controller

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/flouressaint/schedule-service/internal/service"
	"github.com/labstack/echo/v4"
)

const (
	userIdCtx    = "userId"
	userRolesCtx = "userRoles"
)

type AuthMiddleware struct {
	authService service.Auth
}

func (h *AuthMiddleware) UserIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := bearerToken(c.Request())
		if !ok {
			log.Printf("AuthMiddleware.UserIdentity: bearerToken: %v", fmt.Errorf("invalid auth header"))
			newErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("invalid auth header").Error())
			return nil
		}

		userId, userRoles, err := h.authService.ParseToken(token)
		if err != nil {
			log.Printf("AuthMiddleware.UserIdentity: h.authService.ParseToken: %v", err)
			newErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("cannot parse token: %v", err).Error())
			return err
		}
		c.Set(userIdCtx, userId)
		c.Set(userRolesCtx, userRoles)

		return next(c)
	}
}

func (h *AuthMiddleware) isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRoles := c.Get(userRolesCtx).([]string)
		for _, role := range userRoles {
			if role == "admin" {
				return next(c)
			}
		}

		newErrorResponse(c, http.StatusForbidden, "forbidden")
		return nil
	}
}

func (h *AuthMiddleware) isTeacher(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRoles := c.Get(userRolesCtx).([]string)
		for _, role := range userRoles {
			if role == "teacher" {
				return next(c)
			}
		}

		newErrorResponse(c, http.StatusForbidden, "forbidden")
		return nil
	}
}

func bearerToken(r *http.Request) (string, bool) {
	const prefix = "Bearer "

	header := r.Header.Get(echo.HeaderAuthorization)
	if header == "" {
		return "", false
	}

	if len(header) > len(prefix) && strings.EqualFold(header[:len(prefix)], prefix) {
		return header[len(prefix):], true
	}

	return "", false
}
