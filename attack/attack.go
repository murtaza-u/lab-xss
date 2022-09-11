package attack

import (
	"fmt"
	"net/http"
)

func cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set(
		"Access-Control-Allow-Methods",
		"POST, GET, OPTIONS, PUT, DELETE",
	)
	w.Header().Set(
		"Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Authorization",
	)
}

func handler(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	for _, v := range r.URL.Query() {
		fmt.Printf("%s\n", v[0])
	}
}

func Start(addr string) error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(addr, nil)
}
