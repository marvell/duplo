package main

import (
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type mux struct{}

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.SetOutput(w)
	log.SetLevel(log.WarnLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})

	r.ParseForm()

	var err error

	name := strings.ToLower(r.URL.Path[1:])
	tag := strings.ToLower(r.FormValue("tag"))

	if len(name) == 0 {
		w.WriteHeader(500)
		log.Errorln("you should specify name of spec")
		return
	}
	if len(tag) == 0 {
		tag = "latest"
	}

	spec := newSpec(name)
	err = spec.load()
	if err != nil {
		w.WriteHeader(500)
		log.Errorln(err)
		return
	}

	spec.Tag = tag

	err = pullImage(spec, false)
	if err != nil {
		w.WriteHeader(500)
		log.Errorln(err)
		return
	}
	stopContainer(spec)
	deleteContainer(spec)
	err = runContainer(spec, false)
	if err != nil {
		w.WriteHeader(500)
		log.Errorln(err)
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
