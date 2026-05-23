package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type KubeConfig struct {
	Name string   `yaml:"name" validate:"required"`
	Path string   `yaml:"path" validate:"required"`
	Tags []string `yaml:"tags"`
}

type Server struct {
	ApiListenAddr  string `yaml:"api_listen_addr" validate:"required"`
	GrpcListenAddr string `yaml:"grpc_listen_addr" validate:"required"`
	MysqlDSN       string `yaml:"mysql_dsn" validate:"required"`
	RedisDSN       string `yaml:"redis_dsn" validate:"required"`
}

type OIDC struct {
	Provider     string `yaml:"provider" validate:"required"`
	ClientID     string `yaml:"client_id" validate:"required"`
	ClientSecret string `yaml:"client_secret" validate:"required"`
	RedirectURI  string `yaml:"redirect_uri" validate:"required"`
}

type config struct {
	KubeConfigs  []KubeConfig `yaml:"kubeconfig"`
	ServerConfig Server       `yaml:"server"`
	OIDCConfig   OIDC         `yaml:"oidc"`
}

var Config config

func LoadConfig(cfgPath string) error {
	return loadConfig(cfgPath, &Config)
}

func loadConfig(cfgPath string, out interface{}) error {
	configData, err := os.ReadFile(cfgPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configData, out)
	if err != nil {
		return err
	}

	return nil
}
