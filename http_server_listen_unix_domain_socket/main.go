package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"status":200"}`)
}

func main() {
	sock := "/tmp/app.sock"

	listener, err := net.Listen("unix", sock)
	if err != nil {
		log.Fatal(err)
	}
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(listener net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		listener.Close()
		os.Exit(0)
	}(listener, sigc)

	m := http.NewServeMux()
	m.Handle("/", http.HandlerFunc(handler))
	server := http.Server{
		Handler: m,
	}

	server.Serve(listener)
}
