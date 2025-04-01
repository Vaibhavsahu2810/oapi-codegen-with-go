package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type HandlerImpl struct{}

func (h *HandlerImpl) CreateLicense(c *gin.Context) {
	var license License
	if err := c.ShouldBindJSON(&license); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "License created", "license": license})
}

func (h *HandlerImpl) UpdateLicense(c *gin.Context, shortname string) {
	var license License
	if err := c.ShouldBindJSON(&license); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "License updated", "license": license})
}
