package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

func main() {

	s := serve{C: company{ID: "100", Name: "XYZ", employee: employee{Fname: "First"}}}

	http.HandleFunc("/json", s.json)

	http.HandleFunc("/xml", s.xml)

	fmt.Println("Starting web server...")
	go func() { log.Fatal(http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)) }()

	fmt.Println("Starting redirect...")

	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(s.https)))
}

type serve struct {
	C company
}

func (s *serve) json(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.C)
}

func (s *serve) xml(w http.ResponseWriter, r *http.Request) {
	xml.NewEncoder(w).Encode(s.C)
}

func (s *serve) https(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	http.Redirect(w, r, "https://localhost:8081"+r.RequestURI, http.StatusMovedPermanently)
}

type company struct {
	ID   string
	Name string
	employee
}

type employee struct {
	Fname string
}
