package UnitTesting

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/benzend/goalboard/routes"
	"github.com/benzend/goalboard/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func initializeMockDB(ctx context.Context) (context.Context, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	if err != nil {
		return nil, nil, fmt.Errorf("failed to create mock database: %v", err)
	}
	return context.WithValue(ctx, "db", db), mock, nil
}

func TestUserRegistration(t *testing.T) {
	ctx := context.Background()

	// Initialize mock DB and add it to the context
	ctx, mock, err := initializeMockDB(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Extract DB from the context
	db := ctx.Value("db").(*sql.DB)
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
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	username := "testuser"
	password := "testpass"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	userId := int64(123)
	mock.ExpectQuery("SELECT password, id FROM user_ WHERE username = ?").
		WithArgs(username).
		WillReturnRows(sqlmock.NewRows([]string{"password", "id"}).AddRow(hashedPassword, userId))

	// Prepare the request body
	requestBody := map[string]interface{}{
		"username": username,
		"password": password,
	}
	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	// Crate the HTTP request
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}
	rr := httptest.NewRecorder()
	ctx := context.Background()
	ctxWithDB := context.WithValue(ctx, utils.CTX_KEY_DB, db)
	routes.Login(ctxWithDB, rr, req)

	// Verify expectations and HTTP status code
	assert.NoError(t, mock.ExpectationsWereMet())
	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rr.Code)
	}
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
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(reqBody))
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

// func TestGetGoals(t *testing.T) {
// 	query := "SELECT id, name, target_per_day, long_term_target FROM goal WHERE user_id = $1"

// 	type Goal struct {
// 		ID             string `json:"id"`
// 		Name           string `json:"name"`
// 		TargetPerDay   string `json:"target_per_day"`
// 		LongTermTarget string `json:"long_term_target"`
// 	}
// }
