package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlzipper/internal/v1/zipper/errors"
	"urlzipper/internal/v1/zipper/models/dto"
	"urlzipper/internal/v1/zipper/services"
)

type URLController interface {
	Setup(engine *gin.Engine)
}

type urlController struct {
	service services.URLService
}

func NewURLController(service services.URLService) URLController {
	return &urlController{
		service: service,
	}
}

func (controller *urlController) Setup(c *gin.Engine) {
	group := c.Group("/urlzipper/v1/urls")
	{
		group.POST("", controller.Compress)
		group.GET("/:hash", controller.FindURL)
	}
}

func (controller *urlController) Compress(c *gin.Context) {
	var req dto.URLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		apiError := errors.NewBadRequestApiError("")
		_ = c.AbortWithError(http.StatusBadRequest, apiError)
		return
	}

	res, err := controller.service.Compress(&req)
	if err != nil {
		_ = c.AbortWithError(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (controller *urlController) FindURL(c *gin.Context) {
	hash := c.Param("hash")

	if hash == "" {
		apiError := errors.NewBadRequestApiError("Invalid Hash Param")
		_ = c.AbortWithError(http.StatusBadRequest, apiError)
		return
	}

	res, err := controller.service.FindURL(hash)
	if err != nil {
		_ = c.AbortWithError(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, res)
}
