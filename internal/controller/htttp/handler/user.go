package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/gin-gonic/gin"
)
type userRegisterRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type userLoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	
	var userRegisterRequestBody userRegisterRequestBody
	if err := c.ShouldBindJSON(&userRegisterRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidInput.Error()})
		return
	}

	user := &entity.User{
		Email: userRegisterRequestBody.Email,
		Password: userRegisterRequestBody.Password,
		Role: userRegisterRequestBody.Role,
	}

	user, err := h.service.Register(ctx, user)
	if err != nil {
		if errors.Is(err, errorsx.ErrUserExists) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrUserExists.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrPasswordHash) {
			c.JSON(http.StatusInternalServerError, gin.H{"message": errorsx.ErrInternal.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrRoleNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrRoleNotFound.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrInvalidPassword) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidPassword.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrEmailTooLong) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrEmailTooLong.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrInvalidEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidEmail.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrEmptyField) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		fmt.Println("error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": errorsx.ErrInternal.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
	
}

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var userLoginRequestBody userLoginRequestBody
	if err := c.ShouldBindJSON(&userLoginRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorsx.ErrInvalidInput.Error()})
		return
	}

	user := &entity.User{
		Email: userLoginRequestBody.Email,
		Password: userLoginRequestBody.Password,
	}

	token, err := h.service.Login(ctx, user)
	if err != nil {
		if errors.Is(err, errorsx.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": errorsx.ErrUserNotFound.Error()})
			return
		}
		if errors.Is(err, errorsx.ErrInvalidPassword) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": errorsx.ErrInvalidPassword.Error()})
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

