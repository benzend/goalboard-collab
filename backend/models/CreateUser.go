package models

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:UserName`
	Password string `json:password`
}

func (u *User) CreateUser(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch req.Method {

	case http.MethodPost:

		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		values := map[string]string{
			"Id":       u.ID,
			"Name":     u.Name,
			"UserName": u.UserName,
			"Password": u.Password,
		}

		jsonResponse, err := json.Marshal(values)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
