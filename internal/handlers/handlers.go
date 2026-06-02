package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Не удалось получить файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	convertedText, err := service.ConvertData(string(fileBytes))
	if err != nil {
		http.Error(w, "Ошибка обработки данных", http.StatusInternalServerError)
		return
	}

	timeStr := time.Now().UTC().Format("2006-01-02_15-04-05")
	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".txt"
	}
	newFileName := timeStr + "_result" + ext

	out, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "Не удалось создать локальный файл", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = out.WriteString(convertedText)
	if err != nil {
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(convertedText))
}
