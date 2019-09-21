package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	}
	http.HandleFunc("/h1", h1)

	j := func(w http.ResponseWriter, r *http.Request) {
		c := company{ID: "100", Name: "XYZ", employee: employee{Fname: "First"}}
		json.NewEncoder(w).Encode(c)
	}
	http.HandleFunc("/json", j)

	x := func(w http.ResponseWriter, r *http.Request) {
		c := company{ID: "100", Name: "XYZ", employee: employee{Fname: "First"}}
		xml.NewEncoder(w).Encode(c)
	}
	http.HandleFunc("/xml", x)

	fmt.Println("Starting web server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type company struct {
	ID   string
	Name string
	employee
}

type employee struct {
	Fname string
}
