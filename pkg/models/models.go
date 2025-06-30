package models

import (
	"time"

	"github.com/sarthak-gc/file_sharing_service/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type FileDetails struct {
	gorm.Model
	FileName    string    `gorm:"column:fileName" json:"fileName"` // by default the PascalCase will be converted to snake_case
	StoragePath string    `json:"storagePath"`
	UploadTime  time.Time `json:"uploadTime"`
	ExpireAt    time.Time `json:"expireAt"`
	ShareToken  string    `gorm:"uniqueIndex" json:"shareToken"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&FileDetails{})
}

func UploadFile(file FileDetails) {
	db.Save(&file)
}
func DownloadFile(token string) FileDetails {
	var file FileDetails
	db.Where("share_token=?", token).First(&file)
	return file
}
