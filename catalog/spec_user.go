package catalog

// UserSpec contains the User standard spec fields.
//
// See: https://backstage.io/docs/features/software-catalog/descriptor-format#kind-user
type UserSpec struct {
	// Profile information about the user, mainly for display purposes.
	Profile Profile `json:"profile,omitempty"`
	// MemberOf is the list of groups that the user is a direct member of.
	MemberOf []string `json:"memberOf"`
}
