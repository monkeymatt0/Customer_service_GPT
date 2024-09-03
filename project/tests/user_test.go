package tests

import (
	"bytes"
	"customer_service_gpt/api/handlers"
	"customer_service_gpt/models"
	"customer_service_gpt/services"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// MockUserService implements services.UserService interface
type MockUserService struct {
	createUserFunc     func(*models.User) error
	getUserByEmailFunc func(string) (*models.User, error)
	createSession      func(session *models.UserSession) error
}

func (m *MockUserService) CreateUser(user *models.User) error {
	return m.createUserFunc(user)
}

func (m *MockUserService) GetUserByEmail(email string) (*models.User, error) {
	return m.getUserByEmailFunc(email)
}

func (m *MockUserService) CreateSession(session *models.UserSession) error {
	return m.createSession(session)
}

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockUserService{
		createUserFunc: func(user *models.User) error {
			userService := new(services.UserService)
			return userService.CreateUser(user)
		},
	}
	handler := handlers.NewUserHandler(mockService)

	router := gin.Default()
	router.POST("/register", handler.Register)

	t.Run("Successful Registration", func(t *testing.T) {
		body := []byte(`{"email":"test@example.com","password":"password123"}`)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.Code)
		}
	})
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockUserService{
		getUserByEmailFunc: func(email string) (*models.User, error) {
			userService := new(services.UserService)
			return userService.GetUserByEmail("test@example.com")
		},
	}
	handler := handlers.NewUserHandler(mockService)

	router := gin.Default()
	router.POST("/login", handler.Login)

	t.Run("Successful Login", func(t *testing.T) {
		body := []byte(`{"email":"test@example.com","password":"password123"}`)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.Code)
		}

		var response map[string]interface{}
		json.Unmarshal(resp.Body.Bytes(), &response)
		if _, exists := response["token"]; !exists {
			t.Error("Expected token in response, but it was not present")
		}
		if _, exists := response["user_id"]; !exists {
			t.Error("Expected user_id in response, but it was not present")
		}
	})
}
