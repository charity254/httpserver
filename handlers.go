package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

)

type User struct {
	Name string  `json:"name"`

}
type errorResponse struct {
	Error string `json:"error"`
}

func writeError( w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse{
		Error: message,
	})
}

func getRoot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	home := "Welcome to the API!\n"
	helloHome := map[string]string{"message": home}
	json.NewEncoder(w).Encode(helloHome)
}

func getGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Only POST method allowed")
	}
	defer r.Body.Close()

	var person User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	if person.Name == ""{
		writeError(w, http.StatusBadRequest, "Name is required")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":"Hello " + person.Name,
	})
}

func getHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		writeError(w, http.StatusBadRequest, "Name is required")
		return
	}

	greeting := fmt.Sprintf("Hello, %s!\n", name)
	helloData := map[string]string{"message": greeting}
	json.NewEncoder(w).Encode(helloData)
}

func  getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	healthData := map[string]string{"status": "healthy"}
	json.NewEncoder(w).Encode(healthData)
}

var startTime = time.Now()

type StatusResponse struct {
	Service string `json:"service"`
	Uptime string `json:"uptime"`
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime)

	w.Header().Set("Content-Type", "application/json")
	response := StatusResponse {
		Service:"running",
		Uptime: uptime.String(),
	}
	json.NewEncoder(w).Encode(response)
	
}