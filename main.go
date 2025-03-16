package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	config := ServerConfig{}
	loadConfigFile("server", &config)
	routes := Routes{}
	loadConfigFile("routes", &routes)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	for route, file := range routes.Routes {
		mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			data, err := os.ReadFile("./site/" + file)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(routes.InternalErrorMessage))
				return
			}
			w.WriteHeader(200)
			w.Write(data)
		})
	}

	log.Printf("Listening on port %v", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, mux))
}
