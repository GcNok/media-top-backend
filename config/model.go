package config

import (
	"fmt"
)

const (
	EnvTesting     = "testing"
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
	EnvLocal       = "local"
)

func ParseEnv(e string) (string, error) {
	EnvList := []string{EnvTesting, EnvDevelopment, EnvStaging, EnvProduction, EnvLocal}
	for _, env := range EnvList {
		if e == string(env) {
			return env, nil
		}
	}
	return "", fmt.Errorf("env not found: %s", e)
}

type (
	database struct {
		Name               string
		Host               string
		Port               int
		User               string
		Password           string
		MaxIdleConnections int
		MaxOpenConnections int
		LogMode            bool
	}

	cache struct {
		Host string
		Port int
	}

	server struct {
		Endpoint string
		Host     string
		Port     int
	}

	logger struct {
		Stdout struct {
			Crit  bool
			Warn  bool
			Info  bool
			Debug bool
			Dump  bool
		}
		File struct {
			Crit  bool
			Warn  bool
			Info  bool
			Debug bool
			Dump  bool
			Path  map[string]string
		}
	}
	config struct {
		Logger   logger
		Server   server
		Database database
		Cache    cache
	}

	Config   config
	Logger   logger
	Server   server
	Database database
	Cache    cache
)
