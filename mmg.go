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

var MMG_MIN_DELAY int
var MMG_MAX_DELAY int
var MMG_CALLBACK_URL string

func init() {
	MMG_MIN_DELAY, _ = strconv.Atoi(getenv("MMG_MIN_DELAY", "100"))
	MMG_MAX_DELAY, _ = strconv.Atoi(getenv("MMG_MAX_DELAY", "1000"))
	MMG_CALLBACK_URL = getenv("MMG_CALLBACK_URL", "http://localhost:6011/notifications/sms/mmg")

	log.Printf("MMG callback: URL %s, with delay %d-%d ms\n", MMG_CALLBACK_URL, MMG_MIN_DELAY, MMG_MAX_DELAY)
}

func MmgEndpoint(w http.ResponseWriter, r *http.Request) {
	var reqData MmgRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("MMG received message %s:  %s  \n", reqData.Cid, reqData.Msg)

	json.NewEncoder(w).Encode(MmgResponse{Reference: rand.Intn(100000)})
	go MmgSendCallback(reqData.Cid, reqData.MSISDN)
}

func MmgSendCallback(cid string, msisdn string) {

	time.Sleep(time.Duration(MMG_MIN_DELAY+rand.Intn(MMG_MAX_DELAY-MMG_MIN_DELAY)) * time.Millisecond)

	data := MmgCallback{Cid: cid, Status: "3", MSISDN: msisdn}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)

	res, err := http.Post(MMG_CALLBACK_URL, "application/json; charset=utf-8", buf)
	if err != nil {
		log.Printf("MMG callback failed: %s\n", err.Error())
		return
	}

	log.Printf("MMG callback sent for %s: %s", cid, res.Status)
}
