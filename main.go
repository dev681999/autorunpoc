package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello world!!!")
	fmt.Fprintf(w, "hello\n")
}

func bye(w http.ResponseWriter, req *http.Request) {
	fmt.Println("bye!!!")
	fmt.Fprintf(w, "bye\n")
}

func exit(w http.ResponseWriter, req *http.Request) {
	os.Exit(-1)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bye", bye)
	http.HandleFunc("/exit", exit)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
