package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sarthak-gc/file_sharing_service/pkg/models"
	"github.com/sarthak-gc/file_sharing_service/pkg/utils"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fullPath, err := utils.UploadFile(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	uploadedFile := &models.FileDetails{}
	token, err := utils.GenerateToken(16)
	if err != nil {
		panic("error while generating token, try again")
	}
	utils.ParseBody(r, uploadedFile)
	uploadedFile.ShareToken = token
	uploadedFile.StoragePath = fullPath
	models.UploadFile(*uploadedFile)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":    "File uploaded successfully",
		"ShareToken": token,
	})
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	token := vars["token"]

	if token == "" {
		http.Error(w, `{"error": "Missing token"}`, http.StatusBadRequest)
		return
	}

	file := models.DownloadFile(token)
	filePath := file.StoragePath
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, `{"error": "File does not exist"}`, http.StatusNotFound)
		return
	}

	base := filepath.Base(filePath)
	underscoreIdx := strings.LastIndex(base, "_")
	dotIdx := strings.LastIndex(base, ".")
	originalName := base
	if underscoreIdx != -1 && dotIdx != -1 && underscoreIdx < dotIdx {
		originalName = base[:underscoreIdx] + base[dotIdx:]
	}
	fileName := originalName

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, filePath)
}
