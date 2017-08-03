package main

import (
	"fmt"
	"github.com/giskook/test_http/http_inl2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/inl2/get_map_list", http_inl2.GetMapsHandler)
	http.HandleFunc("/plc/query_all_routers", http_inl2.QueryAllRoutersHandler)
	http.HandleFunc("/plc/query_all_win_client", http_inl2.QueryAllWinClientHandler)
	http.HandleFunc("/plc/query_router", http_inl2.QueryRouterHandler)
	http.HandleFunc("/plc/query_win_client", http_inl2.QueryWinClientHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe :", err)
	}
	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}
