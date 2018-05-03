package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	goji "goji.io"
	"goji.io/pat"
)

var (
	conf   config
	scopes []string
)

func init() {
	scopes = []string{"accounts:read"}
}

type config struct {
	HTTPPort    int    `envconfig:"http_port" default:"8080"`
	TypeformURL string `envconfig:"typeform_url" default:"https://api.typeform.com"`

	ClientID     string `envconfig:"client_id" required:"true"`
	ClientSecret string `envconfig:"client_secret" required:"true"`
	RedirectURI  string `envconfig:"redirect_uri" required:"true"`
}

func main() {
	if err := envconfig.Process("tfoauth", &conf); err != nil {
		log.Fatal(err)
	}

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), landing)
	mux.HandleFunc(pat.Get("/callback"), callback)

	log.Printf("serving requests in port %d", conf.HTTPPort)
	http.ListenAndServe(fmt.Sprintf(":%d", conf.HTTPPort), mux)
}
