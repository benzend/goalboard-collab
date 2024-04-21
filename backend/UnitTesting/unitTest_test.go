package UnitTesting

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/benzend/goalboard/routes"
	"github.com/benzend/goalboard/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type GoalRequestBody struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	TargetPerDay   string `json:"target_per_day"`
	LongTermTarget string `json:"long_term_target"`
}

func TestUserRegistration(t *testing.T) {
	ctx := context.Background()

	// Initialize mock DB and add it to the context
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	// Generate hashed password using bcrypt with a cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("testpass"), 10)
	if err != nil {
		t.Fatalf("failed to generate hashed password: %v", err)
	}
	username := "testuser"

	// Set up expectation for the SQL query with hashed password
	mock.ExpectExec("INSERT INTO user_ \\(username, password\\) VALUES \\(\\$1, \\$2\\)").
		WithArgs(username, string(hashedPassword)).
		WillReturnResult(sqlmock.NewResult(2, 2))
	if _, err = db.Exec("INSERT INTO user_ (username, password) VALUES ($1, $2)", username, hashedPassword); err != nil {
		return
	}
	// Create request body with username and hashed password
	reqBody := []byte(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, hashedPassword))

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))

	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}

	rr := httptest.NewRecorder()

	ctxWithDB := context.WithValue(ctx, utils.CTX_KEY_DB, db)
	routes.Register(ctxWithDB, rr, req)

	// Verify expectations and HTTP status code
	assert.NoError(t, mock.ExpectationsWereMet())
	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rr.Code)
	}
}

func TestUserLogin(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred when opening a stub database connection: %v", err)
	}
	defer db.Close()

	// Define test data
	username := "testuser"
	plainPassword := "plainpass"

	userId := int64(1)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

	if err != nil {
		t.Fatalf("Failed to generate bcrypt hash: %v", err)
	}

	reqBody, err := json.Marshal(LoginRequestBody{Username: username, Password: plainPassword})
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Mock the expected query and response
	expectedSQL := "SELECT password, id FROM user_ WHERE username = \\$1"
	mock.ExpectQuery(expectedSQL).WithArgs(username).WillReturnRows(sqlmock.NewRows([]string{"password", "id"}).AddRow(hashedPassword, userId))

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	ctxWithDB := context.WithValue(context.Background(), utils.CTX_KEY_DB, db)

	// Call the Login handler
	routes.Login(ctxWithDB, rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

}

func TestGetGoals(t *testing.T) {
	// Step 1: Set up a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred when opening a stub database connection: %v", err)
	}
	defer db.Close()

	// Step 2: Set up the expected SQL query and results
	userID := int64(1) // Assuming user ID 1
	expectedSQL := "SELECT id, name, target_per_day, long_term_target FROM goal WHERE user_id = \\$1"
	// Mock the expected query and response

	reqBody, err := json.Marshal(GoalRequestBody{ID: "1", Name: "testuser", TargetPerDay: "3 hours", LongTermTarget: "end of month"})

	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	mock.ExpectQuery(expectedSQL).WithArgs(userID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "target_per_day", "long_term_target"}).
		AddRow("1", "Goal 1", "3 hours", "end of month"))

	if err != nil {
		t.Fatalf("Failed to generate JWT token: %v", err)
	}

	if err != nil {
		return
	}
	// Create a new HTTP test recorder
	rr := httptest.NewRecorder()
	// Step 3: Create a mock authorization.

	// Create a new context with the mock database and mock authorization
	ctx := context.WithValue(context.Background(), utils.CTX_KEY_DB, db)

	req, err := http.NewRequest("GET", "/goals", bytes.NewBuffer(reqBody))

	if err != nil {
		return
	}
	// Call the GetGoals method with the mocked context
	routes.GetGoals(ctx, rr, req)

	// Verify the HTTP status code
	assert.Equal(t, http.StatusOK, rr.Code, "expected status code 200 but got different")
}

func TestUserLogout(t *testing.T) {
	// Create a mock HTTP request with any necessary headers
	req, err := http.NewRequest("GET", "/logout", nil)

	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}

	// Create a mock HTTP response writer to capture the changes made to the response
	rr := httptest.NewRecorder()

	// Call the Logout function with the mock context, response writer, and request
	ctx := context.Background()
	routes.Logout(ctx, rr, req)
	// Print out cookies in the response for debugging
	fmt.Println("Cookies in response:", rr.Result().Cookies())
	// Verify that the jwt_token cookie has been cleared
	cookies := rr.Result().Cookies()
	for _, cookie := range cookies {
		// Check if jwt_token cookie exists and its value is empty
		if cookie != nil && cookie.Value != "" {
			t.Errorf("expected jwt_token cookie to be cleared, but it still exists with value %q", cookie.Value)
		}
	}

	// Verify that the response status code is 303 (redirect)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("expected status code %d but got %d", http.StatusSeeOther, rr.Code)
	}

	// Verify that the response redirects to the "/login" page
	expectedLocation := "/login"
	actualLocation := rr.Header().Get("Location")
	if actualLocation != expectedLocation {
		t.Errorf("expected redirect to %s but got %s", expectedLocation, actualLocation)
	}
}

func TestCreateGoals(t *testing.T) {

	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}

	defer db.Close()

	userId := int64(123)

	requestBody := map[string]interface{}{
		"name":             "TestUser",
		"target_per_day":   "three housrs",
		"long_term_target": "end of month",
		"user_id":          userId,
	}

	mock.ExpectExec("INSERT INTO goal (name, target_per_day, long_term_target, user_id) VALUES ($1, $2, $3, $4) RETURNING id").
		WithArgs(requestBody["name"], requestBody["target_per_day"], requestBody["long_term_target"], requestBody["user_id"]).
		WillReturnResult(sqlmock.NewResult(4, 1))
	if _, err := db.Exec("INSERT INTO goal (name, target_per_day, long_term_target, user_id) VALUES ($1, $2, $3, $4) RETURNING id",
		requestBody["name"], requestBody["target_per_day"], requestBody["long_term_target"], requestBody["user_id"]); err != nil {
		return
	}

	reqBody, err := json.Marshal(requestBody)

	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	// Crate the HTTP request
	req, err := http.NewRequest("POST", "/goals", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}
	rr := httptest.NewRecorder()
	ctx := context.Background()
	ctxWithDB := context.WithValue(ctx, utils.CTX_KEY_DB, db)
	routes.CreateGoal(ctxWithDB, rr, req)

	// Verify expectations and HTTP status code
	assert.NoError(t, mock.ExpectationsWereMet())
	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rr.Code)
	}
}

// place holder for updating goals route.
func TestUpdateGoals(t *testing.T) {
	// Create a mock database connection
	// Step 1: Set up a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred when opening a stub database connection: %v", err)
	}
	defer db.Close()

	// Step 2: Set up the expected SQL query and results
	userID := int64(1) // Assuming user ID 1
	expectedSQL := " "
	expectedSQL1 := " "

	reqBody, err := json.Marshal(GoalRequestBody{ID: "1", Name: "testuser", TargetPerDay: "3 hours", LongTermTarget: "end of month"})

	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	mock.ExpectQuery(expectedSQL).WithArgs(userID).WillReturnRows(sqlmock.NewRows([]string{""}).
		AddRow("1", "Goal 1", "3 hours", "end of month"))

	mock.ExpectQuery(expectedSQL1).WithArgs(userID).WillReturnRows(sqlmock.NewRows([]string{""}).
		AddRow("1", "Goal 1", "3 hours", "end of month"))

	if err != nil {
		t.Fatalf("Failed to generate JWT token: %v", err)
	}

	if err != nil {
		return
	}
	// Create a new HTTP test recorder
	rr := httptest.NewRecorder()
	// Step 3: Create a mock authorization.

	// Create a new context with the mock database and mock authorization
	ctx := context.WithValue(context.Background(), utils.CTX_KEY_DB, db)

	req, err := http.NewRequest("GET", "/goals", bytes.NewBuffer(reqBody))

	if err != nil {
		return
	}
	// Call the GetGoals method with the mocked context
	routes.UpdateGoals(ctx, rr, req)

	// Verify the HTTP status code
	assert.Equal(t, http.StatusOK, rr.Code, "expected status code 200 but got different")
}
