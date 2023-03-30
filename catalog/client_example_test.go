package catalog_test

import (
	"context"
	"fmt"

	"go.einride.tech/backstage/catalog"
)

func ExampleClient() {
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
