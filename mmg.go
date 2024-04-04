package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type MmgRequest struct {
	ReqType string `json:"reqType"`
	MSISDN  string
	Msg     string `json:"msg"`
	Sender  string `json:"sender"`
	Cid     string `json:"cid"`
	Multi   bool   `json:"multi"`
}

type MmgResponse struct {
	Reference int
}

type MmgCallback struct {
	Status string `json:"status"`
	Cid    string `json:"CID"`
	MSISDN string
}

var MMG_MIN_DELAY_MS int
var MMG_MAX_DELAY_MS int
var MMG_CALLBACK_URL string
var mmgClient *http.Client

func init() {
	MMG_MIN_DELAY_MS, _ = strconv.Atoi(getenv("MMG_MIN_DELAY_MS", "100"))
	MMG_MAX_DELAY_MS, _ = strconv.Atoi(getenv("MMG_MAX_DELAY_MS", "1000"))
	MMG_CALLBACK_URL = getenv("MMG_CALLBACK_URL", "http://localhost:6011/notifications/sms/mmg")
	var maxConns, _ = strconv.Atoi(getenv("MMG_MAX_CONNS", "256"))

	mmgClient = &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			MaxConnsPerHost: maxConns,
		},
	}

	log.Printf("MMG callback: URL %s, with delay %d-%d ms\n", MMG_CALLBACK_URL, MMG_MIN_DELAY_MS, MMG_MAX_DELAY_MS)
}

func MmgEndpoint(w http.ResponseWriter, r *http.Request) {
	var reqData MmgRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("MMG received message %s:\n\n    %s\n\n", reqData.Cid, reqData.Msg)

	json.NewEncoder(w).Encode(MmgResponse{Reference: rand.Intn(100000)})
	go MmgSendCallback(reqData.Cid, reqData.MSISDN)
}

func MmgSendCallback(cid string, msisdn string) {

	time.Sleep(time.Duration(MMG_MIN_DELAY_MS+rand.Intn(MMG_MAX_DELAY_MS-MMG_MIN_DELAY_MS)) * time.Millisecond)

	data := MmgCallback{Cid: cid, Status: "3", MSISDN: msisdn}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)

	res, err := mmgClient.Post(MMG_CALLBACK_URL, "application/json; charset=utf-8", buf)
	if err != nil {
		log.Printf("MMG callback failed: %s\n", err.Error())
		return
	}
	res.Body.Close()

	log.Printf("MMG callback sent for %s: %s", cid, res.Status)
}
