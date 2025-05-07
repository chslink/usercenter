package config

import (
	"strings"

	"github.com/spf13/viper"
)

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

// Init 初始化配置的默认值
func (c *Config) Init() error {
	c.defString(&c.Server.Http.Addr, ":8000")
	c.defString(&c.Server.Grpc.Addr, ":9000")
	c.defString(&c.Server.Http.Timeout, "1s")
	c.defString(&c.Server.Grpc.Timeout, "1s")
	c.defString(&c.Data.Database.Driver, "mysql")
	c.defString(&c.Data.Database.Source, "root:root@tcp(127.0.0.1:3306)/test?parseTime=True&loc=Local")
	// 补充 Redis 配置的默认值
	c.defString(&c.Data.Redis.Addr, "127.0.0.1:6379")
	c.defString(&c.Data.Redis.Password, "")
	c.defInt(&c.Data.Redis.DB, 0)
	c.defString(&c.Data.Redis.ReadTimeout, "0.2s")
	c.defString(&c.Data.Redis.WriteTimeout, "0.2s")

	return nil
}

// defString 如果指针指向的字符串为空，则将其设置为指定值
func (c *Config) defString(ptr *string, val string) {
	if *ptr == "" {
		*ptr = val
	}
}

// defInt 如果指针指向的整数为 0，则将其设置为指定值
func (c *Config) defInt(ptr *int, val int) {
	if *ptr == 0 {
		*ptr = val
	}
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetEnvPrefix("uc")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
