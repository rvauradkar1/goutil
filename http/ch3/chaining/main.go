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

	mux := http.NewServeMux()
	mux.HandleFunc("/json", s.json)
	mux.HandleFunc("/xml", s.xml)
	//mux = logging(mux)
	fmt.Println("Starting web server...")
	go func() { log.Fatal(http.ListenAndServeTLS(":8081", "server.crt", "server.key", mux)) }()

	fmt.Println("Starting redirect...")

	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", s.https)
	// Create a file server for static content like html, css, images, templates etc
	fileServer := http.FileServer(http.Dir("./static/"))
	// stripPrefix to remove the leading /static
	mux1.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Fatal(http.ListenAndServe(":8080", mux1))
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
	fmt.Println("ppp ", r.RequestURI)
	http.Redirect(w, r, "https://localhost:8081"+r.RequestURI, http.StatusTemporaryRedirect)
}

type company struct {
	ID   string
	Name string
	employee
}

type employee struct {
	Fname string
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging middleware....")
		next.ServeHTTP(w, r)
	})
}
