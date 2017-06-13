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

	err := http.ListenAndServe(":8876", nil)
	if err != nil {
		log.Fatal("ListenAndServe :", err)
	}
	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}
