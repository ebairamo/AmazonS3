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

		parts := strings.Split(r.URL.Path, "/")
		fmt.Println(len(parts), parts[1])
		bucketName := parts[1]
		if len(parts) >= 3 {
			objectName := parts[2]
			switch {
			case r.Method == http.MethodPut:
				ObjectCreate(w, r, storageDir, bucketName, objectName)
			}

			fmt.Println(bucketName, "object", objectName)
			return
		}
		if len(parts) >= 2 {
			switch {
			case r.Method == http.MethodPut:
				BucketCreate(w, r, storageDir, bucketName)
			case r.Method == http.MethodGet:
				GetBucket(w, r, storageDir, bucketName)
			case r.Method == http.MethodDelete:
				DeleteBucket(w, r, storageDir, bucketName)
			}

		}
	})

	return mux
}

func BucketCreate(w http.ResponseWriter, r *http.Request, storageDir string, bucketName string) {

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
func ObjectCreate(w http.ResponseWriter, r *http.Request, storageDir string, bucketName string, objectName string) {
	fmt.Println(storageDir, bucketName, objectName)
}
func DeleteBucket(w http.ResponseWriter, r *http.Request, storageDir string, bucketName string) {

	isExist, err := storage.BucketExists(bucketName, storageDir)
	if err != nil {
		sendError(w, http.StatusNotFound, "NotFound", err.Error())
		return
	}
	if isExist {
		sendError(w, http.StatusNotFound, "NotFound", "Error Bucket not exist")
		return
	}
	isEmpty, err := storage.IsBucketEmpty(bucketName, storageDir)
	if err != nil {
		sendError(w, http.StatusConflict, "Conflict", err.Error())
		return
	}
	if !isEmpty {
		sendError(w, http.StatusConflict, "Conflict", "Error Bucket is not empty")
		return
	}
	err = storage.DeleteBucket(bucketName, storageDir)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}
	fmt.Println(bucketName)
}
func GetBucket(w http.ResponseWriter, r *http.Request, storageDir string, bucketName string) {
	err := storage.GetBucket(w, r, storageDir, bucketName)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "InternalError", err.Error())
		return
	}
}
