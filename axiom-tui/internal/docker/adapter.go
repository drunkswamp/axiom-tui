package docker

import (
	"context"
	"strings"
	"time"

	"axiomtui/internal/domain"

	"github.com/moby/moby/client"
)

// Adapter is a Docker client adapter.
type Adapter struct {
	cli *client.Client
}

// NewAdapter creates a new adapter.
func NewAdapter() (*Adapter, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &Adapter{cli: cli}, nil
}

// Close closes the Docker client connection.
func (a *Adapter) Close() error {
	return a.cli.Close()
}

// GetServices retrieves a list of Docker containers and maps them to domain.Service.
func (a *Adapter) GetServices(ctx context.Context) ([]domain.Service, error) {
	result, err := a.cli.ContainerList(ctx, client.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var services []domain.Service
	for _, cont := range result.Items {
		var name string
		if len(cont.Names) > 0 {
			name = strings.TrimPrefix(cont.Names[0], "/")
		}

		var status domain.AxiomStatus
		switch cont.State {
		case "running":
			status = domain.StatusRunning
		case "exited":
			status = domain.StatusStopped
		default:
			status = domain.StatusDegraded
		}

		uptime := time.Duration(0)
		if cont.State == "running" {
			uptime = time.Since(time.Unix(cont.Created, 0))
		}

		service := domain.Service{
			ID:   cont.ID[:12],
			Name: name,
			Type: domain.ServiceTypeDocker,
			State: domain.ServiceState{
				Status: status,
				Uptime: uptime,
			},
		}
		services = append(services, service)
	}

	return services, nil
}
