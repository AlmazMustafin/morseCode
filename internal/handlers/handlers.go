package handlers

import (
	"io"
	"log"
	"net/http"
	"://github.com"
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

		keys := []string{"file", "upload", "text", "message", "morse"}
		var file io.ReadCloser

		for _, key := range keys {
			f, _, err := r.FormFile(key)
			if err == nil {
				file = f
				break
			}
		}

		if file == nil {
			logger.Printf("Error retrieving file: no valid key found in form")
			http.Error(w, "Error retrieving file", http.StatusInternalServerError)
			return
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
