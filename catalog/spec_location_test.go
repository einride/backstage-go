package catalog

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEntity_LocationSpec(t *testing.T) {
	//nolint: lll
	const domain = `{"apiVersion":"backstage.io/v1alpha1","kind":"Location","metadata":{"name":"org-data"},"spec":{"type":"url","targets":["https://github.com/myorg/myproject/org-data-dump/catalog-info-staff.yaml","https://github.com/myorg/myproject/org-data-dump/catalog-info-consultants.yaml"]}}`
	var entity Entity
	assert.NilError(t, json.Unmarshal([]byte(domain), &entity))
	expected := &LocationSpec{
		Type: "url",
		Targets: []string{
			"https://github.com/myorg/myproject/org-data-dump/catalog-info-staff.yaml",
			"https://github.com/myorg/myproject/org-data-dump/catalog-info-consultants.yaml",
		},
	}
	actual, err := entity.LocationSpec()
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, actual)
}
