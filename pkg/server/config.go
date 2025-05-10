package server

type Config struct {
	Server struct {
		Http struct {
			Addr    string `yaml:"addr"`
			Timeout string `yaml:"timeout"`
		} `yaml:"http"`
		Grpc struct {
			Addr    string `yaml:"addr"`
			Timeout string `yaml:"timeout"`
		} `yaml:"grpc"`
	} `yaml:"server"`
}
