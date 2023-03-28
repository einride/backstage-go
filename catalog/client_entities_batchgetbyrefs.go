package catalog

import (
	"bytes"
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
	Entities []*Entity `json:"items"`
}

// BatchGetEntitiesByRefs gets an entity by its kind, namespace and name.
// See: https://backstage.io/docs/features/software-catalog/software-catalog-api#post-entitiesby-refs
func (c *Client) BatchGetEntitiesByRefs(
	ctx context.Context,
	request *BatchGetEntitiesByRefsRequest,
) (*BatchGetEntitiesByRefsResponse, error) {
	const path = "/api/catalog/entities/by-refs"
	var responseBody struct {
		RawEntities []json.RawMessage `json:"items"`
	}
	if err := c.post(ctx, path, request, func(response *http.Response) error {
		return json.NewDecoder(response.Body).Decode(&responseBody)
	}); err != nil {
		return nil, err
	}
	response := BatchGetEntitiesByRefsResponse{
		Entities: make([]*Entity, 0, len(responseBody.RawEntities)),
	}
	for _, rawEntity := range responseBody.RawEntities {
		if bytes.Equal(rawEntity, []byte("null")) {
			response.Entities = append(response.Entities, nil)
		} else {
			entity := Entity{
				Raw: rawEntity,
			}
			if err := json.Unmarshal(rawEntity, &entity); err != nil {
				return nil, err
			}
			response.Entities = append(response.Entities, &entity)
		}
	}
	return &response, nil
}
