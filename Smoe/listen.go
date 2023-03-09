package Smoe

import (
	"crypto/tls"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

func (s *Smoe) Listen() {
	if s.CommandLineArgs.Domain != "" {
		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("user"),
			HostPolicy: autocert.HostWhitelist("smoe.cc", s.CommandLineArgs.Domain),
		}
		server := &http.Server{
			Addr:    ":443",
			Handler: s.E,
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}

		go log.Fatal(http.ListenAndServe(":80", certManager.HTTPHandler(nil)))
		log.Fatal(server.ListenAndServeTLS("", ""))
	}
	if s.CommandLineArgs.SslPort != "" {
		log.Fatal(http.ListenAndServeTLS(":"+s.CommandLineArgs.SslPort, s.CommandLineArgs.SslCert, s.CommandLineArgs.SslKey, s.E))
	}
	log.Fatal(http.ListenAndServe(":"+s.CommandLineArgs.Port, s.E))
}
