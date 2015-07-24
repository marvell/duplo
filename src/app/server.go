package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
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

	var err error

	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Vary", "Accept-Encoding")
	w.Header().Set("x-content-type-options", "nosniff")

	r.ParseForm()

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

	printNow(w, "pulling image")
	err = pullImage(spec, false)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	printNow(w, "stopping old container")
	stopContainer(spec)

	printNow(w, "removing old container")
	deleteContainer(spec)

	printNow(w, "running new container")
	err = runContainer(spec, false)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	lockers[name].Unlock()

	w.WriteHeader(200)
	printNow(w, "ok")
}

func startServer(addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: &mux{},
	}

	return server.ListenAndServe()
}

func printNow(w http.ResponseWriter, msg string) {
	timestamp := time.Now().Format("15:04:05")
	w.Write([]byte(fmt.Sprintf("[%s] %s\n", timestamp, msg)))
	w.(http.Flusher).Flush()
}
