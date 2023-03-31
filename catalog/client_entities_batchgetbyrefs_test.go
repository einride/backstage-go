package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestClient_BatchGetEntitiesByRefs(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		const (
			system1 = `{"apiVersion":"backstage.io/v1alpha1","kind":"System","metadata":{"name":"system1"}}`
			system2 = `{"apiVersion":"backstage.io/v1alpha1","kind":"System","metadata":{"name":"system2"}}`
		)
		expected := []*Entity{
			{
				APIVersion: "backstage.io/v1alpha1",
				Kind:       EntityKindSystem,
				Metadata: EntityMetadata{
					Name: "system1",
				},
				Raw: json.RawMessage(system1),
			},
			{
				APIVersion: "backstage.io/v1alpha1",
				Kind:       EntityKindSystem,
				Metadata: EntityMetadata{
					Name: "system2",
				},
				Raw: json.RawMessage(system2),
			},
		}
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "/api/catalog/entities/by-refs", r.URL.Path)
			assert.Equal(t, "application/json", r.Header.Get("content-type"))
			body, err := io.ReadAll(r.Body)
			assert.NilError(t, err)
			assert.NilError(t, r.Body.Close())
			assert.Equal(t, `{"entityRefs":["foo","bar"],"fields":["baz"]}`, string(body))
			_, _ = w.Write([]byte(fmt.Sprintf(`{"items":[%s,%s]}`, system1, system2)))
		})
		actual, err := client.BatchGetEntitiesByRefs(ctx, &BatchGetEntitiesByRefsRequest{
			EntityRefs: []string{"foo", "bar"},
			Fields:     []string{"baz"},
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, expected, actual.Entities)
	})

	t.Run("fail", func(t *testing.T) {
		const statusCode = http.StatusInternalServerError
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "/api/catalog/entities/by-refs", r.URL.Path)
			assert.Equal(t, "application/json", r.Header.Get("content-type"))
			body, err := io.ReadAll(r.Body)
			assert.NilError(t, err)
			assert.NilError(t, r.Body.Close())
			assert.Equal(t, `{"entityRefs":["foo","bar"],"fields":["baz"]}`, string(body))
			w.WriteHeader(statusCode)
		})
		response, err := client.BatchGetEntitiesByRefs(ctx, &BatchGetEntitiesByRefsRequest{
			EntityRefs: []string{"foo", "bar"},
			Fields:     []string{"baz"},
		})
		assert.Assert(t, response == nil)
		var errStatus *StatusError
		assert.Assert(t, errors.As(err, &errStatus))
		assert.Equal(t, statusCode, errStatus.StatusCode)
	})
}
