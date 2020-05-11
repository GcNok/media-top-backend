package config

import (
	"bytes"
	"fmt"
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var (
	configValue *config
)

const (
	osEnvParam = "env"
)

func Loads() error {
	filePath := fmt.Sprintf("/env/%s.yml", GetEnv())
	file, err := Assets.Open(filePath)
	if err != nil {
		return err
	}

	by := new(bytes.Buffer)
	_, err = io.Copy(by, file)
	if err != nil {
		return err
	}
	buf := by.Bytes()

	configValue = &config{}
	if err := yaml.Unmarshal(buf, configValue); err != nil {
		return err
	}

	return nil
}

func Reset() error {
	configValue = nil
	return os.Setenv(osEnvParam, "")
}

func SetEnv(env string) error {
	_, err := ParseEnv(env)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("set %s environment", env))
	return os.Setenv(osEnvParam, env)
}

func GetEnv() string {
	env := os.Getenv(osEnvParam)
	if env == "" {
		panic("Environment not set")
	}
	e, err := ParseEnv(env)
	if err != nil {
		panic(err.Error())
	}
	return e
}

func GetEnvValue() *config {
	if configValue == nil {
		panic("Must run Loads first")
	}
	return configValue
}
