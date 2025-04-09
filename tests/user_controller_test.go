package services

import (
	"bytes"
	"encoding/json"
	"microservicio/config"
	"microservicio/models"
	"microservicio/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var userID string

func setup() http.Handler {
	os.Setenv("MONGODB_URI", "mongodb://localhost:27017")
	os.Setenv("MONGODB_NAME", "usersdb_test") // base de pruebas
	config.ConnectDB()
	return routes.SetupRoutes()
}

func Test1_CreateUser(t *testing.T) {
	router := setup()

	user := models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "123456",
	}

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var res map[string]interface{}
	err := json.NewDecoder(rr.Body).Decode(&res)
	assert.Nil(t, err)

	insertedID, ok := res["InsertedID"].(string)
	assert.True(t, ok)
	userID = insertedID // guardamos para los próximos tests
}

func Test2_GetAllUsers(t *testing.T) {
	router := setup()

	req, _ := http.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var res []models.User
	err := json.NewDecoder(rr.Body).Decode(&res)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)
}

func Test3_GetUserByID(t *testing.T) {
	router := setup()

	req, _ := http.NewRequest("GET", "/users/"+userID, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var res models.User
	err := json.NewDecoder(rr.Body).Decode(&res)
	assert.Nil(t, err)
	assert.Equal(t, userID, res.ID.Hex())
}

func Test4_UpdateUser(t *testing.T) {
	router := setup()

	updated := models.User{
		Name:     "Updated User",
		Email:    "updated@example.com",
		Password: "abc123",
	}

	body, _ := json.Marshal(updated)
	req, _ := http.NewRequest("PUT", "/users/"+userID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func Test5_DeleteUser(t *testing.T) {
	router := setup()

	req, _ := http.NewRequest("DELETE", "/users/"+userID, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
