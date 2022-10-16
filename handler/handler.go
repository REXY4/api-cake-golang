package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}
	g := c.R.Group("/api/v1")

	g.GET("/", h.Test)
}

func (h *Handler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "succees Test",
	})
}
