package main

import "testing"

func TestToken(t *testing.T) {
	login := "golaniyule0"
	sig, token, err := getPlayBackToken(login)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("sig: %s, token: %s", sig, token)
}

func TestHLSUri(t *testing.T) {
	login := "nanajam777"
	sig, token, err := getPlayBackToken(login)
	if err != nil {
		t.Fatal(err.Error())
	}

	uri := getHLSUri(login, sig, token)
	t.Logf("token: %s", token)
	t.Logf("uri: %s", uri)
}
