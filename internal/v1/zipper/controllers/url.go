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
	}
}

func (controller *urlController) Compress(c *gin.Context) {
	var req dto.URLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
