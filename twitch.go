package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type StreamPlaybackToken struct {
	Value     string
	Signature string
}

type PlaybackData struct {
	StreamPlaybackAccessToken StreamPlaybackToken
}

type PlaybackTokenBody struct {
	Data PlaybackData
}

func getPlayBackToken(login string) (string, string, error) {
	client := http.Client{}

	body := fmt.Sprintf(`{"operationName":"PlaybackAccessToken","variables":{"isLive":true,"login":"%s","isVod":false,"vodID":"","playerType":"frontpage"},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"0828119ded1c13477966434e15800ff57ddacf13ba1911c129dc2200705b0712"}}}`, login)

	req, err := http.NewRequest(http.MethodPost, "https://gql.twitch.tv/gql", bytes.NewBufferString(body))
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Client-Id", "kimne78kx3ncx6brgo4mv6wki5h1ko")
	res, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	// b, err := io.ReadAll(res.Body)
	// fmt.Println(string(b))

	var tokenData PlaybackTokenBody
	json.NewDecoder(res.Body).Decode(&tokenData)

	sig := tokenData.Data.StreamPlaybackAccessToken.Signature
	token := tokenData.Data.StreamPlaybackAccessToken.Value

	return sig, token, nil
}

func getHLSUri(login string, sig string, token string) string {
	encodedToken := url.QueryEscape(token)
	url := fmt.Sprintf(`https://usher.ttvnw.net/api/channel/hls/%s.m3u8?acmb=e30%%3D&allow_source=true&fast_bread=true&player_backend=mediaplayer&playlist_include_framerate=true&reassignments_supported=true&sig=%s&supported_codecs=avc1&token=%s&cdm=wv&player_version=1.17.0`, login, sig, encodedToken)

	return url
}
