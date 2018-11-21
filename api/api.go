package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type API struct {
	versions map[string]http.Handler
	fallback http.Handler
}

func (api API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.Path, "/")

	if len(splitPath) < 2 {
		api.fallback.ServeHTTP(w, r)
		return
	}

	version := splitPath[1]

	versionHandler, present := api.versions[version]
	if (!present) {
		api.fallback.ServeHTTP(w, r)
		return
	}

	http.StripPrefix(
		fmt.Sprintf("/%s/", version),
		versionHandler,
	).ServeHTTP(w, r)
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
// /query/:queryName[/:formatter][?<params>]
// /queries
// /formatters
// examples:
//     /query/users-by-email?email=name@example.com
//     /query/users-by-email/yaml?email=name@example.com
//     /query/yield-report/csv?client_id=1234&last_n_days=7
//     /queries
//     /formatters
func v1Handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	fmt.Println(parts)
	io.WriteString(w, "v1 handler")
}

func GetApi() API {
	var v1 http.HandlerFunc = v1Handler
	var noVersion http.HandlerFunc = noVersionHandler
	versions := map[string]http.Handler{
		"v1": http.Handler(v1),
	}
	api := API{versions, noVersion}
	return api
}
