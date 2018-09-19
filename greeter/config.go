package greeter

// Config determines the runtime characteristics of our service.
type Config struct {
	Port     string
	Location string
}

// DefaultConfig returns a config object with the default parameters
func DefaultConfig() *Config {
	return &Config{
		Port:     "4000",
		Location: "paradise",
	}
}
