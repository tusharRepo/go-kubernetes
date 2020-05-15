package main

import (
	"log"
	"net/http"
	"webservice/webservicehandler"
)

func main() {
	http.HandleFunc("/getpod", webservicehandler.Getpod)
	http.HandleFunc("/namespace", webservicehandler.Namespace)
	http.HandleFunc("/", webservicehandler.DefaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
