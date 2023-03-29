package catalog

// GroupSpec contains the Group standard spec fields.
//
// See: https://backstage.io/docs/features/software-catalog/descriptor-format#kind-group
type GroupSpec struct {
	// The type of group. This field is required.
	//
	// There is currently no enforced set of values for this field,
	// so it is left up to the adopting organization to choose a nomenclature
	// that matches their org hierarchy.
	Type string `json:"type"`
	// Profile information about the group, mainly for display purposes.
	Profile Profile `json:"profile,omitempty"`
	// Parent group in the hierarchy, if any.
	Parent string `json:"parent,omitempty"`
	// Child groups of this group in the hierarchy (whose parent field points to this group).
	// The items are not guaranteed to be ordered in any particular way.
	Children []string `json:"children"`
	// Members are the users that are direct members of this group.
	// The items are not guaranteed to be ordered in any particular way.
	Members []string `json:"members,omitempty"`
}
