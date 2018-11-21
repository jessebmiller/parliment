package main

import (
	"fmt"
	"github.com/jessebmiller/parliment/api"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Parliment starting up")
	log.Fatal(http.ListenAndServe(":8080", api.GetApi()))
}
