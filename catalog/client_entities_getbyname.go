package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetEntityByNameRequest is the request to the [Client.GetEntityByName] method.
type GetEntityByNameRequest struct {
	// Kind of the entity to get.
	Kind string
	// Namespace of the entity to get.
	Namespace string
	// Name of the entity to get.
	Name string
}

// GetEntityByName gets an entity by its kind, namespace and name.
//
// See: https://backstage.io/docs/features/software-catalog/software-catalog-api/#get-entitiesby-namekindnamespacename
func (c *Client) GetEntityByName(ctx context.Context, request *GetEntityByNameRequest) (*Entity, error) {
	const pathTemplate = "/api/catalog/entities/by-name/%s/%s/%s"
	path := fmt.Sprintf(
		pathTemplate,
		url.PathEscape(request.Kind),
		url.PathEscape(request.Namespace),
		url.PathEscape(request.Name),
	)
	var rawEntity json.RawMessage
	if err := c.get(ctx, path, nil, func(response *http.Response) error {
		data, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		rawEntity = data
		return nil
	}); err != nil {
		return nil, err
	}
	entity := Entity{
		Raw: rawEntity,
	}
	if err := json.Unmarshal(rawEntity, &entity); err != nil {
		return nil, err
	}
	return &entity, nil
}
