package main

import (
	"net/http"
	"strings"
	"sync"
)

type mux struct{}

var lockers map[string]*sync.Mutex

func init() {
	lockers = make(map[string]*sync.Mutex)
}

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != config.AuthUser || pass != config.AuthPass {
		w.Header().Set("WWW-Authenticate", "Basic realm=duplo auth")
		http.Error(w, "access denied", http.StatusUnauthorized)
		return
	}

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

	if lockers[name] == nil {
		lockers[name] = new(sync.Mutex)
	}
	lockers[name].Lock()

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

	lockers[name].Unlock()

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
