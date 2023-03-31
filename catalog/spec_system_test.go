package catalog

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEntity_SystemSpec(t *testing.T) {
	//nolint: lll
	const system = `{"apiVersion":"backstage.io/v1alpha1","kind":"System","metadata":{"name":"podcast","description":"Podcast playback"},"spec":{"owner":"team-b","domain":"playback"}}`
	var entity Entity
	assert.NilError(t, json.Unmarshal([]byte(system), &entity))
	expected := &SystemSpec{
		Owner:  "team-b",
		Domain: "playback",
	}
	actual, err := entity.SystemSpec()
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, actual)
}
