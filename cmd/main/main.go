package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarthak-gc/file_sharing_service/pkg/config"
	"github.com/sarthak-gc/file_sharing_service/pkg/routes"
)

func main() {
	config.ConnectDB()
	router := mux.NewRouter()
	routes.Routes(router)
	fmt.Println("File sharing service in go")
	log.Fatal(http.ListenAndServe(":9090", router))
}
