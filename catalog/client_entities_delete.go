package catalog

import (
	"context"
	"fmt"
	"net/url"
)

// DeleteEntityByUIDRequest is the request to the [Client.DeleteEntityByUID] method.
type DeleteEntityByUIDRequest struct {
	// UID of the entity to get.
	UID string
}

// DeleteEntityByUID gets an entity by its kind, namespace and name.
//
// See: https://backstage.io/docs/features/software-catalog/software-catalog-api/#get-entitiesby-uiduid
func (c *Client) DeleteEntityByUID(ctx context.Context, request *DeleteEntityByUIDRequest) error {
	const pathTemplate = "/api/catalog/entities/by-uid/%s"
	return c.delete(ctx, fmt.Sprintf(pathTemplate, url.PathEscape(request.UID)))
}
