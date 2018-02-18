package main

import (
	"bytes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/webhook", func(rw http.ResponseWriter, req *http.Request) {
		log.Infoln("Trigger /webhook")
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		log.Println(buf.String())
	})

	r.HandleFunc("/callback", func(rw http.ResponseWriter, req *http.Request) {
		log.Infoln("Trigger /callback")
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		log.Println(buf.String())
	})

	r.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		log.Infoln("Trigger /")
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		log.Println(buf.String())
	})

	http.ListenAndServe("0:9876", r)
}
