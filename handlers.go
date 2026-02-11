package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

)

func getRoot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	home := "Welcome to the API!\n"
	helloHome := map[string]string{"message": home}
	json.NewEncoder(w).Encode(helloHome)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
	w.Header().Set("Content-Type", "application/JSON")
	w.WriteHeader(http.StatusBadRequest)
	// w.Write([]byte("Error: name parameter is required"))
	errorData := map[string]string {"error":"name parameter is required"}
	json.NewEncoder(w).Encode(errorData)
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