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
	log.Fatal(http.ListenAndServe(":8080", nil))
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

type company struct {
	ID   string
	Name string
	employee
}

type employee struct {
	Fname string
}
