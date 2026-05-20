package config

var RunnerUUID = ""

type CIConfig struct {
	Server     string `yaml:"server" validate:"required"`
	StorageDir string `yaml:"storage_dir" validate:"required"`
}

var CIConf CIConfig

func LoadCIConfig(cfgPath string) error {
	return loadConfig(cfgPath, &CIConf)
}
