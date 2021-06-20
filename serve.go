package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// gotta use TXT instead of CNAME so that people can specify the protocol

		queryHost := "_redirect." + r.Host
		txt, err := net.LookupTXT(queryHost)
		if err != nil {
			log.Printf("ERROR: %v.\n", r.Host)
			w.WriteHeader(404) // could also use 410 Gone
			w.Header().Add("Content-Type", "text/plain")
			fmt.Fprintf(w, "To enable redirect, create a DNS TXT record at %v.\n", queryHost)
			fmt.Fprintf(w, "No response or failed TXT lookup for %v\n", queryHost)
			return
		}
		//target := ips
		log.Printf("OK: %v.\n", r.Host)
		w.WriteHeader(302)
		w.Header().Add("Location", txt[0])
	})
	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
