package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Message string `json:"Message"`
}

func CreateResponseAndLogger(w http.ResponseWriter, statusCode int, err error, message string) {
	dashes := strings.Repeat("-", 100)
	if err != nil {
		log.Printf("\nError occurred: %v\nMessage: %s\n%s", err, message, dashes)
	} else {
		log.Printf("\nSuccess: %s\n%s", message, dashes)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := Response{
		Message: message,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(errorResponse); err != nil {
		log.Printf("\nFailed to encode error response: %v\n%s", err, dashes)
	}
}
