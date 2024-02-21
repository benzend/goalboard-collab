package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func enableCorsAgain(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func (u *User) CreateUser(w http.ResponseWriter, req *http.Request) {
	enableCorsAgain(&w)
	w.Header().Set("Content-Type", "application/json")

	switch req.Method {

	case http.MethodPost:

		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		jsonResponse, err := json.Marshal(&u)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(string(jsonResponse))

	default:
		//Replace with better error for end user
		jsonResponse, err := json.Marshal("Unable to find values")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
		return
	}

}

//TO DO ADD IN UPDATE AND DELETE REQUEST ONCE DATABASE HOOKED UP
