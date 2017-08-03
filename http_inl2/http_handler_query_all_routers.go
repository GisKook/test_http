package http_inl2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RouterProperty struct {
	RegisterID                 uint64 `json:"RegisterID"`
	PeerID                     uint64 `json:"PeerID"`
	TransparentTransmissionKey uint32 `json:"Key"`
	EstablishedTime            string `json:"EstablishedTime"`
	RecvByteCount              uint32 `json:"RecvByteCount"`
	SendByteCount              uint32 `json:"SendByteCount"`
	Mode                       uint8  `json:"IsTTMode"`
}

func QueryAllRouters() string {
	routers := make([]*RouterProperty, 0)
	routers = append(routers, &RouterProperty{
		RegisterID: 123,
		PeerID:     111,
		TransparentTransmissionKey: 1234,
		EstablishedTime:            "2017-08-03 09:18:11",
		RecvByteCount:              1024,
		SendByteCount:              1024,
		Mode:                       1,
	})

	response, _ := json.Marshal(routers)

	return string(response)
}

func QueryAllRoutersHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Fprint(w, QueryAllRouters())
}
