package utils

import (
	"net/http"
	"encoding/json"

	"github.com/menasehk13/ResturnatSystem/backend/config"
)

func RespondWithError(w http.ResponseWriter, statusCode int, errorResponse config.ErrorResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse)

	return nil
}
