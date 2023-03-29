package catalog

// ComponentSpec contains the Component standard spec fields.
//
// See: https://backstage.io/docs/features/software-catalog/descriptor-format#kind-component
type ComponentSpec struct {
	// The type of component. This field is required.
	Type string `json:"type"`
	// The lifecycle state of the component. This field is required.
	Lifecycle string `json:"lifecycle"`
	// An entity reference to the owner of the component. This field is required.
	Owner string `json:"owner"`
	// An entity reference to the system that the component belongs to.
	System string `json:"system,omitempty"`
	// An entity reference to another component of which the component is a part.
	SubcomponentOf string `json:"subcomponentOf,omitempty"`
	// An array of entity references to the APIs that are provided by the component.
	ProvidesAPIs []string `json:"providesApis,omitempty"`
	// An array of entity references to the APIs that are consumed by the component.
	ConsumesAPIs []string `json:"consumesApis,omitempty"`
	// An array of entity references to the components and resources that the component depends on.
	DependsOn []string `json:"dependsOn,omitempty"`
}
