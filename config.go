package main

import (
	"encoding/json"
	"log"
	"os"
)

type ServerConfig struct {
	Port string
}

type Routes struct {
	Routes               map[string]string
	MissingMessage       string
	InternalErrorMessage string
}

func loadConfigFile(file string, into any) {
	data, err := os.ReadFile("./config/" + file + ".json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, into)
	if err != nil {
		log.Fatal(err)
	}
}
