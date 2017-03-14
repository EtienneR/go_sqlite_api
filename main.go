package main

import (
	"log"
	"net/http"

	"github.com/EtienneR/go_sqlite_api/api"
)

func main() {
	err := http.ListenAndServe(":3000", api.Handlers())
	//err := http.ListenAndServeTLS(":3000", "certificate/localhost.pem", "certificate/localhost.key", api.Handlers())

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
