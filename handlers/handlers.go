package handlers

import (
	"fmt"
	"net/http"
	"s3/bucket"
	"s3/storage"
	"strings"
)

func ServerMux(storageDir string) *http.ServeMux {
	mux := http.NewServeMux()

	// Используем замыкание тут!
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPut:
			BucketCreate(w, r, storageDir)
		case r.Method == http.MethodGet:
			GetBucket(w, r, storageDir)
		}

	})

	return mux
}

func BucketCreate(w http.ResponseWriter, r *http.Request, storageDir string) {

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
	if isExist {
		sendError(w, http.StatusConflict, "Conflict", "Bucket is alredy exsist")
		return
	}

	err = storage.CreateBucket(bucketName, storageDir)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}
}

func GetBucket(w http.ResponseWriter, r *http.Request, storageDir string) {
	err := storage.GetBucket(w, r, storageDir)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}
}
