package main

import (
	"encoding/json"
	"log"
	"os"
)

var Settings = struct {
	TLS        bool   `json:"tls"`
	LogSuccess bool   `json:"log_success"`
	LogErrors  bool   `json:"log_errors"`
	CertPath   string `json:"cert_path"`
	KeyPath    string `json:"key_path"`
	Addr       string `json:"address"`
}{
	TLS:        false,
	LogSuccess: true,
	LogErrors:  true,
	Addr:       ":8193",
}

var Files = map[string]string{"/": "./html/index.html", "/index.html": "./html/index.html"}

func init() {
	settings, err := os.ReadFile("static-server-settings.json")
	if err != nil {
		log.Println("[Warning] Error opening settings file:", err, "Settings will be default.")
	} else {
		err = json.Unmarshal(settings, &Settings)
		if err != nil {
			log.Println("[Warning] Error decoding JSON data. Settings will be default.")

			Settings.TLS = false
			Settings.Addr = ":8193"
		}
	}

	files, err := os.ReadFile("file-list.json")
	if err != nil {
		log.Println("[Warning] Error opening filelist:", err, "Filelist will be default.")
	} else {
		err = json.Unmarshal(files, &Files)
		if err != nil {
			log.Println("[Warning] Error decoding JSON data:", err, " Filelist will be default.")
		}
	}
}
