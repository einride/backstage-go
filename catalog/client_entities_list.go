package catalog

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ListEntitiesRequest is the request to the [Client.ListEntities] method.
type ListEntitiesRequest struct {
	// Filter for selecting only a subset of all entities.
	Filter string
	// Fields for selecting only parts of the full data structure of each entity.
	Fields string
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
func (c *Client) ListEntities(ctx context.Context, request *ListEntitiesRequest) (*ListEntitiesResponse, error) {
	const path = "/api/catalog/entities"
	query := make(url.Values)
	if request.Offset > 0 {
		query.Set("offset", strconv.FormatInt(request.Offset, 10))
	}
	if request.Limit > 0 {
		query.Set("limit", strconv.FormatInt(request.Limit, 10))
	}
	if request.Filter != "" {
		query.Set("filter", request.Filter)
	}
	if request.Fields != "" {
		query.Set("fields", request.Fields)
	}
	if request.After != "" {
		query.Set("after", request.After)
	}
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
		entity := &Entity{
			Raw: rawEntity,
		}
		if err := json.Unmarshal(rawEntity, &entity); err != nil {
			return nil, err
		}
		response.Entities = append(response.Entities, entity)
	}
	return &response, nil
}
