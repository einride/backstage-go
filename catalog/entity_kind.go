package catalog

// EntityKind represents a known entity kind.
type EntityKind string

// Known entity kinds.
const (
	// EntityKindComponent represents a Component entity kind.
	EntityKindComponent EntityKind = "Component"

	// EntityKindSystem represents a System entity kind.
	EntityKindSystem EntityKind = "System"

	// EntityKindDomain represents a Domain entity kind.
	EntityKindDomain EntityKind = "Domain"

	// EntityKindUser represents a User entity kind.
	EntityKindUser EntityKind = "User"

	// EntityKindAPI represents an API entity kind.
	EntityKindAPI EntityKind = "API"

	// EntityKindResource represents a Resource entity kind.
	EntityKindResource EntityKind = "Resource"

	// EntityKindLocation represents a Location entity kind.
	EntityKindLocation EntityKind = "Location"

	// EntityKindTemplate represents a Template entity kind.
	EntityKindTemplate EntityKind = "Template"

	// EntityKindGroup represents a Group entity kind.
	EntityKindGroup EntityKind = "Group"
)
