package moe

import "flag"

type BindFlag struct {
	Domain  *string
	Port    *string
	SslPort *string
	SslCert *string
	SslKey  *string
	DbConf  *string
	logLvl  *string
}

func (s *Smoe) Bind() {
	s.param = &BindFlag{
		Domain:  flag.String("d", "", "绑定域名，用于申请ssl证书"),
		Port:    flag.String("p", "80", "运行端口，默认80"),
		SslPort: flag.String("tlsp", "", "tls运行端口，默认不开启"),
		SslCert: flag.String("tlsc", "", "tls证书路径"),
		SslKey:  flag.String("tlsk", "", "tls密钥路径"),
		DbConf: flag.String("dbconf", "", "数据库配置：\n"+
			"postgres:user=postgres password=your_password host=127.0.0.1 dbname=your_database sslmode=disable\n"+
			"mysql:root:your_password@tcp(127.0.0.1:3306)/your_database\n"+
			"sqlserver:server=your_server_name;user id=your_user_name;password=your_password;database=your_database"),
	}
	flag.Parse()
}
