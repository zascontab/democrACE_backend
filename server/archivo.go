package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sonderkevin/governance/application"
	"github.com/sonderkevin/governance/db"
	"github.com/sonderkevin/governance/graph/model"
	"github.com/sonderkevin/governance/graph/resolver"
	"github.com/sonderkevin/governance/infrastructure/persistence"
	"golang.org/x/crypto/bcrypt"
)

const MaxFileSize = 5000000

func UploadArchivo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if !resolver.IsAuthorized(ctx, "uploadArchivo") {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Read the archivo file from the request body
	file, header, err := r.FormFile("archivo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check the size of the file
	if header.Size > MaxFileSize {
		http.Error(w, "file is too large", http.StatusBadRequest)
		return
	}

	// Get the file extension
	if !strings.Contains(header.Filename, ".") {
		http.Error(w, "file has no extension", http.StatusBadRequest)
		return
	}
	splitted := strings.Split(header.Filename, ".")
	ext := splitted[len(splitted)-1]

	// Generate a unique ID for the archivo file
	id := strconv.FormatInt(time.Now().UnixNano(), 10)

	// Create a secure hash of the ID
	hash, err := bcrypt.GenerateFromPassword([]byte(id), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := fmt.Sprintf("%x.%s", hash, ext)

	// Create a new file in the "archivos" directory with the generated filename
	out, err := os.Create("archivos/" + filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Copy the archivo file from the request to the new file
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ar := persistence.NewArchivoRepository(db.New())
	as := application.ArchivoService{R: ar}

	// Save Archivo Record in db
	archivo, err := as.Save(r.Context(), &model.Archivo{Nombre: filename})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the ID and the hash of the saved archivo file
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(archivo.ID))
}

func ServeArchivo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if !resolver.IsAuthorized(ctx, "serveArchivo") {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	ar := persistence.NewArchivoRepository(db.New())
	as := application.ArchivoService{R: ar}

	id := mux.Vars(r)["id"]
	// Get archivo filename from db
	archivo, err := as.Get(&id)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	// Open the archivo file with the specified ID
	file, err := os.Open("archivos/" + archivo.Nombre)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the content type of the response based on the file extension
	splitted := strings.Split(archivo.Nombre, ".")
	ext := splitted[len(splitted)-1]
	w.Header().Set("Content-Type", "archivo/"+ext)

	// Serve the archivo file to the client
	io.Copy(w, file)
}

func DeleteArchivo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if !resolver.IsAuthorized(ctx, "deleteArchivo") {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	ar := persistence.NewArchivoRepository(db.New())
	as := application.ArchivoService{R: ar}

	id := mux.Vars(r)["id"]
	// Get archivo filename from db
	archivo, err := as.Get(&id)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	// Update the database
	_, err = as.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete the archivo file with the specified ID
	err = os.Remove("archivos/" + archivo.Nombre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
