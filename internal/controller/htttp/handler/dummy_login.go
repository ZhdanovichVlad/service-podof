package handler

import (
	"errors"
	"net/http"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/gin-gonic/gin"
)



func (h *Handler) DummyLogin(c *gin.Context) {
	ctx := c.Request.Context()
	
	var dummyLogin entity.DummyLogin
	if err := c.ShouldBindJSON(&dummyLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidInput.Error()})
		return
	}

	token, err := h.service.DummyLogin(ctx, &dummyLogin)
	if err != nil {
		if errors.Is(err, errorsx.ErrRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrEmptyField) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": errorsx.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusOK, token.Token)
}


