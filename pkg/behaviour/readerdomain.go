package behaviour

import "context"

type ReaderDomain interface {
	// All returns all the entities based on paramters
	All(context.Context, map[string]any) (any, error)
	// ByID returns the entity with the given id
	ByID(context.Context, string) (any, error)
}
