// Package config consolidates commandline flags for easier function calls
package config

// Config defines valid commandline args
type Config struct {
  Target string
  PostURL string
}

// New creates new config object
func New() *Config {
  return &Config{}
}
