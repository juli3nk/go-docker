package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func ContainerList(names []string) ([]string, error) {
	var args []filters.KeyValuePair
	var result []string

	apiClient, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation(), client.WithHostFromEnv())
	if err != nil {
		return nil, err
	}
	defer apiClient.Close()

	for _, name := range names {
		args = append(args, filters.Arg("name", name))
	}

	listOpts := container.ListOptions{
		Filters: filters.NewArgs(args...),
	}

	containers, err := apiClient.ContainerList(context.Background(), listOpts)
	if err != nil {
		return nil, err
	}

	for _, ctr := range containers {
		result = append(result, ctr.Names[0][1:])
	}

	return result, nil
}

func (c *Config) ContainerRun() error {
	// c.Client.ContainerCreate(context.Background())
	// c.Client.ContainerStart(context.Background())

	return nil
}

func (c *Config) ContainerStop() error {
	return nil
}

func (c *Config) ContainerRemove() error {
	return nil
}