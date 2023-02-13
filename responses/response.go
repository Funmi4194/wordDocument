package response

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, status bool, statusCode int, message string, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")

	//send an http response header with the provided statusCode
	w.WriteHeader(statusCode)
	respone := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(respone)
}
