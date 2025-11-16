package storage

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"s3/models"
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

	// что нужно сделать?
	// 1. Открыть файл storageDir + "/buckets.csv"
	// 2. Прочитать все строки
	// 3. Найти bucketName
	// 4. Вернуть true если нашли, false если нет
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

func GetBucket(w http.ResponseWriter, r *http.Request, storageDir string) error {
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
		Bucket := models.BucketXml{
			Name:         record[0],
			CreationDate: t,
		}
		buckets = append(buckets, Bucket)
	}
	fmt.Println(buckets)

	return nil
}
