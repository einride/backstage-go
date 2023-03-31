package catalog

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEntity_ComponentSpec(t *testing.T) {
	//nolint: lll
	const component = "{\"apiVersion\":\"backstage.io/v1alpha1\",\"kind\":\"Component\",\"metadata\":{\"name\":\"petstore\",\"description\":\"[The Petstore](http://petstore.example.com) is an example API used to show features of the OpenAPI spec.\\n- First item\\n- Second item\\n\",\"links\":[{\"url\":\"https://github.com/swagger-api/swagger-petstore\",\"title\":\"GitHub Repo\",\"icon\":\"github\"}]},\"spec\":{\"type\":\"service\",\"lifecycle\":\"experimental\",\"owner\":\"team-c\",\"providesApis\":[\"petstore\",\"streetlights\",\"hello-world\"]}}"
	var entity Entity
	assert.NilError(t, json.Unmarshal([]byte(component), &entity))
	expected := &ComponentSpec{
		Type:         "service",
		Lifecycle:    "experimental",
		Owner:        "team-c",
		ProvidesAPIs: []string{"petstore", "streetlights", "hello-world"},
	}
	actual, err := entity.ComponentSpec()
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, actual)
}
