package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/blend/go-sdk/env"
	"github.com/blend/go-sdk/r2"
	"github.com/gorilla/mux"
)

func NewServer(recport int, sendaddr string) *http.Server {
	r := mux.NewRouter()
	sendChan := make(chan string)
	go func() {
		for addr := range sendChan {
			log.Println("addr received:", addr)
			time.Sleep(2 * time.Second)
			_, err := r2.New(addr).Do()
			if err != nil {
				log.Println("error:", err.Error())
			}
		}
	}()
	// GET /status
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		log.Println("status ok")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
	// GET /ping
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received")
		sendChan <- sendaddr
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", recport),
		Handler: r,
	}
}

func main() {
	e := env.Env()
	// log.Println(e)
	log.Println("FOO =", e.Get("FOO"))
	log.Println("SEND_ADDR =", e.Get("SEND_ADDR"))
	log.Println("TRIGGER =", e.Get("TRIGGER"))
	log.Println("Hello World")

	err := e.Require("FOO", "SEND_ADDR", "TRIGGER")
	if err != nil {
		log.Fatal(err)
	}

	trig, err := strconv.ParseBool(e.Get("TRIGGER"))
	if err != nil {
		log.Fatal(err)
	}
	if trig {
		res, err := r2.New(e.Get("SEND_ADDR")).Do()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("res code:", res.StatusCode)
	} else {
		s := NewServer(4000, e.Get("SEND_ADDR"))
		var seconds int
		for seconds < 4 {
			log.Println("waiting to start server", seconds)
			time.Sleep(1 * time.Second)
			seconds++
		}
		log.Println("done starting server")
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}

}
