package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	port = flag.Int("port", 8080, "Port to listen")
	path = flag.String("path", "./", "Path served as document root.")
)

const Version = "0.1.1"

func main() {
	flag.Parse()

	docroot, err := filepath.Abs(*path)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("Static file server running at %s:%d. CTRL + C to shutdown\n", "http://locahost", *port)
	err = http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(docroot)))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
