package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"error"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
