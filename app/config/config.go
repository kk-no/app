package config

import "github.com/kelseyhightower/envconfig"

var Conf *Config

type Config struct {
	Port string `envconfig:"PORT" default:"5001"`
}

func (c *Config) load() error {
	return envconfig.Process("", c)
}

func Load() error {
	conf := &Config{}
	if err := conf.load(); err != nil {
		return err
	}
	Conf = conf
	return nil
}
