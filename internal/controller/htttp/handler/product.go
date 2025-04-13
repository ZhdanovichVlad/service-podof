package handler

import (
	"errors"
	"net/http"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type productCreateRequestBody struct {
	Type  string `json:"type"`
	PvzId string `json:"pvzId"`
}


func (h *Handler) CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()
	var productCreateRequestBody productCreateRequestBody

	if err := c.ShouldBindJSON(&productCreateRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidInput.Error()})
		return
	}

	pvzID, err := uuid.Parse(productCreateRequestBody.PvzId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidUUID.Error()})
		return
	}

	product := &entity.Product{
		Type: productCreateRequestBody.Type,
	}

	product, err = h.service.CreateProduct(ctx, product, pvzID)
	if err != nil {
		if errors.Is(err, errorsx.ErrReceptionIsClosed) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrReceptionIsClosed.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrReceptionNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrReceptionIsClosed.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrPVZNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": errorsx.ErrPVZNotFound.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrEmptyField) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}	

func (h *Handler) DeleteLastProduct(c *gin.Context) {
	ctx := c.Request.Context()
	pvzID := c.Param("pvzId")

	pvzUUID, err := uuid.Parse(pvzID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidUUID.Error()})
		return
	}
	
	err = h.service.DeleteLastProduct(ctx, pvzUUID)
	if err != nil {
		if errors.Is(err, errorsx.ErrReceptionIsClosed) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrReceptionIsClosed.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrReceptionNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrReceptionIsClosed.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrPVZNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": errorsx.ErrPVZNotFound.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}	




