package models

import (
	"fmt"
	"time"

	"github.com/sarthak-gc/file_sharing_service/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type FileDetails struct {
	gorm.Model
	Id          string    `json:"id"`
	FileName    string    `json:"fileName"`
	StoragePath string    `json:"storagePath"`
	UploadTime  time.Time `json:"uploadTime"`
	ExpireAt    time.Time `json:"expireAt"`
	ShareToken  string    `json:"shareToken"`
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&FileDetails{})
}

func UploadFile() {
	fmt.Println("File uploaded")
}
func DownloadFile(token string) {
	fmt.Println("File downloaded")
}
