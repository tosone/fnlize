package main

import (
	"net/http"
)

// Handler is the entry point for this fission function

func Handler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, CNCF Webinar!\n"
	var _, _ = w.Write([]byte(msg))
}
