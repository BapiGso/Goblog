package Smoe

import (
	"crypto/tls"
	"flag"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

type BindFlag struct {
	Domain  string
	Port    string
	SslPort string
	SslCert string
	SslKey  string
}

func (s *Smoe) BindFlag() {
	s.CommandLineArgs.Domain = *flag.String("d", "", "绑定域名，用于申请ssl证书")
	s.CommandLineArgs.Port = *flag.String("p", "80", "运行端口，默认80")
	s.CommandLineArgs.SslPort = *flag.String("tlsp", "", "tls运行端口，默认不开启")
	s.CommandLineArgs.SslCert = *flag.String("tlsc", "", "tls证书路径")
	s.CommandLineArgs.SslKey = *flag.String("tlsk", "", "tls密钥路径")
}

func (s *Smoe) Listen() {
	if s.CommandLineArgs.Domain != "" {
		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("certs"),
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
