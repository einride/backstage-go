package catalog

import (
	"encoding/json"
	"fmt"
)

// An Entity in the software catalog.
type Entity struct {
	// APIVersion is the version of specification format for this particular entity.
	APIVersion string

	// Kind is the high-level entity type.
	Kind EntityKind

	// Metadata related to the entity.
	Metadata EntityMetadata

	// Relations that this entity has with other entities.
	Relations []EntityRelation

	// Raw entity JSON message.
	Raw json.RawMessage
}

// UnmarshalJSON implements [json.Unmarshaler].
func (e *Entity) UnmarshalJSON(data []byte) error {
	var fields struct {
		APIVersion string           `json:"apiVersion"`
		Kind       EntityKind       `json:"kind"`
		Metadata   EntityMetadata   `json:"metadata"`
		Relations  []EntityRelation `json:"relations,omitempty"`
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	e.APIVersion = fields.APIVersion
	e.Kind = fields.Kind
	e.Metadata = fields.Metadata
	e.Relations = fields.Relations
	e.Raw = data
	return nil
}

// APISpec decodes the entity's spec as a [APISpec].
func (e *Entity) APISpec() (*APISpec, error) {
	return unmarshalSpec[APISpec](EntityKindAPI, e)
}

// ComponentSpec decodes the entity's spec as a [ComponentSpec].
func (e *Entity) ComponentSpec() (*ComponentSpec, error) {
	return unmarshalSpec[ComponentSpec](EntityKindComponent, e)
}

// DomainSpec decodes the entity's spec as a [DomainSpec].
func (e *Entity) DomainSpec() (*DomainSpec, error) {
	return unmarshalSpec[DomainSpec](EntityKindDomain, e)
}

// GroupSpec decodes the entity's spec as a [GroupSpec].
func (e *Entity) GroupSpec() (*GroupSpec, error) {
	return unmarshalSpec[GroupSpec](EntityKindGroup, e)
}

// LocationSpec decodes the entity's spec as a [LocationSpec].
func (e *Entity) LocationSpec() (*LocationSpec, error) {
	return unmarshalSpec[LocationSpec](EntityKindLocation, e)
}

// ResourceSpec decodes the entity's spec as a [ResourceSpec].
func (e *Entity) ResourceSpec() (*ResourceSpec, error) {
	return unmarshalSpec[ResourceSpec](EntityKindResource, e)
}

// SystemSpec decodes the entity's spec as a [SystemSpec].
func (e *Entity) SystemSpec() (*SystemSpec, error) {
	return unmarshalSpec[SystemSpec](EntityKindSystem, e)
}

// TemplateSpec decodes the entity's spec as a [TemplateSpec].
func (e *Entity) TemplateSpec() (*TemplateSpec, error) {
	return unmarshalSpec[TemplateSpec](EntityKindTemplate, e)
}

// UserSpec decodes the entity's spec as a [UserSpec].
func (e *Entity) UserSpec() (*UserSpec, error) {
	return unmarshalSpec[UserSpec](EntityKindUser, e)
}

func unmarshalSpec[T any](kind EntityKind, e *Entity) (*T, error) {
	if e.Kind != kind {
		return nil, fmt.Errorf("expected kind %s but was %s", kind, e.Kind)
	}
	var result struct {
		Spec *T `json:"spec"`
	}
	if err := json.Unmarshal(e.Raw, &result); err != nil {
		return nil, err
	}
	return result.Spec, nil
}
