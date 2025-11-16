package storage

import (
	"encoding/csv"
	"fmt"
	"os"
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
	file, err := os.OpenFile(storageDir, flag, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	record := []string{
		bucketName,
		time.Now().Format(time.RFC3339),
	}
	fmt.Println("record to write", record)
	fmt.Println(len(record))
	writer := csv.NewWriter(file)

	err = writer.Write(record)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
