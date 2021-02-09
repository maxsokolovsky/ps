package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"time"

	"ps/pkg/scheduler"

	"github.com/bmizerany/pat"
)

func makeHandler() http.Handler {
	mux := pat.New()

	mux.Post("/process/create", http.HandlerFunc(handler.SubmitProcess))
	mux.Post("/process/cancel/:pid", http.HandlerFunc(handler.CancelProcess))
	mux.Get("/process/:pid", http.HandlerFunc(handler.GetProcessStatus))

	return mux
}

var handler = Handler{
	s: scheduler.New(),
}

func main() {
	port := flag.String("port", ":4000", "HTTP network port")
	flag.Parse()

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *port,
		Handler:      makeHandler(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting server on", *port)
	log.Fatal(srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem"))
}
