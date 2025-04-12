package docker

import (
	"context"
)

func (c *Config) Version() (string, error) {
	sv, err := c.Client.ServerVersion(context.Background())
	if err != nil {
		return "", err
	}

	return sv.Version, nil
}