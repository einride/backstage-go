package catalog

import "encoding/json"

// TemplateSpec contains the Template standard spec fields.
//
// See: https://backstage.io/docs/features/software-catalog/descriptor-format#kind-template
type TemplateSpec struct {
	// The type of component created by the template, e.g. website.
	// This is used for filtering templates, and should ideally match the Component spec.type created by the template.
	Type string `json:"type"`

	// An entity reference to the owner of the template.
	Owner string `json:"owner,omitempty"`

	// RawParameters contains the parameter specs.
	RawParameters []json.RawMessage `json:"parameters"`

	// RawSteps contains the step specs.
	RawSteps []json.RawMessage `json:"steps"`
}
