package catalog

import (
	"context"
	"encoding/json"
	"net/http"
)

// BatchGetEntitiesByRefsRequest is the request to the [Client.BatchGetEntitiesByRefs] method.
type BatchGetEntitiesByRefsRequest struct {
	// EntityRefs to fetch.
	// See: https://backstage.io/docs/features/software-catalog/references
	EntityRefs []string `json:"entityRefs"`

	// Fields to fetch.
	Fields []string `json:"fields,omitempty"`
}

// BatchGetEntitiesByRefsResponse is the response from the [Client.BatchGetEntitiesByRefs] method.
type BatchGetEntitiesByRefsResponse struct {
	// Entities returned.
	// Has the same length and the same order as the input entityRefs array.
	Entities []*Entity
}

// BatchGetEntitiesByRefs gets an entity by its kind, namespace and name.
// See: https://backstage.io/docs/features/software-catalog/software-catalog-api#post-entitiesby-refs
func (c *Client) BatchGetEntitiesByRefs(
	ctx context.Context,
	request *BatchGetEntitiesByRefsRequest,
) (*BatchGetEntitiesByRefsResponse, error) {
	const path = "/api/catalog/entities/by-refs"
	var responseBody struct {
		Items []*Entity
	}
	if err := c.post(ctx, path, request, func(r *http.Response) error {
		return json.NewDecoder(r.Body).Decode(&responseBody)
	}); err != nil {
		return nil, err
	}
	return &BatchGetEntitiesByRefsResponse{Entities: responseBody.Items}, nil
}
