package routes

import (
	"context"
	"net/http"
	"time"
)

func Logout(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// Clear the JWT cookie by setting its expiration time to a past date
	cookie := http.Cookie{
		Name:     "jwt_token",
		Value:    "",                         // Clear the value
		Expires:  time.Now().Add(-time.Hour), // Set expiration to a past time
		HttpOnly: true,
		Secure:   false, // Set to true if using HTTPS
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
