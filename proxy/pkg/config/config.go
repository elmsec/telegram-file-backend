package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Secret string `yaml:"secret"`
	IV     string `yaml:"iv"`
	Redis  struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		DB   int    `yaml:"db"`
	} `yaml:"redis"`
}

func InitConfig() *Config {
	var configs Config
	filename, _ := filepath.Abs("config.yaml")
	yamlFile, _ := os.ReadFile(filename)
	yaml.Unmarshal(yamlFile, &configs)
	return &configs
}
