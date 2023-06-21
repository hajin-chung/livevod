package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	defaultLogin := "golaniyule0"
	loginPtr := flag.String("id", defaultLogin, "target streamer id")
	limitPtr := flag.Int64("limit", 10*1000*1000*1000, "video size limit in bytes")
	flag.Parse()
	dest := fmt.Sprintf("./videos/%s %s.mp4", *loginPtr, currentTime.Format("2006-01-02 15 04 05"))

	login := *loginPtr
	sig, token, err := getPlayBackToken(login)
	if err != nil {
		return
	}

	uri := getHLSUri(login, sig, token)

	c := make(chan string)
	go downloadHLSFromUri(uri, dest, *limitPtr, c)
	for {
		msg := <-c
		// fmt.Printf("%s", msg)
		if msg == "ERR" {
			break
		}
	}
}
