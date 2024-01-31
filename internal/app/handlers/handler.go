package handlers

import (
	"movierental/internal/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetEndPoint(ctx *gin.Context)
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) (Handler, error) {
	return &handler{service: service}, nil
}

func (h *handler) GetEndPoint(ctx *gin.Context) {
	message, err := h.service.GetEndPoint()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, message)
}