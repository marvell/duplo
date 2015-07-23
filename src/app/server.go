package main

import (
	"net/http"
	"strings"
)

type mux struct{}

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var err error

	name := strings.ToLower(r.URL.Path[1:])
	tag := strings.ToLower(r.FormValue("tag"))

	if len(name) == 0 {
		http.Error(w, "you should specify name of spec", 400)
		return
	}
	if len(tag) == 0 {
		tag = "latest"
	}

	spec := newSpec(name)
	err = spec.load()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	spec.Tag = tag

	err = pullImage(spec, false)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	stopContainer(spec)
	deleteContainer(spec)
	err = runContainer(spec, false)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func startServer(addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: &mux{},
	}

	return server.ListenAndServe()
}
