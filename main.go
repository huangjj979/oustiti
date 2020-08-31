package main

import (
	"github.com/chfanghr/oustiti/greeting"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
)

var addr = os.Getenv("ADDR")

func main() {
	router := gin.Default()

	router.GET("/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, greeting.Greet(name))
	})

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
	})

	listener, err := net.Listen("tcp", addr)
	checkError(err)

	log.Println("listening on", listener.Addr())

	checkError(http.Serve(listener, router))
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
