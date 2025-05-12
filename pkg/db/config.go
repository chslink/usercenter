package db

type Config struct {
	Data struct {
		Database struct {
			Driver string `yaml:"driver"`
			Source string `yaml:"source"`
		} `yaml:"database"`
		Redis struct {
			Addr         string `yaml:"addr"`
			Password     string `yaml:"password"`
			DB           int    `yaml:"db"`
			ReadTimeout  string `yaml:"read_timeout"`
			WriteTimeout string `yaml:"write_timeout"`
		}
	} `yaml:"data"`
}
