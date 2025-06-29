package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarthak-gc/file_sharing_service/pkg/models"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	models.UploadFile()
	fmt.Fprint(w, "File downloaded")
}
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	models.DownloadFile(token)
	fmt.Fprint(w, "File downloaded")
}
