package catalog

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEntity_APISpec(t *testing.T) {
	//nolint: lll
	const component = `{"apiVersion":"backstage.io/v1alpha1","kind":"API","metadata":{"name":"artist-api","description":"Retrieve artist details"},"spec":{"type":"openapi","lifecycle":"production","owner":"artist-relations-team","system":"artist-engagement-portal","definition":"openapi: \"3.0.0\"\ninfo:\n  version: 1.0.0\n  title: Artist API\n  license:\n    name: MIT\nservers:\n  - url: http://artist.spotify.net/v1\npaths:\n  /artists:\n    get:\n      summary: List all artists\n...\n"}}`
	var entity Entity
	assert.NilError(t, json.Unmarshal([]byte(component), &entity))
	expected := &APISpec{
		Type:      "openapi",
		Lifecycle: "production",
		Owner:     "artist-relations-team",
		System:    "artist-engagement-portal",
		//nolint: lll
		Definition: "openapi: \"3.0.0\"\ninfo:\n  version: 1.0.0\n  title: Artist API\n  license:\n    name: MIT\nservers:\n  - url: http://artist.spotify.net/v1\npaths:\n  /artists:\n    get:\n      summary: List all artists\n...\n",
	}
	actual, err := entity.APISpec()
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, actual)
}
