package catalog

// LocationSpec contains the Location standard spec fields.
//
// See: https://backstage.io/docs/features/software-catalog/descriptor-format#kind-location
type LocationSpec struct {
	// Type is the single location type, that's common to the targets specified in the spec.
	//
	// If it is left out, it is inherited from the location type that originally read the
	// entity data.
	//
	// For example, if you have a URL type location, that when read results in a Location kind
	// entity with no spec.type, then the referenced targets in the entity will implicitly also
	// be of url type.
	//
	// This is useful because you can define a hierarchy of things in a directory structure
	// using relative target paths (see below), and it will work out no matter if it's consumed
	// locally on disk from a file location, or as uploaded on a VCS.
	Type string `json:"type,omitempty"`

	// Target as a string.
	//
	// Can be either an absolute path/URL (depending on the type), or a relative path
	// such as ./details/catalog-info.yaml which is resolved relative to the location of
	// the Location entity itself.
	Target string `json:"target,omitempty"`

	// Targets as strings.
	//
	// They can all be either absolute paths/URLs (depending on the type), or relative paths
	// such as ./details/catalog-info.yaml which are resolved relative to the location of
	// the Location entity itself.
	Targets []string `json:"targets,omitempty"`

	// Describes whether the target of a location is required to exist or not.
	// It defaults to 'required' if not specified, can also be 'optional'.
	Presence LocationPresence `json:"presence,omitempty"`
}

// LocationPresence represents a location presence.
type LocationPresence string

// Known LocationPresence values.
const (
	// LocationPresenceRequired describes a location that is required to exist.
	LocationPresenceRequired = "required"
	// LocationPresenceOptional describes a location that is not required to exist.
	LocationPresenceOptional = "optional"
)
