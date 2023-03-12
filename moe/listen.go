package moe

import (
	"crypto/tls"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

func (s *Smoe) Listen() {
	if s.CommandLineArgs.Domain != "" {
		autoTLSManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("user"),
			HostPolicy: autocert.HostWhitelist("smoe.cc", s.CommandLineArgs.Domain),
		}
		server := http.Server{
			Addr:    ":443",
			Handler: s.E,
			TLSConfig: &tls.Config{
				GetCertificate: autoTLSManager.GetCertificate,
				NextProtos:     []string{acme.ALPNProto},
			},
		}

		go log.Fatal(http.ListenAndServe(":80", autoTLSManager.HTTPHandler(s.E)))
		log.Fatal(server.ListenAndServeTLS("", ""))
	}
	if s.CommandLineArgs.SslPort != "" {
		log.Printf(banner, "=> https server started on :"+s.CommandLineArgs.SslPort)
		log.Fatal(http.ListenAndServeTLS(":"+s.CommandLineArgs.SslPort, s.CommandLineArgs.SslCert, s.CommandLineArgs.SslKey, s.E))
	}
	log.Printf(banner, "=> http server started on :"+s.CommandLineArgs.Port)
	log.Fatal(http.ListenAndServe(":"+s.CommandLineArgs.Port, s.E))
}
