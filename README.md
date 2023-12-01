<img src="./docs/backstage-go.svg" align="left" width="140" />

# Backstage Go SDK

A Go SDK and CLI tool for working with [Backstage](https://backstage.io).

## CLI tool

The `backstage` CLI tool provides command-line access to Backstage APIs.

```bash
# Authenticate to your Backstage instance.
$ backstage auth login --base-url "https://your-backstage.com" --token "<TOKEN>"

# List component entities in the catalog.
$ backstage catalog entities list --filter "kind=Component"

# Get an entity in the catalog.
$ backstage catalog entities get-by-name --kind "User" --name "odsod"

# Validate catalog entities in the ".backstage" dir.
$ backstage catalog entities validate ".backstage"
```

The CLI tool can be downloaded from the
[Releases](https://github.com/einride/backstage-go/releases) page.

## Software Catalog API

The [`catalog`](https://pkg.go.dev/go.einride.tech/backstage/catalog) package
provides a Go client to the
[Software Catalog API](https://backstage.io/docs/features/software-catalog/software-catalog-api).

```go
package main

import (
	"context"
	"fmt"

	"go.einride.tech/backstage/catalog"
)

func main() {
	ctx := context.Background()
	// Create a Software Catalog API client.
	client := catalog.NewClient(
		catalog.WithBaseURL("https://your-backstage-instance.example.com"),
		catalog.WithToken("YOUR_API_AUTH_TOKEN"),
	)
	// List component entities.
	response, err := client.ListEntities(ctx, &catalog.ListEntitiesRequest{
		Filters: []string{"kind=Component"},
	})
	if err != nil {
		panic(err)
	}
	for _, entity := range response.Entities {
		// Standard fields are parsed into Go structs.
		fmt.Println(entity.Metadata.Name)
		// Standard fields on specs can be parsed into Go structs.
		spec, err := entity.ComponentSpec()
		if err != nil {
			panic(err)
		}
		fmt.Println(spec.Lifecycle)
		// Custom fields can be accessed via raw JSON.
		fmt.Println(string(entity.Raw))
	}
}
```
