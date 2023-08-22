package router

import (
	"github.com/labstack/echo/v4"
	"mesh-uat-go-two/api"
)

func Routes(e *echo.Echo) {
	e.GET("/", api.MeshApi().Index)
	e.GET("/api", api.MeshApi().ApiIndex)
	e.GET("/health", api.MeshApi().Health)
	e.GET("/api/v1/hello/one", api.MeshApi().HelloApiOne)
	e.GET("/api/v1/hello/two", api.MeshApi().HelloApiTwo)
	e.GET("/api/v1/mesh/one", api.MeshApi().MeshUatOneApiOne)
	e.GET("/api/v1/mesh/two", api.MeshApi().MeshUatOneApiTwo)
}
