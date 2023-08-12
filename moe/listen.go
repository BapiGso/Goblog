package moe

import (
	"crypto/tls"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

func (s *Smoe) Listen() {
	if *s.CommandLineArgs.Domain != "" {
		autoTLSManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("user"),
			HostPolicy: autocert.HostWhitelist("smoe.cc", *s.CommandLineArgs.Domain),
		}
		server := http.Server{
			Addr:    ":443",
			Handler: s.e,
			TLSConfig: &tls.Config{
				GetCertificate: autoTLSManager.GetCertificate,
				NextProtos:     []string{acme.ALPNProto},
			},
		}

		go log.Fatal(http.ListenAndServe(":80", autoTLSManager.HTTPHandler(s.e)))
		log.Fatal(server.ListenAndServeTLS("", ""))
	}
	if *s.CommandLineArgs.SslPort != "" {
		if err := http.ListenAndServeTLS(":"+*s.CommandLineArgs.SslPort, *s.CommandLineArgs.SslCert, *s.CommandLineArgs.SslKey, s.e); err != nil {
			log.Fatal(err)
		} else {
			log.Printf(banner, "=> https server started on :"+*s.CommandLineArgs.SslPort)
		}
		log.Printf(banner, "=> https server started on :"+*s.CommandLineArgs.SslPort)
	}
	if err := http.ListenAndServe(":"+*s.CommandLineArgs.Port, s.e); err != nil {
		log.Fatal(err)
	} else {
		log.Printf(banner, "=> http server started on :"+*s.CommandLineArgs.Port)
	}
}
