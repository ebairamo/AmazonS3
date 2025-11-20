package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"net/http"
	"os"
	"s3/bucket"
	"s3/handlers"
)

func main() {
	helpFlag := flag.Bool("help", false, "help flag")
	portFlag := flag.String("port", ":8000", "port flag")
	dirFlag := flag.String("dir", "main", "dir name flag")
	flag.Parse()
	if *helpFlag {
		help()
	}
	_, err := os.Stat("data/" + *dirFlag)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll("data/"+*dirFlag, 0755)
			if err != nil {
				panic(err)
			}
		}
	}

	bucketPath := "data/" + *dirFlag + "/buckets.csv"
	_, err = os.Stat(bucketPath)
	if err != nil {
		if os.IsNotExist(err) {
			WriteTitle(*dirFlag)

		}
	}

	storageDir := "data/" + *dirFlag
	mux := handlers.ServerMux(storageDir)
	err = bucket.ValidateContainerName("my-bucket")
	fmt.Println(err) // должно быть nil

	http.ListenAndServe(*portFlag, mux)
}

func help() {
	fmt.Println(`$ ./triple-s --help  
Simple Storage Service.

**Usage:**
    triple-s [-port <N>] [-dir <S>]  
    triple-s --help

**Options:**
- --help     Show this screen.
- --port N   Port number
- --dir S    Path to the directory`)
}

func WriteTitle(dataDir string) {
	file, err := os.Create("data/" + dataDir + "/buckets.csv")
	if err != nil {
		fmt.Println("ошибка создания файла")
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	data := [][]string{
		{"Name", "CreationTime"},
	}
	for _, record := range data {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("ошибка записи файла")
		}
	}
	writer.Flush()
}
