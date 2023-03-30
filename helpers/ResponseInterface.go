package helpers

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/models"
)

func APIResponse(ctx *gin.Context, Message string, statusCode int, Method string, Data interface{}) {
	jsonResponse := models.ModelResponses{
		StatusCode: statusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}
	if statusCode >= 400 {
		ctx.AbortWithStatusJSON(statusCode, jsonResponse)
	} else {
		ctx.JSON(statusCode, jsonResponse)
	}
}
