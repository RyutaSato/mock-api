package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const baseDir = "responses"

func serveJSON(w http.ResponseWriter, r *http.Request) {
	subpath := strings.TrimPrefix(r.URL.Path, "/")
	jsonFilePath := filepath.Join(baseDir, subpath + ".json")

	if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	file, err := os.ReadFile(jsonFilePath)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}

func main() {
	http.HandleFunc("/", serveJSON)
	http.ListenAndServe(":8080", nil)
}
