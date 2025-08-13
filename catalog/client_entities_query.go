package catalog

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// QueryEntitiesRequest is the request to the [Client.QueryEntities] method.
type QueryEntitiesRequest struct {
	// Filters for selecting only a subset of all entities.
	// Multiple filter sets with AND/OR conditions are supported.
	Filters []string
	// Fields for selecting only parts of the full data structure of each entity.
	Fields []string
	// Limit for the number of returned entities (default 20).
	Limit int64
	// OrderField for customizing entity sorting.
	// Format: "field,direction" where direction is "asc" or "desc".
	OrderField string
	// Cursor for cursor-based pagination.
	// Mutually exclusive with other query parameters.
	Cursor string
}

// PageInfo contains pagination information for cursor-based pagination.
type PageInfo struct {
	// NextCursor contains the cursor for the next page.
	NextCursor string `json:"nextCursor,omitempty"`
	// PrevCursor contains the cursor for the previous page.
	PrevCursor string `json:"prevCursor,omitempty"`
}

// QueryEntitiesResponse is the response from the [Client.QueryEntities] method.
type QueryEntitiesResponse struct {
	// Items contains the entities in the response.
	Entities []*Entity `json:"items"`
	// PageInfo contains pagination information.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
}

// QueryEntities queries entities in the catalog using the more efficient by-query endpoint.
//
// This method supports cursor-based pagination, flexible querying options, and customizable sorting.
// It provides better performance compared to the traditional ListEntities method.
//
// See: https://backstage.io/docs/features/software-catalog/software-catalog-api/#get-entities
func (c *Client) QueryEntities(ctx context.Context, request *QueryEntitiesRequest) (*QueryEntitiesResponse, error) {
	const path = "/api/catalog/entities/by-query"
	query := make(url.Values)

	if request.Limit > 0 {
		query.Set("limit", strconv.FormatInt(request.Limit, 10))
	}
	for _, filter := range request.Filters {
		query.Add("filter", filter)
	}
	if len(request.Fields) > 0 {
		query.Set("fields", strings.Join(request.Fields, ","))
	}
	if request.OrderField != "" {
		query.Set("orderField", request.OrderField)
	}
	if request.Cursor != "" {
		query.Set("cursor", request.Cursor)
	}
	var response QueryEntitiesResponse
	if err := c.get(ctx, path, query, func(httpResponse *http.Response) error {
		return json.NewDecoder(httpResponse.Body).Decode(&response)
	}); err != nil {
		return nil, err
	}

	return &response, nil
}
