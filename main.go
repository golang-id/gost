package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
)

var (
	listen = flag.String("listen", ":8080", "Address to bind to")
	path   = flag.String("path", "./", "Path served as document root.")
)

const Version = "0.1.1"

func main() {
	flag.Parse()

	docroot, err := filepath.Abs(*path)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("Static file server running at %v. CTRL + C to shutdown\n",
		*listen)

	handler := http.FileServer(http.Dir(docroot))

	if err = http.ListenAndServe(*listen, handler); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
