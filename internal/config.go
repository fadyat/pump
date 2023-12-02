package internal

type Config struct {
	Driver string
}

func NewConfig() *Config {
	return &Config{
		Driver: "asana",
	}
}

func (c *Config) GetDriverOpts() map[string]any {
	switch c.Driver {
	case "asana":
		return map[string]any{"token": "token"}
	case "fs":
		return map[string]any{"file": ".pump/tasks.json"}
	}

	return map[string]any{}
}
