package main

import (
	"log"
	"net/http"
)

const ssl_crt = "../server.crt"
const ssl_key = "../server.key"

func main() {
	http.Handle("/w", http.FileServer(http.Dir("webroot")))

	log.Fatal(http.ListenAndServeTLS(":443", ssl_crt, ssl_key, nil))
	//log.Fatal(http.ListenAndServe(":80", nil))
}
