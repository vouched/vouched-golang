package client

// Config structure
type Config struct {
	URL, Key string
}

// NewConfig create a new config
func NewConfig() *Config {
	url := getEnv("VOUCHED_SERVER", "https://verify.vouched.id/graphql")
	return &Config{URL: url}
}
