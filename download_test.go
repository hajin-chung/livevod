package main

import "testing"

func TestDownload(t *testing.T) {
	sizeLimit := int64(10 * 1000 * 1000)
	login := "golaniyule0"
	sig, token, err := getPlayBackToken(login)
	if err != nil {
		t.Fatal(err.Error())
	}

	uri := getHLSUri(login, sig, token)
	dest := "./test.mp4"
	c := make(chan string)

	go downloadHLSFromUri(uri, dest, sizeLimit, c)
	for {
		t.Log(<-c)
	}
}
