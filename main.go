package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
)

var (
	ip   = flag.String("ip", "", "Address to bind to")
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

	address := fmt.Sprintf("%v:%v", *ip, *port)

	fmt.Printf("Static file server running at %v. CTRL + C to shutdown\n",
		address)

	handler := http.FileServer(http.Dir(docroot))

	if err = http.ListenAndServe(address, handler); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
