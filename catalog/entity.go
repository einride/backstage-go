package catalog

import (
	"encoding/json"
	"fmt"
)

// An Entity in the software catalog.
type Entity struct {
	// APIVersion is the version of specification format for this particular entity.
	APIVersion string `json:"apiVersion"`

	// Kind is the high-level entity type.
	Kind EntityKind `json:"kind"`

	// Metadata related to the entity.
	Metadata EntityMetadata `json:"metadata"`

	// Relations that this entity has with other entities.
	Relations []EntityRelation `json:"relations,omitempty"`

	// Raw entity JSON message.
	Raw json.RawMessage `json:"-"`
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
