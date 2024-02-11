package models

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:UserName`
	Password string `json:password`
}

func (u *User) CreateUser(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch req.Method {

	case http.MethodPost:

		Id := req.PostFormValue("Id")
		Name := req.PostFormValue("fristName")
		Target := req.PostFormValue("UserName")
		Password := req.PostFormValue("password")

		values := map[string]string{
			"Name":     Name,
			"Id":       Id,
			"Target":   Target,
			"Password": Password,
		}

		jsonResponse, err := json.Marshal(values)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonResponse)

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
