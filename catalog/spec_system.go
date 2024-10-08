package catalog

// SystemSpec contains the System standard spec fields.
//
// See: https://backstage.io/docs/features/software-catalog/descriptor-format#kind-system
type SystemSpec struct {
	// An entity reference to the owner of the system. This field is required.
	Owner string `json:"owner"`
	// An entity reference to the domain that the system belongs to.
	Domain string `json:"domain,omitempty"`
	// The type of system. There is currently no enforced set of values for this field, so it is left up to the adopting organization to choose a nomenclature that matches their catalog hierarchy. This field is optional.
	Type string `json:"type,omitempty"`
}
