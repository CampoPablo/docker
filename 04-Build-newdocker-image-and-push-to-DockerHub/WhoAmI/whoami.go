package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	socket := fmt.Sprintf(":%v", port)
	log.Println(fmt.Sprintf("Listening on %v...", socket))
	log.Fatal(http.ListenAndServe(socket, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()

	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintln(w, fmt.Sprintf("%v\n", host))

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Fprintln(w, err)
	}

	for _, a := range addrs {
		fmt.Fprintln(w, a)
	}

	fmt.Fprintln(w, fmt.Sprintf("\n%s", os.Getenv("SOME_VALUE")))

	fmt.Println(fmt.Sprintf("timestamp=%s method=%s route=%s protocol=%s status_code=%d", time.Now().Format("2006-01-02T15:04:05"), r.Method, r.URL, r.Proto, 200))

	fmt.Fprintln(w, fmt.Sprintf("timestamp=%s method=%s route=%s protocol=%s status_code=%d", time.Now().Format("2006-01-02T15:04:05"), r.Method, r.URL, r.Proto, 200))
}
