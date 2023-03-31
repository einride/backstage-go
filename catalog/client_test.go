package catalog

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/v3/assert"
)

const testToken = "HELLO_WORLD"

func TestNewClient(t *testing.T) {
	ctx := context.Background()

	t.Run("authorization", func(t *testing.T) {
		var authorization string
		client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			authorization = r.Header.Get("authorization")
			_, err := w.Write([]byte("{}"))
			assert.NilError(t, err)
		})
		_, err := client.GetEntityByUID(ctx, &GetEntityByUIDRequest{
			UID: "test",
		})
		assert.NilError(t, err)
		assert.Equal(t, "Bearer "+testToken, authorization)
	})
}

func newTestClient(t *testing.T, handler func(http.ResponseWriter, *http.Request)) *Client {
	server := httptest.NewServer(http.HandlerFunc(handler))
	t.Cleanup(server.Close)
	return NewClient(
		WithBaseURL(server.URL),
		WithToken(testToken),
	)
}
