package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const baseDir = "responses"

func serveJSON(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	subpath := strings.TrimPrefix(r.URL.Path, "/")

	// replace ? with _ to avoid issues with file names
	subpath = strings.ReplaceAll(subpath, "?", "_")

	jsonFilePath := ""

	// GET
	if r.Method == http.MethodGet {
		jsonFilePath1 := filepath.Join(baseDir, subpath+".json")
		jsonFilePath2 := filepath.Join(baseDir, subpath+".get.json")
		file1Exists := false
		file2Exists := false

		if _, err := os.Stat(jsonFilePath1); err == nil {
			file1Exists = true
		}
		if _, err := os.Stat(jsonFilePath2); err == nil {
			file2Exists = true
		}

		if file1Exists && file2Exists {
			jsonFilePath = jsonFilePath1 // or jsonFilePath2, depending on your preference
		} else if file1Exists {
			jsonFilePath = jsonFilePath1
		} else if file2Exists {
			jsonFilePath = jsonFilePath2
		} else {
			http.Error(w, "Resource not found", http.StatusNotFound)
			log.Printf("Resource not found for GET request: %s", subpath)
			return
		}
	// POST, PUT, DELETE
	} else if r.Method == http.MethodPost {
		jsonFilePath = filepath.Join(baseDir, subpath+".post.json")
		if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
			http.Error(w, "Resource not found", http.StatusNotFound)
			log.Printf("Resource not found for POST request: %s", jsonFilePath)
			return
		}
	} else if r.Method == http.MethodPut {
		jsonFilePath = filepath.Join(baseDir, subpath+".put.json")
		if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
			http.Error(w, "Resource not found", http.StatusNotFound)
			log.Printf("Resource not found for PUT request: %s", jsonFilePath)
			return
		}
	} else if r.Method == http.MethodDelete {
		jsonFilePath = filepath.Join(baseDir, subpath+".delete.json")
		if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
			http.Error(w, "Resource not found", http.StatusNotFound)
			log.Printf("Resource not found for DELETE request: %s", jsonFilePath)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("Method not allowed: %s", r.Method)
		return
	}

	// if file exists, read and return it
	file, err := os.ReadFile(jsonFilePath)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		log.Printf("Error reading file: %s, error: %v", jsonFilePath, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
	log.Printf("Successfully served file: %s", jsonFilePath)
}

func main() {
	log.Println("Starting server on :8080")
	http.HandleFunc("/", serveJSON)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
