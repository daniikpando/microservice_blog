package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRoutes() *echo.Echo{

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	ArticleRouterV1(e)

	return e;
}