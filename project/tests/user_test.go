package tests

import (
	"bytes"
	"customer_service_gpt/api/handlers"
	"customer_service_gpt/config"
	"customer_service_gpt/db"
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
	// User related functions
	createUserFunc     func(*models.User) (uint, error)
	getUserByEmailFunc func(string) (*models.User, error)
	deleteUser         func(id uint) error

	// Session related functions
	createSession func(session *models.UserSession) error
	deleteSession func(id uint) error
	getSession    func(id uint) (*models.UserSession, error)
}

func (m *MockUserService) CreateUser(user *models.User) (uint, error) {
	return m.createUserFunc(user)
}

func (m *MockUserService) GetUserByEmail(email string) (*models.User, error) {
	return m.getUserByEmailFunc(email)
}

func (m *MockUserService) CreateSession(session *models.UserSession) error {
	return m.createSession(session)
}

func (m *MockUserService) DeleteUser(id uint) error {
	return m.deleteUser(id)
}

func (m *MockUserService) DeleteSession(id uint) error {
	return m.deleteSession(id)
}

func (m *MockUserService) GetSession(id uint) (*models.UserSession, error) {
	return m.getSession(id)
}

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db.InitDB(config.LoadConfig())

	mockService := &MockUserService{
		createUserFunc: func(user *models.User) (uint, error) {
			userService := new(services.UserService)
			return userService.CreateUser(user)
		},
		deleteUser: func(id uint) error {
			userService := new(services.UserService)
			return userService.DeleteUser(id)
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
	db.InitDB(config.LoadConfig())

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
