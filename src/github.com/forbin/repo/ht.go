package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {

	port := flag.String("port", "8080", "http port")
	path := flag.String("docroot", ".", "document root")
	flag.Parse()

	fmt.Println("Start http service")
	fmt.Println("  port :", *port)
	fmt.Println("  path :", *path)
	http.Handle("/", http.FileServer(http.Dir(*path)))
	http.ListenAndServe(":"+*port, nil)
}
