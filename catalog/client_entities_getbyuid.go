package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetEntityByUIDRequest is the request to the [Client.GetEntityByUID] method.
type GetEntityByUIDRequest struct {
	// UID of the entity to get.
	UID string
}

// GetEntityByUID gets an entity by its kind, namespace and name.
//
// See: https://backstage.io/docs/features/software-catalog/software-catalog-api/#get-entitiesby-uiduid
func (c *Client) GetEntityByUID(ctx context.Context, request *GetEntityByUIDRequest) (*Entity, error) {
	const pathTemplate = "/api/catalog/entities/by-uid/%s"
	path := fmt.Sprintf(
		pathTemplate,
		url.PathEscape(request.UID),
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
	var entity Entity
	if err := json.Unmarshal(rawEntity, &entity); err != nil {
		return nil, err
	}
	return &entity, nil
}
