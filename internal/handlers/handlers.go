package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "index.html")
}

func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method does't allowed", http.StatusMethodNotAllowed)
		return
	}
	file, header, err := request.FormFile("myFile")
	if err != nil {
		log.Printf("Error retrieving file: %v", err)
		http.Error(writer, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fileCont, err := io.ReadAll(file)
	if err != nil {
		log.Printf("File read error: %v", err)
		http.Error(writer, "File read error", http.StatusInternalServerError)
		return
	}
	convCont, err := service.ConvertString(string(fileCont))
	if err != nil {
		log.Printf("Convertation error: %v", err)
		http.Error(writer, fmt.Sprintf("Convertation error: %v", err), http.StatusInternalServerError)
		return
	}
	newFN := time.Now().Format("2026-07-06_10-31-48") + filepath.Ext(header.Filename)
	if err := os.WriteFile(newFN, []byte(convCont), 0755); err != nil {
		log.Printf("Error saving file: %v", err)
		http.Error(writer, "Error saving file", http.StatusInternalServerError)
		return
	}
	log.Printf("File creation done successfully: %s", newFN)
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := fmt.Fprintf(writer, "Initial content:\n%s\nConvertation result:\n%s\nFile saved as: %s", string(fileCont), convCont, newFN); err != nil {
		log.Printf("Error sending responce: %v", err)
	}
}
