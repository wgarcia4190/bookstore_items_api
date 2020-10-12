package utils

import (
	"encoding/json"
	"net/http"

	"github.com/wgarcia4190/bookstore_utils_go/rest_errors"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err *rest_errors.RestErr) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)

	_ = json.NewEncoder(w).Encode(err)
}
