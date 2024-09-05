package handlers

import (
	"customer_service_gpt/models"
	"customer_service_gpt/services"
	"customer_service_gpt/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword
	id, err := h.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "id": id})
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUserByEmail(loginData.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	session := models.UserSession{
		UserID: user.ID,
		Token:  token,
	}
	if err := h.userService.CreateSession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user_id": user.ID})
}

func (h *UserHandler) Logout(c *gin.Context) {
	var userSessionID struct {
		UserID uint `json:"userID" binding:"required"`
	}
	if err := c.ShouldBindBodyWithJSON(&userSessionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id is not sent"})
		return
	}
	_, err := h.userService.GetSession(userSessionID.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}
	err2 := h.userService.DeleteSession(userSessionID.UserID)
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Something went wrong during session deletion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var userUpdate struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong parameters"})
		return
	}

	// @todo : fetch the user
	// @todo : update the field

	c.JSON(http.StatusOK, gin.H{"user": user})
}
