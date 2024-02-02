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
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, message)
}

func (h *handler) GetMoviesEndPoint(ctx *gin.Context) {
	message, err := h.service.GetMoviesEndPoint()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, message)
}

func (h *handler) GetAllMovieData(ctx *gin.Context) {
	title := ctx.Query("title")
	year := ctx.Query("year")
	actors := ctx.Query("actors")
	movielist, err := h.service.GetAllMovieData(title, year, actors)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		//ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, movielist)
}
