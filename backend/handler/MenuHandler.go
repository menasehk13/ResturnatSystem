package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/menasehk13/ResturnatSystem/backend/model"
	"github.com/menasehk13/ResturnatSystem/backend/service"
)


func RegisterNewMenuItem(w http.ResponseWriter, r *http.Request){

	err := r.ParseMultipartForm(32 << 20) // max image size
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	} 
	defer file.Close()

	filename := generateFileName(handler.Filename)
	err = saveImage(file, filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	price, err := strconv.ParseFloat(r.FormValue("price"),64)
	if err != nil {
		http.Error(w, "Invalid price value", http.StatusBadRequest)
		return
	}
	catagoriesId, err:=  strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		http.Error(w, "Invalid price value", http.StatusBadRequest)
		return
	}

	var menu = model.Menu{
		MenuName: r.FormValue("menu_name"),
		Description: r.FormValue("description"),
		Price: price,
		Picture: filename,
		CategoryID:  catagoriesId,
	}
err = service.SaveMenu(menu) 
if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}

// Return a success response
w.WriteHeader(http.StatusOK)
fmt.Fprintf(w, "Menu created successfully")

}
// image display 

func ServeImage(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	staticDir := "static/menuimage"
	filePath := filepath.Join(staticDir, filename)
	fmt.Print(filePath)

	fmt.Printf("yes")
}

func getContentType(filename string) string {
	switch filepath.Ext(filename) {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	default:
		return "application/octet-stream"
	}
}


func generateFileName(originalfilename string) string {
	extension := filepath.Ext(originalfilename)
	filename := strings.TrimSuffix(originalfilename,extension)
	return fmt.Sprint("%s_%d%s",filename,time.Now().UnixNano(),extension)
}

func saveImage(file io.Reader,filename string) error {
	
	err:= os.MkdirAll("static/menuimage",os.ModePerm)
	if err != nil {
		return err
	}
	path :=filepath.Join("static/menuimage",filename)
	outputfile,err := os.Create(path)
	if err != nil {
		return err
	}

	defer outputfile.Close()

	_,err = io.Copy(outputfile,file)
	if err != nil {
		return err
	}
	return nil
}