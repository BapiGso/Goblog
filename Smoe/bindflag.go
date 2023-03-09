package Smoe

import "flag"

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
