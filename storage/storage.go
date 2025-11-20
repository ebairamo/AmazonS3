package storage

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"s3/models"
	"strconv"
	"time"
)

func BucketExists(bucketName string, storageDir string) (bool, error) {
	bucketCSV := storageDir + "/buckets.csv"
	file, err := os.Open(bucketCSV)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return false, err
	}
	for i, record := range records {
		if i == 0 {
			continue
		}
		if bucketName == record[0] {
			return true, nil
		}
	}
	return false, nil
}

func CreateBucket(bucketName string, storageDir string) error {
	err := os.Mkdir(storageDir+"/"+bucketName, 0755)
	if err != nil {
		return err
	}
	flag := os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile(storageDir+"/bucket.csv", flag, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	records := [][]string{
		{bucketName,
			time.Now().Format(time.RFC3339)},
	}
	fmt.Println("record to write", records)
	fmt.Println(len(records))
	writer := csv.NewWriter(file)
	for _, record := range records {
		err = writer.Write(record)
	}

	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}

func DeleteBucket(bucketName string, storageDir string) error {

	bucketDir := storageDir + "/" + bucketName
	err := os.RemoveAll(bucketDir)
	if err != nil {
		return err
	}

	file, err := os.Open(storageDir + "/bucket.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	ReaderCsv := csv.NewReader(file)
	records, err := ReaderCsv.ReadAll()
	if err != nil {
		return err
	}
	newRecords := [][]string{}
	for _, record := range records {
		if record[0] != bucketName {
			newRecords = append(newRecords, record)
		}
	}
	file, err = os.Create(storageDir + "/bucket.csv")
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	writer.WriteAll(newRecords)
	file.Close()
	return nil
}
func GetBucket(w http.ResponseWriter, r *http.Request, storageDir string, bucketName string) error {
	bucketCsvPath := storageDir + "/bucket.csv"
	var buckets []models.Bucket
	file, err := os.Open(bucketCsvPath)
	if err != nil {
		return err
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for i, record := range records {
		if i == 0 {
			continue
		}
		t, err := time.Parse(time.RFC3339, record[1])
		if err != nil {
			return err
		}
		Bucket := models.Bucket{
			Name:         record[0],
			CreationTime: t,
		}
		buckets = append(buckets, Bucket)
	}

	response := models.ListAllBucketsResult{
		Buckets: models.BucketsAll{
			Buckets: buckets,
		},
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	xml.NewEncoder(w).Encode(response)
	return nil
}

func IsBucketEmpty(bucketName, storageDir string) (bool, error) {
	objectPath := storageDir + "/" + bucketName + "/objects.csv"
	file, err := os.Open(objectPath)
	if err != nil {
		if os.IsNotExist(err) {
			return true, nil
		}
		return true, err
	}
	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return false, err
	}
	for i, _ := range records {
		if i == 1 {
			return false, nil
		}
	}
	// проверить существует ли objects.csv
	// если нет → return true, nil
	// если есть → прочитать, проверить количество строк
	return true, nil
}

func CreateObject(storageDir, bucketName, objectName string, body io.Reader, contentType string, size int64) error {
	fileNamePath := storageDir + "/" + bucketName + "/" + objectName
	fileName := storageDir + "/" + bucketName + "/object.csv"
	file, err := os.Create(fileNamePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, body)
	if err != nil {
		return err
	}

	if _, err = os.Stat(fileName); os.IsNotExist(err) {

		f, err := os.Create(fileName)
		if err != nil {
			return err
		}
		writer := csv.NewWriter(f)
		sizeString := strconv.Itoa(int(size))
		data := [][]string{
			{"ObjectKey", "Size", "ContentType", "LastModified"},
			{objectName, sizeString, contentType, time.Now().Format(time.RFC3339)},
		}
		for _, record := range data {
			err := writer.Write(record)
			if err != nil {
				fmt.Println("ошибка записи файла")
				return err
			}
		}
		writer.Flush()
	} else {
		sizeString := strconv.Itoa(int(size))
		flag := os.O_APPEND | os.O_WRONLY
		f, _ := os.OpenFile(fileName, flag, 0660)

		writer := csv.NewWriter(f)
		data := [][]string{
			{objectName, sizeString, contentType, time.Now().Format(time.RFC3339)},
		}
		for _, record := range data {
			err = writer.Write(record)
		}
		writer.Flush()
		fmt.Println("OPEN FILE")

	}
	return nil
}

func GetObject(w http.ResponseWriter, r *http.Request, storageDir string, bucketName string, objectName string) error {
	objectDir := storageDir + "/" + bucketName + "/" + objectName

	file, err := os.Open(objectDir)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, file)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/xml")

	return nil
}
