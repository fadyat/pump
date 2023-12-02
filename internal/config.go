package internal

type Config struct {
	TasksFile string
}

func NewConfig() *Config {
	return &Config{
		TasksFile: ".pump/tasks.json",
	}
}
