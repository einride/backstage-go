package catalog

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEntity_UserSpec(t *testing.T) {
	//nolint: lll
	const component = `{"apiVersion":"backstage.io/v1alpha1","kind":"User","metadata":{"name":"jdoe"},"spec":{"profile":{"displayName":"Jenny Doe","email":"jenny-doe@example.com","picture":"https://example.com/staff/jenny-with-party-hat.jpeg"},"memberOf":["team-b","employees"]}}`
	var entity Entity
	assert.NilError(t, json.Unmarshal([]byte(component), &entity))
	expected := &UserSpec{
		Profile: Profile{
			DisplayName: "Jenny Doe",
			Email:       "jenny-doe@example.com",
			Picture:     "https://example.com/staff/jenny-with-party-hat.jpeg",
		},
		MemberOf: []string{"team-b", "employees"},
	}
	actual, err := entity.UserSpec()
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, actual)
}
