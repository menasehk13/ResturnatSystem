package routes


import (
	"net/http"
)

func SetUpRoutes() *http.ServeMux{
	mux := http.NewServeMux()

	mux.Handle("/auth/",http.StripPrefix("/auth",authRoutes()))
	mux.Handle("/menu/",http.StripPrefix("/menu",menuRoutes()))
	
	return mux
}

