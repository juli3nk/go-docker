package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func (c *Config) ContainerList(names []string) ([]string, error) {
	var args []filters.KeyValuePair
	var result []string

	defer c.Close()

	for _, name := range names {
		args = append(args, filters.Arg("name", name))
	}

	listOpts := container.ListOptions{
		Filters: filters.NewArgs(args...),
	}

	containers, err := c.Client.ContainerList(context.Background(), listOpts)
	if err != nil {
		return nil, err
	}

	for _, ctr := range containers {
		result = append(result, ctr.Names[0][1:])
	}

	return result, nil
}

func (c *Config) ContainerInspect(containerID string) (container.InspectResponse, error) {
	response, err := c.Client.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return container.InspectResponse{}, err
	}

	return response, nil
}

func (c *Config) ContainerRun(
	config *container.Config,
	hostConfig *container.HostConfig,
	networkingConfig *network.NetworkingConfig,
	platform *ocispec.Platform,
	containerName string,
) error {
	ctx := context.Background()

	createResponse, err := c.Client.ContainerCreate(ctx, config, hostConfig, networkingConfig, platform, containerName)
	if err != nil {
		return err
	}

	if err := c.Client.ContainerStart(ctx, createResponse.ID, container.StartOptions{}); err != nil {
		return err
	}

	return nil
}

func (c *Config) ContainerStop() error {
	return nil
}

func (c *Config) ContainerRemove() error {
	return nil
}