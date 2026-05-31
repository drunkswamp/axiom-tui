package dockercomponent

import "axiomtui/internal/domain"

// ServicesLoaded is a message sent when the list of services has been loaded.
type ServicesLoaded struct {
	Services []domain.Service
}

// ServicesError is a message sent when an error occurs.
type ServicesError struct {
	Err error
}
