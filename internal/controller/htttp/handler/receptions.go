package handler

import (
	"errors"
	"net/http"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
type receptionRequestBodyCreate struct {
	PvzID         string    `json:"pvzId" binding:"required"`
}

func (h *Handler) CreateReception(c *gin.Context) {
	ctx := c.Request.Context()
	var receptionRequestBody receptionRequestBodyCreate
	if err := c.ShouldBindJSON(&receptionRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidInput.Error()})
		return
	}


	pvzID, err := uuid.Parse(receptionRequestBody.PvzID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidUUID.Error()})
		return
	}

	reception := &entity.Reception{
		PvzID: pvzID,
	}

	reception, err = h.service.CreateReception(ctx, reception)
	if err != nil {
		if errors.Is(err, errorsx.ErrPVZNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrEmptyField) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, reception)
}


func (h *Handler) CloseReception(c *gin.Context) {
	ctx := c.Request.Context()
	pvzID := c.Param("pvzId")

	pvzUUID, err := uuid.Parse(pvzID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidUUID.Error()})
		return
	}	

	reception, err := h.service.CloseReception(ctx, pvzUUID)
	if err != nil {
		if errors.Is(err, errorsx.ErrReceptionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reception)
}

