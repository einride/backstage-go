package catalog

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// ListEntitiesRequest is the request to the [Client.ListEntities] method.
type ListEntitiesRequest struct {
	// Filters for selecting only a subset of all entities.
	Filters []string
	// Fields for selecting only parts of the full data structure of each entity.
	Fields []string
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
	// NextPageToken contains the next page token.
	NextPageToken string
}

var linkURLRegexp = regexp.MustCompile(`<(.*)>; *rel="next"`)

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
	for _, filter := range request.Filters {
		query.Add("filter", filter)
	}
	if len(request.Fields) > 0 {
		query.Set("fields", strings.Join(request.Fields, ","))
	}
	if request.After != "" {
		query.Set("after", request.After)
	}
	var rawEntities []json.RawMessage
	var nextPageToken string
	if err := c.get(ctx, path, query, func(response *http.Response) error {
		for _, link := range response.Header.Values("link") {
			if matches := linkURLRegexp.FindStringSubmatch(link); len(matches) > 1 {
				linkURL, err := url.ParseRequestURI(matches[1])
				if err != nil {
					return err
				}
				nextPageToken = linkURL.Query().Get("after")
			}
		}
		return json.NewDecoder(response.Body).Decode(&rawEntities)
	}); err != nil {
		return nil, err
	}
	response := ListEntitiesResponse{
		Entities:      make([]*Entity, 0, len(rawEntities)),
		NextPageToken: nextPageToken,
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
