package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	controller_employee "service-employee/controller"
	"service-employee/model"
	"service-user/config"
	"service-user/controller"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	config.GetPostgresDatabase()

	requestBody := model.Employee{
		Name: "John Doe",
	}
	body, err := json.Marshal(requestBody)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/employee", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("access_token", "valid_access_token")

	res := httptest.NewRecorder()

	controller_employee.CreateEmployee(req)

	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestRegister(t *testing.T) {
	config.GetPostgresDatabase()

	requestBody := model.User{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, err := json.Marshal(requestBody)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	controller.Register(req)

	assert.Equal(t, http.StatusCreated, res.Code)
}
