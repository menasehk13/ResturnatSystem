package routes

import (
	"net/http"

	"github.com/menasehk13/ResturnatSystem/backend/handler"
)

func menuRoutes() http.Handler {
	max := http.NewServeMux()

	max.HandleFunc("/addNewMenu", handler.RegisterNewMenuItem)
	max.HandleFunc("/image", handler.ServeImage)
	return max
}