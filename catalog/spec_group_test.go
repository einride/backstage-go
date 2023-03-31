package catalog

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEntity_GroupSpec(t *testing.T) {
	//nolint: lll
	const domain = `{"apiVersion":"backstage.io/v1alpha1","kind":"Group","metadata":{"name":"infrastructure","description":"The infra business unit"},"spec":{"type":"business-unit","profile":{"displayName":"Infrastructure","email":"infrastructure@example.com","picture":"https://example.com/groups/bu-infrastructure.jpeg"},"parent":"ops","children":["backstage","other"],"members":["jdoe"]}}`
	var entity Entity
	assert.NilError(t, json.Unmarshal([]byte(domain), &entity))
	expected := &GroupSpec{
		Type: "business-unit",
		Profile: Profile{
			DisplayName: "Infrastructure",
			Email:       "infrastructure@example.com",
			Picture:     "https://example.com/groups/bu-infrastructure.jpeg",
		},
		Parent:   "ops",
		Children: []string{"backstage", "other"},
		Members:  []string{"jdoe"},
	}
	actual, err := entity.GroupSpec()
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, actual)
}
