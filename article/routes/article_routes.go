package routes

import (
	"github.com/labstack/echo"
	"github.com/daniel/basic_microservice/article/controllers"
)




func ArticleRouterV1(e *echo.Echo)  {

	GeneralRoute := "/api/v1/articles"
	UniqueRoute := GeneralRoute + "/:id"

	e.GET(GeneralRoute, controllers.Index)
	e.POST(GeneralRoute, controllers.Create)

	e.PUT(UniqueRoute, controllers.Update)
	e.GET(UniqueRoute, controllers.Show)
	e.DELETE(UniqueRoute, controllers.Delete)

}