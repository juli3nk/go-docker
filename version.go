package docker

func Version() (string, error) {
	apiClient, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation(), client.WithHostFromEnv())
	if err != nil {
		return "", err
	}
	defer apiClient.Close()

	sv, err := apiClient.ServerVersion(context.Background())
	if err != nil {
		return "", err
	}

	return sv.Version, nil
}