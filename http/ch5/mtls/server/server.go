package main

// code mostly from article https://github.com/PrakharSrivastav/tls-certificates
import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CAs, cert and key
const (
	CertPath            string = "server-cert.pem"
	KeyPath             string = "server-key.pem"
	RootCertificatePath string = "../ca-cert.pem"
)

func main() {
	// add an endpoint
	mux := http.NewServeMux()
	mux.HandleFunc("/server", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "I am a mutually authenticated TLS server, I have authenticated the client and the client has authenticated me.")
	})
	// create a certificate pool and load all the CA certificates that you
	// want to validate a client against
	clientCA, err := ioutil.ReadFile(RootCertificatePath)
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	clientCAPool := x509.NewCertPool()
	clientCAPool.AppendCertsFromPEM(clientCA)
	log.Println("ClientCA loaded")
	// configure http server with tls configuration
	s := &http.Server{
		Handler: mux,
		Addr:    ":8080",
		TLSConfig: &tls.Config{
			ClientCAs:  clientCAPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	log.Println("starting server")
	// use server.ListenAndServeTLS instead of http.ListenAndServeTLS
	log.Fatal(s.ListenAndServeTLS(CertPath, KeyPath))
}
