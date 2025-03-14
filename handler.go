package main

import (
	"log"
	"net/http"
)

var version = []byte("lumentro/static-file-server/1.0")

func HandleFile(w http.ResponseWriter, r *http.Request) {
	if path, exists := Files[r.URL.Path]; exists {
		http.ServeFile(w, r, path)
		if Settings.LogSuccess {
			log.Println("Success serving:", r.URL.Path)
		}
	} else {
		http.NotFound(w, r)
		if Settings.LogErrors {
			log.Println("Error 404 (file not found):", r.URL.Path)
		}
	}
}

func HandleVersion(w http.ResponseWriter, r *http.Request) {
	w.Write(version)
}
