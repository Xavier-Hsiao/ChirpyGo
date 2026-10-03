package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/mail"
	"strings"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	// logged server-side for debugging
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	// user-facing error message to client
	type returnError struct {
		Error string `json:"error"`
	}
	respBody := returnError{
		Error: msg,
	}

	respondWithJSON(w, code, respBody)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(dat)
}

func validateProfane(msg string, badWords []string) string {
	words := strings.Split(msg, " ")

	for i, word := range words {
		for _, badWord := range badWords {
			if strings.ToLower(word) == badWord {
				words[i] = "****"
			}
		}
	}

	return strings.Join(words, " ")
}

func isEmailValid(email string) bool {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	// ensure that the domain name includes "."
	parts := strings.Split(addr.Address, "@")
	if len(parts) != 2 {
		return false
	}

	domain := parts[1]
	if !strings.Contains(domain, ".") {
		return false
	}

	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return false
	}

	return true
}
