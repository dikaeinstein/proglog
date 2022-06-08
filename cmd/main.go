package main

import (
	"log"

	"github.com/dikaeinstein/proglog/server"
)

func main() {
	srv := server.NewHTTPServer(":8000")
	log.Fatal(srv.ListenAndServe())
}
