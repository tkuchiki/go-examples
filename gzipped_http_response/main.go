package main

import (
	"bytes"
	"compress/gzip"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body := []byte(`{"status":200}`)
	w.Header().Set("Content-Type", "application/json")
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		var buf bytes.Buffer
		gw := gzip.NewWriter(&buf)
		if _, err := gw.Write(body); err != nil {
			log.Println(err)
			return
		}
		if err := gw.Close(); err != nil {
			log.Println(err)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")

		w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
		if _, err := buf.WriteTo(w); err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Write(body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
