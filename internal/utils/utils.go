package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	// decoder := json.NewDecoder(r.Body)
	// decoder.DisallowUnknownFields()
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteMessage(w http.ResponseWriter, status int, msg string) {
	WriteJSON(w, status, map[string]string{"message": msg})
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func Sanitize(input string) string {
	// Remove all non-alphanumeric characters except spaces and basic punctuation
	reg := regexp.MustCompile(`[^a-zA-Z0-9\s\-_,.']+`)
	safe := reg.ReplaceAllString(input, "")

	// Trim extra spaces and limit length
	safe = strings.TrimSpace(safe)
	if len(safe) > 1000 {
		safe = safe[:1000]
	}

	return safe
}
