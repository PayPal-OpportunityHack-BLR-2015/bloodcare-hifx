package app

import (
	"fmt"
)

// Envs
const (
	MODE_DEV    string = "dev"
	MODE_PROD   string = "prod"
	MODE_DEBUG  string = "debug"
	COOKIE_NAME string = "bloodcare"
)

type Config struct {
	AssetsUrl  string
	UploadPath string
}

func (c Config) String() string {
	return fmt.Sprintf(":%s:", c.AssetsUrl)
}

//NewConfig creates a config object
func NewConfig(assetsurl, uploadPath string) *Config {
	return &Config{AssetsUrl: assetsurl, UploadPath: uploadPath}
}
