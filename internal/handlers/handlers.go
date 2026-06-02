package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
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

		file, handler, err := r.FormFile("file")
		if err != nil {
			logger.Printf("Error retrieving file: %v", err)
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

		timestamp := time.Now().UTC().String()
		ext := filepath.Ext(handler.Filename)
		outputFilename := timestamp + ext

		outFile, err := os.Create(outputFilename)
		if err != nil {
			logger.Printf("Error creating output file: %v", err)
			http.Error(w, "Error creating output file", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		_, err = outFile.WriteString(result)
		if err != nil {
			logger.Printf("Error writing to output file: %v", err)
			http.Error(w, "Error writing to output file", http.StatusInternalServerError)
			return
		}

		logger.Printf("Successfully processed file %s, saved result to %s", handler.Filename, outputFilename)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(result))
	}
}
