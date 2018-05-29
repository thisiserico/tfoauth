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

const logoPath = "logo.png"

func landing(w http.ResponseWriter, r *http.Request) {
	log.Println("serving landing to user")

	qs := url.Values{}
	qs.Add("client_id", conf.ClientID)
	qs.Add("redirect_uri", conf.RedirectURI)
	qs.Add("scope", strings.Join(scopes, " "))

	log.Println("providing link to authenticate against typeform")
	url := fmt.Sprintf("%s/authorize?%s", conf.TypeformURL, qs.Encode())
	body := `<body><a href="%s">authorize this app against typeform</a></body>`
	body = fmt.Sprintf(body, url)

	fmt.Fprint(w, body)
}

func callback(w http.ResponseWriter, r *http.Request) {
	log.Println("processing callback from typeform")

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

	log.Println("oauth flow completed, token granted!")
	w.Write(bodyBytes)
}

func modifyScopes(w http.ResponseWriter, r *http.Request) {
	scopes = strings.Split(r.URL.Query().Get("scopes"), " ")
}

func serveLogo(w http.ResponseWriter, _ *http.Request) {
	if data, err := ioutil.ReadFile(string(logoPath)); err == nil {
		w.Write(data)
		return
	}

	w.WriteHeader(http.StatusNotFound)
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
