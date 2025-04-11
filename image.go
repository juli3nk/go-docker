package docker

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func (c *Config) ImagePull(name, username, password string) error {
    ctx := context.Background()
    options := image.PullOptions{}

    if username != "" && password != "" {
        auth := base64.StdEncoding.EncodeToString([]byte(`{"username": "` + username + `", "password": "` + password + `"}`))
        options.RegistryAuth = auth
    }

    out, err := c.Client.ImagePull(ctx, name, options)
    if err != nil {
        return err
    }
    defer out.Close()

    io.Copy(io.Discard, out)

    return nil
}
