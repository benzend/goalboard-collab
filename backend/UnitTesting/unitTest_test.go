package UnitTesting

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benzend/goalboard/routes"
	// Import Testify/assert package
)

func TestCreateActivity(t *testing.T) {
	// Mock request body
	requestBody := []byte(`{"progress": "some_progress_value"}`)
	// Create a mock request with the request body
	req, err := http.NewRequest("POST", "/activities", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	// Create a mock response recorder
	rr := httptest.NewRecorder()

	var ctx = context.Background()
	// Hello world, the web server

	if err != nil {
		panic(err)
	}

	// Call the CreateActivity handler function
	routes.CreateActivity(ctx, rr, req)

	// Check the response status code
	if rr.Code != http.StatusCreated {
		t.Errorf("expected status code %d but got %d", http.StatusCreated, rr.Code)
	}

	// Check the response body
	expectedResponse := "Goal data inserted successfully\n"
	if rr.Body.String() != expectedResponse {
		t.Errorf("expected response body %q but got %q", expectedResponse, rr.Body.String())
	}
}
