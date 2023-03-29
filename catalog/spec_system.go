package catalog

// SystemSpec contains the System standard spec fields.
//
// See: https://backstage.io/docs/features/software-catalog/descriptor-format#kind-system
type SystemSpec struct {
	// An entity reference to the owner of the system. This field is required.
	Owner string `json:"owner"`
	// An entity reference to the domain that the system belongs to.
	Domain string `json:"system,omitempty"`
}
