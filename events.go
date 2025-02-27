package docker

import (
	"context"

	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
)

func (c *Config) Events(filterType events.Type) (<-chan events.Message, <-chan error) {
	var opts events.ListOptions

	if filterType != "" {
		filterE := filters.NewArgs()
		filterE.Add("type", string(filterType))

		opts.Filters = filterE
	}

	return c.Client.Events(context.Background(), opts)
}
