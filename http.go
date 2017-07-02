package main

import (
	"log"
	"net/http"
	"crypto/tls"
)

const ssl_crt = "../server.crt"
const ssl_key = "../server.key"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	config := &tls.Config{MinVersion: tls.VersionTLS10}
	server := &http.Server{Addr: ":443", Handler: nil, TLSConfig: config}
	log.Fatal(server.ListenAndServeTLS(ssl_crt, ssl_key))

	log.Fatal(http.ListenAndServeTLS(":443", ssl_crt, ssl_key, nil))

	//log.Fatal(http.ListenAndServe(":80", nil))
}
