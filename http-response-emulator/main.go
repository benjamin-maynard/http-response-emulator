package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

func main() {

	// Get the Cloud Run $PORT variable from the environment
	listenPort := os.Getenv("PORT")

	// Define the request handler for the default / route
	http.HandleFunc("/", httpDefaultHandler)
	http.ListenAndServe(":"+listenPort, nil)

}

type logEntry struct {
	ResponseCode int    `json:"responseCode"`
	RequestURL   string `json:"requestURL"`
	RequestIP    string `json:"requestIP"`
}

type httpResponse struct {
	HTTPCode int    `json:"httpCode"`
	Message  string `json:"message"`
}

func jsonLogRequest(responseCode int, requestURL string, requestIP string) {

	// Define a new JSON Encoder
	var jsonEncoder = json.NewEncoder(os.Stdout)

	// Craft a log entry
	logEntry := logEntry{ResponseCode: responseCode, RequestURL: requestURL, RequestIP: requestIP}

	// Return the Response
	if err := jsonEncoder.Encode(&logEntry); err != nil {
		panic(err)
	}

}

func httpDefaultHandler(w http.ResponseWriter, r *http.Request) {

	// Set the Response Type to Application JSON
	w.Header().Set("Content-Type", "application/json")

	// Get the Version of the Service
	responseCode := os.Getenv("RESPONSE_CODE")

	// Convert to int
	responseCodeInt, err := strconv.Atoi(responseCode)

	// Define a new JSON Encoder
	var jsonEncoder = json.NewEncoder(w)

	if err != nil {
		// Define a response
		response := httpResponse{HTTPCode: http.StatusInternalServerError, Message: "Error converting RESPONSE_CODE environment variable to string"}
		// Return the Response
		w.WriteHeader(http.StatusInternalServerError)
		if err := jsonEncoder.Encode(&response); err != nil {
			panic(err)
		}
		return
	}

	// Define a response
	response := httpResponse{HTTPCode: responseCodeInt, Message: "Responded with configured HTTP Code"}

	// Log the Response
	jsonLogRequest(responseCodeInt, r.URL.Path, r.RemoteAddr)

	// Return the Response
	w.WriteHeader(responseCodeInt)
	if err := jsonEncoder.Encode(&response); err != nil {
		panic(err)
	}
	return

}
