package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/menasehk13/ResturnatSystem/backend/config"
	"github.com/menasehk13/ResturnatSystem/backend/model"
	"github.com/menasehk13/ResturnatSystem/backend/service"
	"github.com/menasehk13/ResturnatSystem/backend/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)


func RegisterUser ( w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		errorResponce := config.ErrorResponse{
			Message: "Method not allowed",
		}
		utils.RespondWithError(w, http.StatusMethodNotAllowed, errorResponce)
		return
	}
	var userfile *model.User

	err := json.NewDecoder(r.Body).Decode(&userfile)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	token, err := generateToken(userfile)
	if err != nil {
		errorResponse := config.ErrorResponse{
			Message: "Error generating token",
		}
		utils.RespondWithError(w, http.StatusInternalServerError, errorResponse)
		return
	}

	// Hash the password
	hashedPassword, err := hashPassword(userfile.Password)
	if err != nil {
		errorResponse := config.ErrorResponse{
			Message: "Error hashing password",
		}
		utils.RespondWithError(w, http.StatusInternalServerError, errorResponse)
		return
	}

	userfile.Password = hashedPassword //
	err = service.SaveUser(userfile)
	if err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}
	response := struct {
		Token string      `json:"token"`
		User  interface{} `json:"user"`
		Message string  `json:"message"`
	}{
		Token: token,
		User:  userfile,
		Message: "User Registered Successfully",
	}

	// Send the JSON response
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}


func generateToken(user *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["firstname"] = user.FirstName
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

	// Sign the token with a secret key
	// Replace "your-secret-key" with your own secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func hashPassword(password string) (string, error) {  

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}