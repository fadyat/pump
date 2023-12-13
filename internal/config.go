package internal

import "github.com/fadyat/pump/pkg"

type Config struct {
	ConfigPath string

	Driver     string         `json:"driver"`
	DriverOpts map[string]any `json:"driver_opts"`
}

func NewConfig(configPath string) (*Config, error) {
	var config Config
	if err := pkg.ReadJson(configPath, &config); err != nil {
		return nil, err
	}

	config.ConfigPath = configPath
	return &config, nil
}

func (c *Config) GetDriverOpts() map[string]any {
	return c.DriverOpts
}
