package handlers

import (
	"movierental/internal/app/dto"
	"movierental/internal/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetEndPoint(ctx *gin.Context)
	GetMoviesEndPoint(ctx *gin.Context)
	GetAllMovieData(ctx *gin.Context)
	GetMovieDetail(ctx *gin.Context)
	AddMovieToCart(ctx *gin.Context)
	GetMoviesInCart(ctx *gin.Context)
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
	//ctx.BindQuery()// revisit
	movielist, err := h.service.GetAllMovieData(title, year, actors)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, movielist)
}

func (h *handler) GetMovieDetail(ctx *gin.Context) {
	param := ctx.Param("imdbID")
	movieDetail, err := h.service.GetMovieDetail(param)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, movieDetail)
}

func (h *handler) AddMovieToCart(ctx *gin.Context) {
	var addtoCartBody dto.AddtoCartRequestBody
	err := ctx.ShouldBindJSON(&addtoCartBody)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	message, err := h.service.AddMovieToCart(addtoCartBody.ImdbID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, message)
}

func (h *handler) GetMoviesInCart(ctx *gin.Context) {
	movielist, err := h.service.GetMoviesInCart()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, movielist)
}
