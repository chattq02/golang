package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) GetUserByID(c *gin.Context) {
	users := []string{"User 1", "User 2", "User 3"}
	c.JSON(http.StatusOK, gin.H{"users": users})
}