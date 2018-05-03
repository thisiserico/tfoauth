package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

func landing(w http.ResponseWriter, r *http.Request) {
	log.Println("serving landing")

	qs := url.Values{}
	qs.Add("client_id", conf.ClientID)
	qs.Add("redirect_uri", conf.RedirectURI)
	qs.Add("scope", strings.Join(scopes, "+"))

	url := fmt.Sprintf("%s/authorize?%s", conf.TypeformURL, qs.Encode())
	body := `<body><a href="%s">authorize this app</a></body>`
	body = fmt.Sprintf(body, url)

	fmt.Fprint(w, body)
}

func callback(w http.ResponseWriter, r *http.Request) {
	log.Println("running callback")

	form := url.Values{}
	form.Add("code", r.URL.Query().Get("code"))
	form.Add("client_id", conf.ClientID)
	form.Add("client_secret", conf.ClientSecret)
	form.Add("redirect_uri", conf.RedirectURI)

	url := fmt.Sprintf("%s/token", conf.TypeformURL)
	body := bytes.NewBufferString(form.Encode())

	resp, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil {
		sendFailure(w, errors.Wrap(err, "post /token"))
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sendFailure(w, errors.Wrap(err, "extracting callback body bytes"))
		return
	}

	w.Write(bodyBytes)
	log.Println("oauth flow completed, token granted")
}

func sendFailure(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	payload := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	jsPayload, _ := json.Marshal(payload)
	w.Write(jsPayload)
}
