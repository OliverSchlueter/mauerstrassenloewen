package containers

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"log/slog"
)

func StartMongoDB(ctx context.Context) (string, error) {
	cReq := testcontainers.ContainerRequest{
		Image:        "mongo",
		ExposedPorts: []string{"27017/tcp"},
	}
	gReq := testcontainers.GenericContainerRequest{
		ContainerRequest: cReq,
		Started:          true,
		Reuse:            false,
	}
	container, err := testcontainers.GenericContainer(ctx, gReq)
	if err != nil {
		return "", fmt.Errorf("could not start mongodb container: %w", err)
	}

	port, err := container.MappedPort(ctx, "27017")
	if err != nil {
		return "", fmt.Errorf("could not get port: %w", err)
	}

	slog.Info("Started MongoDB", slog.Any("port", port))

	return port.Port(), nil
}
