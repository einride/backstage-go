package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestClient_GetEntityByUID(t *testing.T) {
	ctx := context.Background()
	t.Run("success", func(t *testing.T) {
		const system = `{"apiVersion":"backstage.io/v1alpha1","kind":"System","metadata":{"name":"podcast"}}`
		expected := &Entity{
			APIVersion: "backstage.io/v1alpha1",
			Kind:       EntityKindSystem,
			Metadata: EntityMetadata{
				Name: "podcast",
			},
			Raw: json.RawMessage(system),
		}
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/catalog/entities/by-uid/foo", r.URL.Path)
			_, _ = w.Write([]byte(system))
		})
		actual, err := client.GetEntityByUID(ctx, &GetEntityByUIDRequest{
			UID: "foo",
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, expected, actual)
	})

	t.Run("fail", func(t *testing.T) {
		const statusCode = http.StatusNotFound
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/catalog/entities/by-uid/foo", r.URL.Path)
			w.WriteHeader(statusCode)
		})
		entity, err := client.GetEntityByUID(ctx, &GetEntityByUIDRequest{
			UID: "foo",
		})
		assert.Assert(t, entity == nil)
		var errStatus *StatusError
		assert.Assert(t, errors.As(err, &errStatus))
		assert.Equal(t, statusCode, errStatus.StatusCode)
	})
}
