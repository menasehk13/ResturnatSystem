package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/menasehk13/ResturnatSystem/backend/config"
	"github.com/menasehk13/ResturnatSystem/backend/model"
	"github.com/menasehk13/ResturnatSystem/backend/service"
	"github.com/menasehk13/ResturnatSystem/backend/utils"
)


func RegisterNewMenuItem(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		errorResponce := config.ErrorResponse{
			Message: "Method not allowed",
		}
	 utils.RespondWithError(w,http.StatusBadRequest,errorResponce)
	 return
	}

	var menuitems *model.Menu

	err := json.NewDecoder(r.Body).Decode(&menuitems)
	if err != nil {
		errorResponce := config.ErrorResponse{
			Message: "Invalid Payload request",
		}
		utils.RespondWithError(w,http.StatusBadRequest,errorResponce)
		return
	}

	err = service.SaveMenu(*menuitems)
	if err != nil {
		errorResponse := config.ErrorResponse{
			Message: "Error saving menu",
		}
		utils.RespondWithError(w, http.StatusInternalServerError, errorResponse)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Menu saved successfully")

}