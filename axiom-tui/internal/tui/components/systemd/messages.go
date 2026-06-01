package systemdcomponent

import "axiomtui/internal/domain"

// ServicesLoaded is emitted when systemd services have been fetched.
type ServicesLoaded struct {
	Services []domain.Service
}

// ServicesError is emitted when fetching systemd services fails.
type ServicesError struct {
	Err error
}
