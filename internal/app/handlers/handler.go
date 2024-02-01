package handlers

import (
	"movierental/internal/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetEndPoint(ctx *gin.Context)
	GetMoviesEndPoint(ctx *gin.Context)
	GetAllMovieData(ctx *gin.Context)
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{service: service}
}

func (h *handler) GetEndPoint(ctx *gin.Context) {
	message, err := h.service.GetEndPoint()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, message)
}

func (h *handler) GetMoviesEndPoint(ctx *gin.Context) {
	message, err := h.service.GetMoviesEndPoint()
    if err!= nil {
        ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
        return
    }
    ctx.AbortWithStatusJSON(http.StatusOK, message)
}

func (h *handler) GetAllMovieData(ctx *gin.Context) {
	message, err := h.service.GetAllMovieData()
    if err!= nil {
        ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
        return
    }
    ctx.AbortWithStatusJSON(http.StatusOK, message)
}
