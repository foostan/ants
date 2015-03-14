package ants

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"os"
)

type Config struct {
	Nodes []NodeConfig `mapstructure:"nodes"`
}

type NodeConfig struct {
	Tags        []string `mapstructure:"tags"`
	Ip          string   `mapstructure:"ip"`
	SshPort     int      `mapstructure:"ssh-port"`
	SshUser     string   `mapstructure:"ssh-user"`
	SshPassword string   `mapstructure:"ssh-password"`
}

func DecodeConfig(r io.Reader) (*Config, error) {
	var raw interface{}
	var result Config
	dec := json.NewDecoder(r)
	if err := dec.Decode(&raw); err != nil {
		return nil, err
	}

	// Decode
	var md mapstructure.Metadata
	msdec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &result,
	})

	if err != nil {
		return nil, err
	}

	if err := msdec.Decode(raw); err != nil {
		return nil, err
	}

	return &result, nil
}

func Read(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading '%s': %s", path, err)
	}

	config, err := DecodeConfig(f)
	f.Close()
	if err != nil {
		return nil, fmt.Errorf("Error decoding '%s': %s", path, err)
	}

	return config, nil
}
