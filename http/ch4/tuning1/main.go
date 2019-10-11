package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func startWebserver() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":9090", nil))

}

func startLoadTest() {
	count := 0
	client := http.Client{}
	fmt.Print(client)
	for {
		time.Sleep(100 * time.Millisecond)
		resp, err := http.Get("http://localhost:9090/")
		if err != nil {
			panic(fmt.Sprintf("Got error: %v", err))
		}
		//io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()

		//client:=http.Client{}
		keepAliveTimeout:= 600 * time.Second
		timeout:= 2 * time.Second
		defaultTransport := &http.Transport{
	    Dial: (&net.Dialer{
                     KeepAlive: keepAliveTimeout,}
		   ).Dial, MaxIdleConns: 100, MaxIdleConnsPerHost: 100,}
client:= &http.Client{
           Transport: defaultTransport,
           Timeout:   timeout,
}

		log.Printf("Finished GET request #%v", count)
		count++
	}

}

func main() {
	go startWebserver()
	startLoadTest()
}
