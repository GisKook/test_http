package http_inl2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WinClientProperty struct {
	ClientID                   uint64 `json:"ClientID"`
	PeerID                     uint64 `json:"PeerID"`
	TransparentTransmissionKey uint32 `json:"Key"`
	EstablishedTime            string `json:"EstablishedTime"`
	RecvByteCount              uint32 `json:"RecvByteCount"`
	SendByteCount              uint32 `json:"SendByteCount"`
	Mode                       uint8  `json:"IsTTMode"`
}

func QueryAllWinClients() string {
	clients := make([]*WinClientProperty, 0)
	clients = append(clients, &WinClientProperty{
		ClientID: 111,
		PeerID:   123,
		TransparentTransmissionKey: 1234,
		EstablishedTime:            "2017-08-03 09:20:11",
		RecvByteCount:              1024,
		SendByteCount:              1024,
		Mode:                       1,
	})

	response, _ := json.Marshal(clients)

	return string(response)
}

func QueryAllWinClientHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprint(w, QueryAllWinClients())
}
