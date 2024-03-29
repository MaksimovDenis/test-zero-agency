package handler

import (
	logger "test-zero-agency/logs"

	"github.com/gin-gonic/gin"
)

type Err struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.Log.Println(message)
	c.AbortWithStatusJSON(statusCode, Err{Message: message})
}
