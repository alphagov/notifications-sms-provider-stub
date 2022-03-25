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

type ReachResponse struct {
  // TODO
}

//type ReachCallback struct {
	//Status    string `json:"status"`
	//Reference string `json:"reference"`
	//Mobile    string `json:"mobile"`
//}

var REACH_MIN_DELAY_MS int
var REACH_MAX_DELAY_MS int
var REACH_CALLBACK_URL string
var reachClient *http.Client

func init() {
	REACH_MIN_DELAY_MS, _ = strconv.Atoi(getenv("REACH_MIN_DELAY_MS", "100"))
	REACH_MAX_DELAY_MS, _ = strconv.Atoi(getenv("REACH_MAX_DELAY_MS", "1000"))
	REACH_CALLBACK_URL = getenv("REACH_CALLBACK_URL", "http://localhost:6011/notifications/sms/reach")
	var maxConns, _ = strconv.Atoi(getenv("REACH_MAX_CONNS", "256"))

	log.Printf("Reach callback: URL %s, with delay %d-%d ms\n", REACH_CALLBACK_URL, REACH_MIN_DELAY_MS, REACH_MAX_DELAY_MS)

	reachClient = &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			MaxConnsPerHost: maxConns,
		},
	}
}

func ReachEndpoint(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Reach received message TODO")

  json.NewEncoder(w).Encode(ReachResponse{})
	//go ReachSendCallback(r.FormValue("reference"), r.FormValue("to"))
}

func ReachSendCallback(reference string, to string) {

	time.Sleep(time.Duration(REACH_MIN_DELAY_MS+rand.Intn(REACH_MAX_DELAY_MS-REACH_MIN_DELAY_MS)) * time.Millisecond)

	res, err := reachClient.PostForm(REACH_CALLBACK_URL, url.Values{
    // TODO
  })
	if err != nil {
		log.Printf("Reach callback failed: %s\n", err.Error())
		return
	}
	res.Body.Close()

	log.Printf("Reach callback sent for %s: %s", reference, res.Status)
}
