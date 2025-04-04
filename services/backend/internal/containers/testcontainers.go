package containers

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"log/slog"
)

var mongoContainer testcontainers.Container
var natsContainer testcontainers.Container

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

	var err error
	mongoContainer, err = testcontainers.GenericContainer(ctx, gReq)
	if err != nil {
		return "", fmt.Errorf("could not start mongodb container: %w", err)
	}

	port, err := mongoContainer.MappedPort(ctx, "27017")
	if err != nil {
		return "", fmt.Errorf("could not get port: %w", err)
	}

	slog.Info("Started MongoDB test container", slog.Any("port", port))

	return port.Port(), nil
}

func StartNATS(ctx context.Context) (string, error) {
	cReq := testcontainers.ContainerRequest{
		Image:        "nats",
		ExposedPorts: []string{"4222/tcp"},
	}
	gReq := testcontainers.GenericContainerRequest{
		ContainerRequest: cReq,
		Started:          true,
		Reuse:            false,
	}

	var err error
	natsContainer, err = testcontainers.GenericContainer(ctx, gReq)
	if err != nil {
		return "", fmt.Errorf("could not start nats container: %w", err)
	}

	port, err := natsContainer.MappedPort(ctx, "4222")
	if err != nil {
		return "", fmt.Errorf("could not get port: %w", err)
	}

	slog.Info("Started NATS test container", slog.Any("port", port))

	return port.Port(), nil
}

func StopAllContainers(ctx context.Context) error {
	err := mongoContainer.Terminate(ctx)
	if err != nil {
		return fmt.Errorf("could not stop mongodb container: %w", err)
	}
	slog.Info("Stopped MongoDB test container")

	err = natsContainer.Terminate(ctx)
	if err != nil {
		return fmt.Errorf("could not stop nats container: %w", err)
	}
	slog.Info("Stopped NATS test container")

	return nil
}
