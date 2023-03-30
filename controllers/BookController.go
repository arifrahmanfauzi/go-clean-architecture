package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-clean-architecture/models"
	"go-clean-architecture/usecase"
	"net/http"
	"strconv"
)

type BookController struct {
	BookUsecase usecase.BookUsecaseInterface
}

func NewBookController(book usecase.BookUsecaseInterface) *BookController {
	return &BookController{BookUsecase: book}
}

func (bc *BookController) Store(ctx *gin.Context) {
	var input = models.BookRequest{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]models.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = models.ErrorMsg{fe.Field(), models.GetErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": out})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error", "message": err.Error(),
		})
		return
	}

	res, err := bc.BookUsecase.Store(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (bc *BookController) Fetch(ctx *gin.Context) {
	res := bc.BookUsecase.Fetch()
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (bc *BookController) Update(ctx *gin.Context) {
	var input models.Book
	var id, _ = strconv.Atoi(ctx.Param("id"))
	input.ID = id
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		return
	}
	res, err := bc.BookUsecase.Update(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to update resource"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK", "data": res,
	})
}
