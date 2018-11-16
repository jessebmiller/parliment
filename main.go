package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
//	"github.com/jessebmiller/parliment/format"
)

type API struct {
	versions map[string]*http.Handler
	fallback *http.Handler
}

func (api API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
}

func noVersionHandler(w http.ResponseWriter, r *http.Request) {
	// respond with a helpful error about needing to specify an API version
	usage := `
usage: /v1/query/:queryName[/:formatter][?<params>]
       /v1/queries
       /v1/formatters
`
	io.WriteString(w, usage)
}

// v1Handler executes version one of the api
// /v1/query/:queryName[/:formatter][?<params>]
// /v1/queries
// /v1/formatters
// examples:
//     /v1/query/users-by-email?email=name@example.com
//     /v1/query/users-by-email/yaml?email=name@example.com
//     /v1/query/yield-report/csv?client_id=1234&last_n_days=7
//     /v1/queries
//     /v1/formatters
func v1Handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.Path, "/")
	io.WriteString(w, "v1 handler")
}


func main() {
	fmt.Println("Parliment starting up")
	versions := map[string]*http.Handler{"v1": &v1Handler}
	api := *&API{versions, &noVersionHandler}
	log.Fatal(http.ListenAndServe(":8080", api))
}
