package catalog

// Profile represents a user or group profile.
type Profile struct {
	// DisplayName of the profile.
	DisplayName string `json:"displayName,omitempty"`
	// Email of the profile.
	Email string `json:"email,omitempty"`
	// Picture URL of the profile.
	Picture string `json:"picture,omitempty"`
}
