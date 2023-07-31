package behaviour

import "context"

// WriterDomain is the interface that wraps the basic CRUD methods.
type WriterDomain interface {
	// Create creates a new entity
	Create(ctx context.Context, entity any) (any, error)
	// Update updates the entity with the given id
	Update(ctx context.Context, id string, entity any) (any, error)
	// Delete deletes the entity with the given id
	Delete(ctx context.Context, id string) error
}
