package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error  bool   `json:"error"`
	Message string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		Username string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err:= json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorLog.Println("Invalid json")

		payload.Error = true
		payload.Message = "Invalid json"

		out, err := json.MarshalIndent(payload, "", "\t")
		if err != nil {
			app.errorLog.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}

	app.infoLog.Println(creds.Username, creds.Password)
	payload.Error = false
	payload.Message = "Login successful"
	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		app.errorLog.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}