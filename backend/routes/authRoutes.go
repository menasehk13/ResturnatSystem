package routes

import (
	"net/http"

	"github.com/menasehk13/ResturnatSystem/backend/handler"
)


func authRoutes() http.Handler{

	max := http.NewServeMux()

	max.HandleFunc("/register",handler.RegisterUser)
	

	return max
} 
