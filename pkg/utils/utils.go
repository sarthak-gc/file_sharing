package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err == nil {
		err := json.Unmarshal([]byte(body), x)
		if err != nil {
			return
		}
	}

}
func GenerateToken(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// https://gowithore.hashnode.dev/go-up-or-go-down
func UploadFile(r *http.Request) (string, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return "", fmt.Errorf("file size is more than 10MB")
	}

	name := r.Form.Get("name")
	if name == "" {
		return "", fmt.Errorf("specify a file name please")
	}

	f, handler, err := r.FormFile("file")
	if err != nil {
		return "", fmt.Errorf("something went wrong")
	}
	defer f.Close()

	fileExtension := strings.ToLower(filepath.Ext(handler.Filename))

	path := filepath.Join(".", "files")
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + name + fileExtension

	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("something went wrong")
	}
	defer file.Close()

	_, err = io.Copy(file, f)
	if err != nil {
		return "", fmt.Errorf("something went wrong")
	}
	return fullPath, nil
}
