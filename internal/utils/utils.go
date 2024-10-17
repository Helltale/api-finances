package handlers

import "encoding/json"

func JsonErrorResponse(message string) string {
	response := map[string]string{"error": message}
	jsonResponse, _ := json.Marshal(response)
	return string(jsonResponse)
}
