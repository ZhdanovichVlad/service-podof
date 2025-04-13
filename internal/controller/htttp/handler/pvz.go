package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)	
type pvzRequestBodyCreate struct {
	Id string `json:"id"`
	RegistrationDate string `json:"registrationDate"`
	City string `json:"city"`
}


func (h *Handler) CreatePvz(c *gin.Context) {
	ctx := c.Request.Context()
	
	var pvzRequestBody pvzRequestBodyCreate
	if err := c.ShouldBindJSON(&pvzRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidInput.Error()})
		return
	}

	pvzUUID, err := uuid.Parse(pvzRequestBody.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidUUID.Error()})
		return
	}

	registrationDate, err := time.Parse(time.RFC3339, pvzRequestBody.RegistrationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidTimeFormat.Error()})
		return
	}

	pvz := &entity.Pvz{
		Id: pvzUUID,
		RegistrationDate: registrationDate,
		City: pvzRequestBody.City,
	}
	

	pvz, err = h.service.CreatePvz(ctx, pvz)
	if err != nil {
		if errors.Is(err, errorsx.ErrInvalidTimeFormat) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidTimeFormat.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrPvzExists) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrPvzExists.Error()})
			return

		}
		if errors.Is(err, errorsx.ErrCityIsNotExists) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrCityIsNotExists.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrEmptyField) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pvz)
}



func (h *Handler) GetPvzList(c *gin.Context) {
	var filter entity.Filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidInput.Error()})
		return
	}

	if err := filter.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	list, err := h.service.GetPvzList(c.Request.Context(), filter)
	if err != nil {
		if errors.Is(err, errorsx.ErrInvalidLimit) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidLimit.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}