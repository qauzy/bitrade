package config

import (
	"gopkg.in/ini.v1"
)

var (
	cfg *ini.File
)

func Init(path string) (err error) {
	cfg, err = ini.Load(path)
	if err != nil {
		return
	}
	return
}

func GetConfig() *ini.File {
	return cfg
}
