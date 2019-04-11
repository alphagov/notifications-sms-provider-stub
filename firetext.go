package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type FiretextResponse struct {
	Description string `json:"description"`
	Code        int    `json:"code"`
}

type FiretextCallback struct {
	Status    string `json:"status"`
	Reference string `json:"reference"`
	Mobile    string `json:"mobile"`
}

var FIRETEXT_MIN_DELAY int
var FIRETEXT_MAX_DELAY int
var FIRETEXT_CALLBACK_URL string

func init() {
	FIRETEXT_MIN_DELAY, _ = strconv.Atoi(getenv("FIRETEXT_MIN_DELAY", "100"))
	FIRETEXT_MAX_DELAY, _ = strconv.Atoi(getenv("FIRETEXT_MAX_DELAY", "1000"))
	FIRETEXT_CALLBACK_URL = getenv("FIRETEXT_CALLBACK_URL", "http://localhost:6011/notifications/sms/firetext")

	log.Printf("Firetext callback: URL %s, with delay %d-%d ms\n", FIRETEXT_CALLBACK_URL, FIRETEXT_MIN_DELAY, FIRETEXT_MAX_DELAY)
}

func FiretextEndpoint(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Firetext received message %s:  %s  \n", r.FormValue("reference"), r.FormValue("message"))

	json.NewEncoder(w).Encode(FiretextResponse{Code: 0, Description: "SMS successfully queued"})
	go FiretextSendCallback(r.FormValue("reference"), r.FormValue("to"))
}

func FiretextSendCallback(reference string, to string) {

	time.Sleep(time.Duration(FIRETEXT_MIN_DELAY+rand.Intn(FIRETEXT_MAX_DELAY-FIRETEXT_MIN_DELAY)) * time.Millisecond)

	res, err := http.PostForm(FIRETEXT_CALLBACK_URL, url.Values{
		"status":    {"0"},
		"reference": {reference},
		"mobile":    {to},
	})
	if err != nil {
		log.Printf("Firetext callback failed: %s\n", err.Error())
		return
	}

	log.Printf("Firetext callback sent for %s: %s", reference, res.Status)
}
