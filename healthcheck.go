package main

import (
	"fmt"
	"net/http"
)

func HealthcheckEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
