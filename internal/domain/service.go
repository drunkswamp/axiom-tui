package domain

import "time"

// AxiomStatus — это наш унифицированный enum статусов.
type AxiomStatus string

const (
	StatusRunning  AxiomStatus = "Running"
	StatusStopped  AxiomStatus = "Stopped"
	StatusDegraded AxiomStatus = "Degraded"
	StatusUnknown  AxiomStatus = "Unknown"
)

// ServiceType — это дискриминатор для источника сервиса.
type ServiceType string

const (
	ServiceTypeDocker  ServiceType = "docker"
	ServiceTypeSystemd ServiceType = "systemd"
)

// PortMapping представляет проброс портов.
type PortMapping struct {
	HostPort      int
	ContainerPort int
	Protocol      string
}

// ServiceState содержит нормализованную информацию о состоянии.
type ServiceState struct {
	Status AxiomStatus
	Uptime time.Duration
}

// ServiceResources содержит информацию об использовании ресурсов.
type ServiceResources struct {
	CPUUsage    float64 // в процентах
	MemoryUsage uint64  // в байтах
}

// ServiceNetwork содержит сетевую информацию.
type ServiceNetwork struct {
	PortMappings []PortMapping
	Labels       map[string]string
}

// Service — это унифицированная абстракция для управляемой сущности.
type Service struct {
	ID        string
	Name      string
	Type      ServiceType
	State     ServiceState
	Resources ServiceResources
	Network   ServiceNetwork
}
