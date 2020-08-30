package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

var addr = os.Getenv("ADDR")

func main() {
	http.DefaultServeMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello world!"))
		writer.WriteHeader(http.StatusOK)
	})

	listener, err := net.Listen("tcp", addr)
	checkError(err)

	log.Println("listening on", listener.Addr())

	err = http.Serve(listener, http.DefaultServeMux)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
