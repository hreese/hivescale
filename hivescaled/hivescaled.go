package main

import (
	"log"

	"github.com/go-ocf/go-coap"
	"github.com/pion/dtls"
)

type AppConfig struct {
	ListenAddr string
}

var (
	Config = AppConfig{
		ListenAddr: ":5688",
	}
)

func HandleHivedataV1(w coap.ResponseWriter, req *coap.Request) {
	log.Printf("Got message in handleA: path=%q: %#v from %v", req.Msg.Path(), req.Msg, req.Client.RemoteAddr())
}

func main() {
	mux := coap.NewServeMux()
	_ = mux.Handle(`/hivedata/v1`, coap.HandlerFunc(HandleHivedataV1))

	log.Printf("Starting DTLS listener on %s", Config.ListenAddr)
	log.Fatal(coap.ListenAndServeDTLS(
		"udp-dtls",
		Config.ListenAddr,
		&dtls.Config{
			Certificate: nil, // TODO
			PrivateKey:  nil, // TODO
			ClientAuth:  dtls.NoClientCert,
			//PSK:                    nil,
			//PSKIdentityHint:        nil,
		},
		mux))
}
