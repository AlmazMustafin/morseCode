package handlers

import (
	"io"
	"log"
	"net/http"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func ServeIndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			logger.Printf("Error parsing form: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			file, _, err = r.FormFile("upload")
			if err != nil {
				logger.Printf("Error retrieving file: %v", err)
				http.Error(w, "Error retrieving file", http.StatusInternalServerError)
				return
			}
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			logger.Printf("Error reading file: %v", err)
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		input := string(data)
		result, err := service.Convert(input)
		if err != nil {
			logger.Printf("Conversion error: %v", err)
			http.Error(w, "Conversion error", http.StatusInternalServerError)
			return
		}

		logger.Printf("Successfully processed file")

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}
}
