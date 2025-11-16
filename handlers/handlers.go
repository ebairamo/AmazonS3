package handlers

import (
	"fmt"
	"net/http"
	"os"
	"s3/bucket"
	"s3/storage"
	"strings"
)

func ServerMux(storageDir string) *http.ServeMux {
	mux := http.NewServeMux()

	// Используем замыкание тут!
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		BucketCreate(w, r, storageDir)
	})

	return mux
}

func HandlerS3(w http.ResponseWriter, r *http.Request, storageDir string) {
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)
	fmt.Fprintf(w, "I received your request!")
}

func BucketCreate(w http.ResponseWriter, r *http.Request, storageDir string) {
	if r.Method != http.MethodPut {
		sendError(w, http.StatusMethodNotAllowed, "InvalidStatus", "The specified method is not allowed")

		return
	}

	bucketName := strings.Trim(r.URL.Path, "/")
	fmt.Println(bucketName)
	err := bucket.ValidateContainerName(bucketName)
	if err != nil {
		sendError(w, http.StatusBadRequest, "InvalidBucketName", err.Error())
		return
	}

	isExist, err := storage.BucketExists(bucketName, storageDir)
	if err != nil {
		sendError(w, http.StatusBadRequest, "BacketNameExists", err.Error())
		return
	}
	if isExist == true {
		sendError(w, http.StatusConflict, "Conflict", "Bucket is alredy exsist")
		return
	}
	err = os.Mkdir(storageDir, 0755)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "InternalServer", err.Error())
		return
	}
}
