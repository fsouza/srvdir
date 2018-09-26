// Copyright 2012 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	var headerList listflag
	flag.Var(&headerList, "H", "list of headers to add to response. Same format as in curl")
	port := flag.String("http", ":9090", "port to be used by the server")
	flag.Parse()

	headers, err := parseHeaders(headerList.v)
	if err != nil {
		log.Fatalf("failed to parse headers: %v", err)
	}

	handler := corsMiddleware(http.FileServer(http.Dir("")))
	http.Handle("/", headersMiddleware(headers, handler))
	log.Printf("Running at http://127.0.0.1%s", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal(err)
	}
}

func parseHeaders(values []string) (http.Header, error) {
	header := make(http.Header)
	for _, value := range values {
		parts := strings.SplitN(value, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid header %q. Should be specified in the format Key: Value", value)
		}
		header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}
	return header, nil
}

func headersMiddleware(header http.Header, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for name, values := range header {
			for _, value := range values {
				w.Header().Add(name, value)
			}
		}
		h.ServeHTTP(w, r)
	})
}

func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, api_key, Authorization")
		h.ServeHTTP(w, r)
	})
}
