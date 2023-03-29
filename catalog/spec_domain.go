package catalog

// DomainSpec contains the Domain standard spec fields.
//
// See: https://backstage.io/docs/features/software-catalog/descriptor-format#kind-domain
type DomainSpec struct {
	// An entity reference to the owner of the domain. This field is required.
	Owner string `json:"owner"`
}
