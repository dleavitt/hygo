package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	HipchatAuthToken  string `json:"hipchat_auth_token"`
	GithubAccessToken string `json:"github_access_token"`
}

func ReadConfig() (*Config, error) {
	path, _ := configFile()
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		return &config, err
	}

	return &config, nil
}

func (config *Config) Write() error {
	b, err := config.ToJSON()
	if err != nil {
		return err
	}

	path, _ := configFile()
	err = ioutil.WriteFile(path, b, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (config *Config) ToJSON() ([]byte, error) {
	b, err := json.MarshalIndent(config, "", "  ")
	return b, err
}

func configFile() (string, error) {
	// TODO: should make sure it exists
	return filepath.Join(os.Getenv("HOME"), ".hygo"), nil
}
