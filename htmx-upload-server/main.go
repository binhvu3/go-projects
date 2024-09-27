package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var tmpl *template.Template

func init() {
	tmpl, _ = template.ParseGlob("templates/*.html")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/upload", UploadHandler).Methods("POST")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "upload", nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// initialize error messages slice
	var serverMessages []string

	// Parse the multipart form, 10 MB max upload size
	r.ParseMultipartForm(10 << 20)

	// Retrieve the file from form data
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		if err == http.ErrMissingFile {
			serverMessages = append(serverMessages, "No file submitted")
		} else {
			serverMessages = append(serverMessages, "Error retrieving the file")
		}

		if len(serverMessages) > 0 {
			tmpl.ExecuteTemplate(w, "messages", serverMessages)
			return
		}

	}

	defer file.Close()

	// Generate a unique filename to prevent overwriting and conflicts
	uuid, err := uuid.NewRandom()
	if err != nil {
		serverMessages = append(serverMessages, "Error generating unique identifier")
		tmpl.ExecuteTemplate(w, "messages", serverMessages)

		return
	}
	filename := uuid.String() + filepath.Ext(handler.Filename) // Append the file extension

	// Create the full path for saving the file
	filePath := filepath.Join("uploads", filename)

	// Save the file to the server
	dst, err := os.Create(filePath)
	if err != nil {
		serverMessages = append(serverMessages, "Error saving the file")
		tmpl.ExecuteTemplate(w, "messages", serverMessages)

		return
	}
	defer dst.Close()
	if _, err = io.Copy(dst, file); err != nil {
		serverMessages = append(serverMessages, "Error saving the file")
		tmpl.ExecuteTemplate(w, "messages", serverMessages)
		return
	}

	serverMessages = append(serverMessages, "File Successfully Saved")
	tmpl.ExecuteTemplate(w, "messages", serverMessages)
}
