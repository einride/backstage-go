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

func TestClient_QueryEntities(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		const (
			system1 = `{"apiVersion":"backstage.io/v1alpha1","kind":"System","metadata":{"name":"system1"}}`
			system2 = `{"apiVersion":"backstage.io/v1alpha1","kind":"System","metadata":{"name":"system2"}}`
		)
		expectedEntities := []*Entity{
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
			assert.Equal(t, "/api/catalog/entities/by-query", r.URL.Path)
			assert.Equal(t, "kind=Component", r.URL.Query().Get("filter"))
			assert.Equal(t, "metadata.name,asc", r.URL.Query().Get("orderField"))
			assert.Equal(t, "100", r.URL.Query().Get("limit"))
			assert.Equal(t, "", r.URL.Query().Get("cursor"))
			assert.Equal(t, "metadata.name,spec.type", r.URL.Query().Get("fields"))

			response := fmt.Sprintf(`{
				"items": [%s,%s],
				"totalItems": 2,
				"pageInfo": {
					"nextCursor": "nextCursor123",
					"prevCursor": "prevCursor456"
				}
			}`, system1, system2)
			_, _ = w.Write([]byte(response))
		})
		actual, err := client.QueryEntities(ctx, &QueryEntitiesRequest{
			Filters:    []string{"kind=Component"},
			Fields:     []string{"metadata.name", "spec.type"},
			Limit:      100,
			OrderField: "metadata.name,asc",
			Cursor:     "",
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, expectedEntities, actual.Entities)
		assert.Equal(t, int64(2), int64(len(actual.Entities)))
		assert.Equal(t, "nextCursor123", actual.PageInfo.NextCursor)
		assert.Equal(t, "prevCursor456", actual.PageInfo.PrevCursor)
	})

	t.Run("success with multiple filters", func(t *testing.T) {
		const system1 = `{"apiVersion":"backstage.io/v1alpha1","kind":"System","metadata":{"name":"system1"}}`
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/catalog/entities/by-query", r.URL.Path)
			filters := r.URL.Query()["filter"]
			assert.Equal(t, 2, len(filters))
			assert.Assert(t, contains(filters, "kind=Component"))
			assert.Assert(t, contains(filters, "spec.type=service"))

			response := fmt.Sprintf(`{
				"items": [%s],
				"totalItems": 1,
				"pageInfo": {
					"nextCursor": "next"
				}
			}`, system1)
			_, _ = w.Write([]byte(response))
		})
		actual, err := client.QueryEntities(ctx, &QueryEntitiesRequest{
			Filters: []string{"kind=Component", "spec.type=service"},
		})
		assert.NilError(t, err)
		assert.Equal(t, 1, len(actual.Entities))
		assert.Equal(t, int64(1), int64(len(actual.Entities)))
		assert.Equal(t, "next", actual.PageInfo.NextCursor)
	})

	t.Run("success with minimal request", func(t *testing.T) {
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/catalog/entities/by-query", r.URL.Path)
			assert.Equal(t, "", r.URL.Query().Get("limit"))
			assert.Equal(t, "", r.URL.Query().Get("cursor"))
			assert.Equal(t, "", r.URL.Query().Get("orderField"))
			assert.Equal(t, "", r.URL.Query().Get("fields"))
			assert.Equal(t, 0, len(r.URL.Query()["filter"]))

			_, _ = w.Write([]byte(`{
				"items": [],
				"totalItems": 0
			}`))
		})
		actual, err := client.QueryEntities(ctx, &QueryEntitiesRequest{})
		assert.NilError(t, err)
		assert.Equal(t, 0, len(actual.Entities))
		assert.Equal(t, int64(0), int64(len(actual.Entities)))
		assert.Assert(t, actual.PageInfo == nil)
	})

	t.Run("success without pageInfo", func(t *testing.T) {
		const system1 = `{"apiVersion":"backstage.io/v1alpha1","kind":"System","metadata":{"name":"system1"}}`
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/catalog/entities/by-query", r.URL.Path)

			response := fmt.Sprintf(`{
				"items": [%s],
				"totalItems": 1
			}`, system1)
			_, _ = w.Write([]byte(response))
		})
		actual, err := client.QueryEntities(ctx, &QueryEntitiesRequest{})
		assert.NilError(t, err)
		assert.Equal(t, 1, len(actual.Entities))
		assert.Equal(t, int64(1), int64(len(actual.Entities)))
		assert.Assert(t, actual.PageInfo == nil)
	})

	t.Run("fail", func(t *testing.T) {
		const statusCode = http.StatusInternalServerError
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/catalog/entities/by-query", r.URL.Path)
			w.WriteHeader(statusCode)
		})
		response, err := client.QueryEntities(ctx, &QueryEntitiesRequest{})
		assert.Assert(t, response == nil)
		var errStatus *StatusError
		assert.Assert(t, errors.As(err, &errStatus))
		assert.Equal(t, statusCode, errStatus.StatusCode)
	})

	t.Run("invalid json response", func(t *testing.T) {
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/catalog/entities/by-query", r.URL.Path)
			_, _ = w.Write([]byte(`invalid json`))
		})
		response, err := client.QueryEntities(ctx, &QueryEntitiesRequest{})
		assert.Assert(t, response == nil)
		assert.Assert(t, err != nil)
	})
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
