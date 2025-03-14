package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleFile)
	http.HandleFunc("/protocol-version", HandleVersion)

	if Settings.TLS {
		log.Println("Static server listening on", Settings.Addr, "with TLS")
		err := http.ListenAndServeTLS(Settings.Addr, Settings.CertPath, Settings.KeyPath, nil)
		if err != nil {
			log.Fatal("Error launching server:", err)
		}
	} else {
		log.Println("Server listening on", Settings.Addr)
		err := http.ListenAndServe(Settings.Addr, nil)
		if err != nil {
			log.Fatal("Error launching server:", err)
		}
	}
}
