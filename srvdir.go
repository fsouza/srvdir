// Copyright 2012 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"net/http"
)

var port string

func init() {
	flag.StringVar(&port, "http", ":9090", "port to be used by the server")
}

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir("")))
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
