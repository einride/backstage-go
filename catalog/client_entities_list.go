package catalog

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// ListEntitiesRequest is the request to the [Client.ListEntities] method.
type ListEntitiesRequest struct {
	// Offset for pagination.
	Offset int64
	// Limit for pagination.
	Limit int64
	// After for returning the next page after the provided cursor.
	After string
}

// ListEntitiesResponse is the response from the [Client.ListEntities] method.
type ListEntitiesResponse struct {
	// Entities in the response.
	Entities []*Entity
}

// ListEntities lists entities in the catalog.
//
// See: https://backstage.io/docs/features/software-catalog/software-catalog-api/#get-entities
func (c *Client) ListEntities(ctx context.Context, _ *ListEntitiesRequest) (*ListEntitiesResponse, error) {
	const path = "/api/catalog/entities"
	// TODO: Set request query parameters.
	query := make(url.Values)
	var rawEntities []json.RawMessage
	if err := c.get(ctx, path, query, func(response *http.Response) error {
		return json.NewDecoder(response.Body).Decode(&rawEntities)
	}); err != nil {
		return nil, err
	}
	response := ListEntitiesResponse{
		Entities: make([]*Entity, 0, len(rawEntities)),
	}
	for _, rawEntity := range rawEntities {
		response.Entities = append(response.Entities, &Entity{
			Raw: rawEntity,
		})
	}
	return &response, nil
}
