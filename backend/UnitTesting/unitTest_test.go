package UnitTesting

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/benzend/goalboard/pw"
	"github.com/benzend/goalboard/router"
	"github.com/benzend/goalboard/routes"
	"github.com/stretchr/testify/assert"
)

type MockDb struct{}

func TestCreateUser(t *testing.T) {
	// Create a new mock SQL database connection.
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	ctx := context.Background()
	ctxWithValueDB := context.WithValue(ctx, "db", db)
	// Define the test user credentials.
	username := "testuser"
	password := "testpass"

	hashedPassword, err := pw.HashPassword(password)
	assert.NoError(t, err)

	// Set up the expectation for the INSERT statement.
	mock.ExpectExec("INSERT INTO user_").
		WithArgs(username, hashedPassword).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Create a request body with user credentials.
	reqBody := []byte(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password))

	// Create a new HTTP request with the request body.
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	// Create a mock response recorder.
	rr := httptest.NewRecorder()

	// Call the Register handler function with the request and recorder.
	routes.Register(ctxWithValueDB, rr, req)

	router := router.NewRouter()
	router.Ctx(ctxWithValueDB)

	router.Post("/register", routes.Register)

	// Check the response status code.
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rr.Code)
	}

}
