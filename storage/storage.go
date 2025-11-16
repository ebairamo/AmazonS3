package storage

import (
	"encoding/csv"
	"os"
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
