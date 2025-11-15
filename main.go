package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	helpFlag := flag.Bool("help", false, "help flag")
	portFlag := flag.String("port", ":8000", "port flag")
	dirFlag := flag.String("dir", "main", "dir name flag")
	flag.Parse()
	if *helpFlag {
		help()
	}
	fmt.Println(*helpFlag, *dirFlag)
	http.HandleFunc("/", HandlerS3)
	http.ListenAndServe(*portFlag, nil)
}

func HandlerS3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)
	fmt.Fprintf(w, "I received your request!")
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
