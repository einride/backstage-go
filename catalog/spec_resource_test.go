package catalog

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEntity_ResourceSpec(t *testing.T) {
	//nolint: lll
	const domain = `{"apiVersion":"backstage.io/v1alpha1","kind":"Resource","metadata":{"name":"artists-db","description":"Stores artist details"},"spec":{"type":"database","owner":"team-a","system":"artist-engagement-portal"}}`
	var entity Entity
	assert.NilError(t, json.Unmarshal([]byte(domain), &entity))
	expected := &ResourceSpec{
		Type:   "database",
		Owner:  "team-a",
		System: "artist-engagement-portal",
	}
	actual, err := entity.ResourceSpec()
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, actual)
}
