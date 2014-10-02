package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var path = flag.String("path", "", "custom path to load the static assets")
var port = flag.Int("port", 4444, "port used by the webserver")

func main() {
	flag.Parse()

	var staticPath string
	var err error

	if *path != "" {
		staticPath = *path
	} else {
		staticPath, err = filepath.Abs("static")
		if err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		log.Fatalf("This server is trying to serve static content from: \n%s\nbut this folder doesn't exist, please create it or pass a 'path' argument when starting the server.", staticPath)
	}
	log.Printf("About to start the server on port %d\nServing content from %s", *port, staticPath)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), http.FileServer(http.Dir(staticPath))))
}
