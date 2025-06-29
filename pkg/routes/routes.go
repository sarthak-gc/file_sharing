package routes

import (
	"github.com/gorilla/mux"
	"github.com/sarthak-gc/file_sharing_service/pkg/controllers"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/upload", controllers.UploadFile)
	router.HandleFunc("/download/{token}", controllers.DownloadFile)
}
