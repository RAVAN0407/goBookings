package main

import (
	"fmt"
	"github.com/RAVAN0407/goBookings/pkg/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler.HelloWorld)
	fmt.Println("Starting Server on Port 9999 ")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("Unable to listen to the Port ", err)
	}
}
