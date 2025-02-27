package compose

import (
	"os"
	"os/exec"
)

const (
	dockerComposeFile = ".devcontainer/compose.yaml"
	envFile           = ".compose.env"
)

func Up() error {
	app, err := exec.LookPath("docker")
	if err != nil {
		return err
	}

	cmd := exec.Command(
		app,
		"compose",
		"--env-file",
		envFile,
		"--file",
		dockerComposeFile,
		"up",
		"-d",
	)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Exec() error {
	app, err := exec.LookPath("docker")
	if err != nil {
		return err
	}

	cmd := exec.Command(
		app,
		"compose",
		"--env-file",
		envFile,
		"--file",
		dockerComposeFile,
		"exec",
		"workspace",
		"bash",
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
