package catalog

// ResourceSpec contains the Resource standard spec fields.
//
// See: https://backstage.io/docs/features/software-catalog/descriptor-format#kind-resource
type ResourceSpec struct {
	// The type of resource. This field is required.
	Type string `json:"type"`
	// An entity reference to the owner of the resource. This field is required.
	Owner string `json:"owner"`
	// An entity reference to the system that the resource belongs to.
	System string `json:"system,omitempty"`
	// An array of entity references to the resources and resources that the resource depends on.
	DependsOn []string `json:"dependsOn,omitempty"`
	// An array of entity references to the components and resources that the resource is a dependency of.
	DependencyOf []string `json:"dependencyOf,omitempty"`
}
