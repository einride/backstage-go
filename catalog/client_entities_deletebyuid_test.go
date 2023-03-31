package catalog

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestClient_DeleteEntityByUID(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Equal(t, "/api/catalog/entities/by-uid/test", r.URL.Path)
			w.WriteHeader(http.StatusNoContent)
		})
		assert.NilError(t, client.DeleteEntityByUID(ctx, &DeleteEntityByUIDRequest{UID: "test"}))
	})

	t.Run("fail", func(t *testing.T) {
		const statusCode = http.StatusOK
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/api/catalog/entities/by-uid/test", r.URL.Path)
			w.WriteHeader(statusCode)
		})
		err := client.DeleteEntityByUID(ctx, &DeleteEntityByUIDRequest{UID: "test"})
		var errStatus *StatusError
		assert.Assert(t, errors.As(err, &errStatus))
		assert.Equal(t, statusCode, errStatus.StatusCode)
	})
}
