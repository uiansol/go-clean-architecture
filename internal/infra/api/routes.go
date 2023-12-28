package api

import (
	"github.com/labstack/echo/v4"
)

func (s *RestServer) SetUpRoutes() {
	s.UserRoutes()
	// Add new routes here, prefered grouped by domain
}

func (s *RestServer) UserRoutes() {
	s.router.GET("/user", func(c echo.Context) error {
		return s.appHandler.userCreateHandler.Handle(c)
	})
}
