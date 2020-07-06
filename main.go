package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var messageArg = flag.String("message", "", "")

func message() string {
	msg := *messageArg
	if msg == "" {
		return os.Getenv("MESSAGE")
	}
	return msg
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s! %s", r.URL.Path[1:], message())
}

func main() {
	flag.Parse()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
