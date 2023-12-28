package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const port = 8002

type RequestBody struct {
	Data []int `json:"data"`
}

func logRequest(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	
	log.Printf("[%s] Received request: %v\n", time.Now().Format("2006-01-02 15:04:05"), requestBody)

	
	time.Sleep(1 * time.Second)

	
	responseTimestamp := time.Now().Format("2006-01-02 15:04:05")

	
	goResponse := fmt.Sprintf("[%s] Go service processed request\n", responseTimestamp)
	log.Print(goResponse)

	
	responseMessage := fmt.Sprintf("Request logged and processed successfully at %s", responseTimestamp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": responseMessage})
}

func main() {
	
	logFile, err := os.OpenFile("go-service.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	
	logger := log.New(io.MultiWriter(os.Stdout, logFile), "", log.LstdFlags)

	
	logger.Printf("Go Service started on port %d\n", port)

	http.HandleFunc("/log-request", logRequest)

	fmt.Printf("Go Service listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

