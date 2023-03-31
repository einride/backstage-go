package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestClient_ListEntities(t *testing.T) {
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
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/catalog/entities", r.URL.Path)
			assert.Equal(t, "kind=Component", r.URL.Query().Get("filter"))
			assert.Equal(t, "10", r.URL.Query().Get("offset"))
			assert.Equal(t, "100", r.URL.Query().Get("limit"))
			assert.Equal(t, "bar", r.URL.Query().Get("after"))
			w.Header().Set("link", `<https://example.com?after=foo>; rel="next"`)
			_, _ = w.Write([]byte(fmt.Sprintf(`[%s,%s]`, system1, system2)))
		})
		actual, err := client.ListEntities(ctx, &ListEntitiesRequest{
			Filters: []string{"kind=Component"},
			Fields:  []string{"baz"},
			Offset:  10,
			Limit:   100,
			After:   "bar",
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, expected, actual.Entities)
		assert.Equal(t, "foo", actual.NextPageToken)
	})

	t.Run("fail", func(t *testing.T) {
		const statusCode = http.StatusInternalServerError
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/catalog/entities", r.URL.Path)
			w.WriteHeader(statusCode)
		})
		response, err := client.ListEntities(ctx, &ListEntitiesRequest{})
		assert.Assert(t, response == nil)
		var errStatus *StatusError
		assert.Assert(t, errors.As(err, &errStatus))
		assert.Equal(t, statusCode, errStatus.StatusCode)
	})
}
