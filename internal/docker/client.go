package docker

import (
	"context"
	"strings"
	"time"

	"axiomtui/internal/domain"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Client is a wrapper around the official Docker API client.
type Client struct {
	api client.APIClient
}

// NewClient creates and initializes a new Docker client.
func NewClient() (*Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &Client{api: cli}, nil
}

// ListContainers gets a list of local containers and maps them to the unified model.
func (c *Client) ListContainers() ([]domain.Service, error) {
	containers, err := c.api.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var services []domain.Service
	for _, cont := range containers {
		services = append(services, c.mapContainerToService(cont))
	}

	return services, nil
}

// mapContainerToService converts the raw container type to our domain model.
func (c *Client) mapContainerToService(cont types.Container) domain.Service {
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
	case "restarting", "paused", "removing", "dead":
		status = domain.StatusDegraded
	default:
		status = domain.StatusUnknown
	}

	uptime := time.Duration(0)
	if cont.State == "running" {
		uptime = time.Since(time.Unix(cont.Created, 0))
	}

	return domain.Service{
		ID:   cont.ID[:12],
		Name: name,
		Type: domain.ServiceTypeDocker,
		State: domain.ServiceState{
			Status: status,
			Uptime: uptime,
		},
		Network: domain.ServiceNetwork{
			Labels: cont.Labels,
		},
	}
}
