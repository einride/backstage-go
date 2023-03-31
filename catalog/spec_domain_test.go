package catalog

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEntity_DomainSpec(t *testing.T) {
	//nolint: lll
	const domain = `{"apiVersion":"backstage.io/v1alpha1","kind":"Domain","metadata":{"name":"playback","description":"Everything related to audio playback"},"spec":{"owner":"user:frank.tiernan"}}`
	var entity Entity
	assert.NilError(t, json.Unmarshal([]byte(domain), &entity))
	expected := &DomainSpec{
		Owner: "user:frank.tiernan",
	}
	actual, err := entity.DomainSpec()
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, actual)
}
