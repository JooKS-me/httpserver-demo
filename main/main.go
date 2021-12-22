package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/pprof"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err := http.ListenAndServe(":80", mux)

	if err != nil {
		fmt.Println(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get request for '/'")
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("Hello %s \n", user))
	} else {
		io.WriteString(w, fmt.Sprintf("Hello stranger\n"))
	}

	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s: %s\n", k, v))
	}
}
