package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Структура для обработки ошибок

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}