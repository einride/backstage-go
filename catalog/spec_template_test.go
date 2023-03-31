package catalog

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEntity_TemplateSpec(t *testing.T) {
	//nolint: lll
	const domain = `{"apiVersion":"backstage.io/v1beta2","kind":"Template","metadata":{"name":"v1beta2-demo","title":"Test Action template","description":"scaffolder v1beta2 template demo"},"spec":{"owner":"backstage/techdocs-core","type":"service","parameters":[{"title":"Fill in some steps","required":["name"],"properties":{"name":{"title":"Name","type":"string","description":"Unique name of the component","ui:autofocus":true,"ui:options":{"rows":5}}}},{"title":"Choose a location","required":["repoUrl"],"properties":{"repoUrl":{"title":"Repository Location","type":"string","ui:field":"RepoUrlPicker","ui:options":{"allowedHosts":["github.com"]}}}}],"steps":[{"id":"fetch-base","name":"Fetch Base","action":"fetch:template","input":{"url":"./template","values":{"name":"{{ parameters.name }}"}}},{"id":"fetch-docs","name":"Fetch Docs","action":"fetch:plain","input":{"targetPath":"./community","url":"https://github.com/backstage/community/tree/main/backstage-community-sessions"}}]}}`
	var entity Entity
	assert.NilError(t, json.Unmarshal([]byte(domain), &entity))
	//nolint: lll
	expected := &TemplateSpec{
		Type:  "service",
		Owner: "backstage/techdocs-core",
		RawParameters: []json.RawMessage{
			json.RawMessage(`{"title":"Fill in some steps","required":["name"],"properties":{"name":{"title":"Name","type":"string","description":"Unique name of the component","ui:autofocus":true,"ui:options":{"rows":5}}}}`),
			json.RawMessage(`{"title":"Choose a location","required":["repoUrl"],"properties":{"repoUrl":{"title":"Repository Location","type":"string","ui:field":"RepoUrlPicker","ui:options":{"allowedHosts":["github.com"]}}}}`),
		},
		RawSteps: []json.RawMessage{
			json.RawMessage(`{"id":"fetch-base","name":"Fetch Base","action":"fetch:template","input":{"url":"./template","values":{"name":"{{ parameters.name }}"}}}`),
			json.RawMessage(`{"id":"fetch-docs","name":"Fetch Docs","action":"fetch:plain","input":{"targetPath":"./community","url":"https://github.com/backstage/community/tree/main/backstage-community-sessions"}}`),
		},
	}
	actual, err := entity.TemplateSpec()
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, actual)
}
