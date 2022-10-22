package config

import (
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	ApiKey string
}

func NewConfig(file string) *Config {
	c := new(Config)

	if f, err := os.Open(file); err != nil {
		panic(err)
	} else {
		defer f.Close()
		if err := toml.NewDecoder(f).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
