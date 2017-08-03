package http_inl2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SingleWinClientProperty struct {
	ClientID                   uint64 `json:"WinID"`
	PeerID                     uint64 `json:"WinPeerID"`
	TransparentTransmissionKey uint32 `json:"WinKey"`
	EstablishedTime            string `json:"WinEstablishedTime"`
	RecvByteCount              uint32 `json:"WinRecvByteCount"`
	SendByteCount              uint32 `json:"WinSendByteCount"`
	Mode                       uint8  `json:"WinIsTTMode"`

	RouterID                         uint64 `json:"RouterID"`
	RouterPeerID                     uint64 `json:"RouterPeerID"`
	RouterTransparentTransmissionKey uint32 `json:"RouterKey"`
	RouterEstablishedTime            string `json:"RouterEstablishedTime"`
	RouterRecvByteCount              uint32 `json:"RouterRecvByteCount"`
	RouterSendByteCount              uint32 `json:"RouterSendByteCount"`
	RouterMode                       uint8  `json:"RouterIsTTMode"`
}

func QueryWinClient(clientid string) string {
	clients := &SingleWinClientProperty{
		ClientID: 111,
		PeerID:   123,
		TransparentTransmissionKey: 1234,
		EstablishedTime:            "2017-08-03 09:18:11",
		RecvByteCount:              1024,
		SendByteCount:              1024,
		Mode:                       1,

		RouterID:                         123,
		RouterPeerID:                     111,
		RouterTransparentTransmissionKey: 1234,
		RouterEstablishedTime:            "2017-08-03 09:20:11",
		RouterRecvByteCount:              1024,
		RouterSendByteCount:              1024,
		RouterMode:                       1,
	}

	response, _ := json.Marshal(clients)

	return string(response)
}

func QueryWinClientHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clientid := r.Form.Get("id")
	fmt.Fprint(w, QueryWinClient(clientid))
}
