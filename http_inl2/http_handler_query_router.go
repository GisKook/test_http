package http_inl2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SingleRouterProperty struct {
	ID                         uint64 `json:"RouterID"`
	PeerID                     uint64 `json:"RouterPeerID"`
	TransparentTransmissionKey uint32 `json:"RouterKey"`
	EstablishedTime            string `json:"RouterEstablishedTime"`
	RecvByteCount              uint32 `json:"RouterRecvByteCount"`
	SendByteCount              uint32 `json:"RouterSendByteCount"`
	Mode                       uint8  `json:"RouterIsTTMode"`

	WinID                         uint64 `json:"WinID"`
	WinPeerID                     uint64 `json:"WinPeerID"`
	WinTransparentTransmissionKey uint32 `json:"WinKey"`
	WinEstablishedTime            string `json:"WinEstablishedTime"`
	WinRecvByteCount              uint32 `json:"WinRecvByteCount"`
	WinSendByteCount              uint32 `json:"WinSendByteCount"`
	WinMode                       uint8  `json:"WinIsTTMode"`
}

func QueryRouter(clientid string) string {
	clients := &SingleRouterProperty{
		ID:     123,
		PeerID: 111,
		TransparentTransmissionKey: 1234,
		EstablishedTime:            "2017-08-03 09:18:11",
		RecvByteCount:              1024,
		SendByteCount:              1024,
		Mode:                       1,

		WinID:                         111,
		WinPeerID:                     123,
		WinTransparentTransmissionKey: 1234,
		WinEstablishedTime:            "2017-08-03 09:20:11",
		WinRecvByteCount:              1024,
		WinSendByteCount:              1024,
		WinMode:                       1,
	}

	response, _ := json.Marshal(clients)

	return string(response)

}

func QueryRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	fmt.Fprint(w, QueryRouter(id))
}
