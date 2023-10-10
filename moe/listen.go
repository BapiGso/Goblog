package moe

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
)

func (s *Smoe) Listen() {

	if *s.param.Domain != "" {
		autoTLSManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("user"),
			HostPolicy: autocert.HostWhitelist("smoe.cc", *s.param.Domain),
		}
		server := http.Server{
			Addr:    ":443",
			Handler: s.e,
			TLSConfig: &tls.Config{
				GetCertificate: autoTLSManager.GetCertificate,
				NextProtos:     []string{acme.ALPNProto},
			},
		}
		go http.ListenAndServe(":80", autoTLSManager.HTTPHandler(s.e))
		go server.ListenAndServeTLS("", "")
		fmt.Printf(banner, "=> http server started on : 80\n")
		fmt.Printf(banner, "=> https server started on : 443\n")
	}
	if *s.param.SslPort != "" {
		go http.ListenAndServeTLS(":"+*s.param.SslPort, *s.param.SslCert, *s.param.SslKey, s.e)
		fmt.Printf(banner, "=> https server started on :"+*s.param.SslPort)
	}
	http.ListenAndServe(":"+*s.param.Port, s.e)
	fmt.Printf(banner, "=> http server started on :"+*s.param.Port)
}
