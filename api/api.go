package api

import (
	"github.com/labstack/echo/v4"
	"log"
	_type "mesh-uat-go-two/types"
	"mesh-uat-go-two/utils"
	"net/http"
)

type HomeInf interface {
	Index(c echo.Context) error
	ApiIndex(c echo.Context) error
	Health(c echo.Context) error
	HelloApiOne(c echo.Context) error
	HelloApiTwo(c echo.Context) error
	MeshUatOneApiOne(c echo.Context) error
	MeshUatOneApiTwo(c echo.Context) error
}

type home struct{}

func MeshApi() HomeInf {
	return &home{}
}

func (h *home) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"appName": "Mesh-UAT-Go-Two",
	})
}

func (h *home) ApiIndex(c echo.Context) error {
	response := _type.ResponseDto{
		Message: "This is Mesh-UAT-Go-Two Api",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *home) Health(c echo.Context) error {
	return c.String(http.StatusOK, "[Mesh UAT Go Two] I am live!")
}

func (h *home) HelloApiOne(c echo.Context) error {
	response := _type.ResponseDto{
		Message: "Hello From Mesh-UAT-Go-Two API One",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *home) HelloApiTwo(c echo.Context) error {
	response := _type.ResponseDto{
		Message: "Hello From Mesh-UAT-Go-Two API Two",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *home) MeshUatOneApiOne(c echo.Context) error {
	resp, err := utils.MeshUatOneAppApiMeshOne()
	if err != nil {
		log.Println("Error Occurred Mesh-UAT-Go-One Api One: " + err.Error())
		return c.JSON(http.StatusBadRequest, _type.Response().Error("Error Occurred Mesh-UAT-Go-One Api One: "+err.Error()))
	}

	response := _type.ResponseDto{
		Message: "Mesh-UAT-Go-Two API --TO-- Mesh-UAT-Go-One Api One",
		Data:    resp,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *home) MeshUatOneApiTwo(c echo.Context) error {
	resp, err := utils.MeshUatOneAppApiMeshTwo()
	if err != nil {
		log.Println("Error Occurred Mesh-UAT-Go-One Api Two: " + err.Error())
		return c.JSON(http.StatusBadRequest, _type.Response().Error("Error Occurred Mesh-UAT-Go-One Api Two: "+err.Error()))
	}

	response := _type.ResponseDto{
		Message: "Mesh-UAT-Go-Two API --TO-- Mesh-UAT-Go-One Api Two",
		Data:    resp,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}
