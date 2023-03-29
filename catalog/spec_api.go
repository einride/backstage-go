package catalog

// APISpec contains the API standard spec fields.
//
// See: https://backstage.io/docs/features/software-catalog/descriptor-format/#kind-api
type APISpec struct {
	// The type of API. This field is required.
	Type string `json:"type"`
	// The lifecycle state of the API. This field is required.
	Lifecycle string `json:"lifecycle"`
	// An entity reference to the owner of the API. This field is required.
	Owner string `json:"owner"`
	// An entity reference to the system that the component belongs to.
	System string `json:"system,omitempty"`
	// The definition of the API, based on the format defined by [APISpec.Type]. This field is required.
	Definition string `json:"definition"`
}
