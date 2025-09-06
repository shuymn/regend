package config

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml/v2"
)

var (
	userHomeDir = os.UserHomeDir
)

type RedshiftConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Config struct {
	Redshift RedshiftConfig
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadTOML(filename string) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return toml.Unmarshal(b, c)
}

func LoadTOMLFilename() string {
	tomlFile := filepath.Join(xdg.ConfigHome, "/regend/regend.toml")
	if fileExists(tomlFile) {
		return tomlFile
	}

	homeDir, err := userHomeDir()
	if err == nil {
		tomlFile = filepath.Join(homeDir, ".regend.toml")
		if fileExists(tomlFile) {
			return tomlFile
		}
	}

	return ""
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil
}
