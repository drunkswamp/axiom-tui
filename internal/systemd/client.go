package systemd

import (
	"axiomtui/internal/domain"
	"time"
)

// Client является мок-клиентом для systemd на не-Linux системах.
type Client struct{}

// NewClient создает новый мок-клиент.
func NewClient() Client {
	return Client{}
}

// ListServices возвращает предопределенный список мок-сервисов.
func (c Client) ListServices() ([]domain.Service, error) {
	mockServices := []domain.Service{
		{
			ID:   "nginx.service",
			Name: "Nginx Web Server",
			Type: domain.ServiceTypeSystemd,
			State: domain.ServiceState{
				Status: domain.StatusRunning,
				Uptime: time.Hour*24*3 + time.Minute*15,
			},
			Resources: domain.ServiceResources{
				CPUUsage:    1.5,
				MemoryUsage: 256 * 1024 * 1024, // 256 MB
			},
		},
		{
			ID:   "postgresql.service",
			Name: "PostgreSQL Database Server",
			Type: domain.ServiceTypeSystemd,
			State: domain.ServiceState{
				Status: domain.StatusRunning,
				Uptime: time.Hour * 24 * 7,
			},
			Resources: domain.ServiceResources{
				CPUUsage:    5.2,
				MemoryUsage: 1024 * 1024 * 1024, // 1 GB
			},
		},
		{
			ID:   "minecraft.service",
			Name: "Minecraft Server",
			Type: domain.ServiceTypeSystemd,
			State: domain.ServiceState{
				Status: domain.StatusDegraded, // Используем Degraded для 'failed'
				Uptime: 0,
			},
			Resources: domain.ServiceResources{
				CPUUsage:    0,
				MemoryUsage: 0,
			},
		},
		{
			ID:   "backup-script.service",
			Name: "Daily Backup Script",
			Type: domain.ServiceTypeSystemd,
			State: domain.ServiceState{
				Status: domain.StatusStopped,
				Uptime: 0,
			},
		},
	}

	return mockServices, nil
}
