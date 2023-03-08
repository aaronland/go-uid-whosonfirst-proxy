package server

import (
	"flag"
	
	"github.com/sfomuseum/go-flags/flagset"
)

var server_uri string
var provider_uri string
var authenticator_uri string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("proxy")

	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")
	fs.StringVar(&provider_uri, "provider-uri", "random://", "A valid aaronland/go-uid.Provider URI.")
	fs.StringVar(&authenticator_uri, "authenticator_uri", "null://", "A valid sfomuseum/go-http-auth URI.")

	return fs
}

	
