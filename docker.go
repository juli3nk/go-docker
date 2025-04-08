package docker

import (
	"github.com/docker/docker/client"
)

type Config struct {
	Client *client.Client
}

func NewDockerClient() (*Config, error) {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation(), client.WithHostFromEnv())
	if err != nil {
		return nil, err
	}

	return &Config{Client: cli}, nil
}

func (c *Config) Close() error {
	return c.Client.Close()
}
