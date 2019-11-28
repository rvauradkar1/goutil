package main

// code mostly from article https://github.com/PrakharSrivastav/tls-certificates
import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	RootCertificatePath string = "../ca-cert.pem"
	ClientCertPath      string = "client-cert.pem"
	ClientKeyPath       string = "client-key.pem"
)

func main() {
	// Create a CA certificate pool for all the servers that you want to authenticate
	rootCA, err := ioutil.ReadFile(RootCertificatePath)
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	rootCAPool := x509.NewCertPool()
	rootCAPool.AppendCertsFromPEM(rootCA)
	log.Println("RootCA loaded")
	cert, err := tls.LoadX509KeyPair(ClientCertPath, ClientKeyPath)
	// configure TLS on http.Client
	c := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{
				RootCAs: rootCAPool,
				// Load clients key-pair. This will be sent to server
				Certificates: []tls.Certificate{cert},
			},
		},
	}
	request(c)
}

func request(c http.Client) {
	// prepare a request
	u := url.URL{Scheme: "https", Host: "localhost:8080", Path: "server"}
	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		log.Fatalf("request failed : %v", err)
	}
	response, err := c.Do(r)
	if err != nil {
		log.Fatalf("request failed : %v", err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("request failed : %v", err)
	}
	log.Println(string(data))
}
