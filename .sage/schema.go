package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"go.einride.tech/sage/sg"
)

func Schema(ctx context.Context) error {
	sg.Deps(
		ctx,
		sg.Fn(downloadSchema, "Entity.schema.json"),
		sg.Fn(downloadSchema, "EntityEnvelope.schema.json"),
		sg.Fn(downloadSchema, "EntityMeta.schema.json"),
		sg.Fn(downloadSchema, "shared/common.schema.json"),
		sg.Fn(downloadSchema, "kinds/API.v1alpha1.schema.json"),
		sg.Fn(downloadSchema, "kinds/Component.v1alpha1.schema.json"),
		sg.Fn(downloadSchema, "kinds/Domain.v1alpha1.schema.json"),
		sg.Fn(downloadSchema, "kinds/Group.v1alpha1.schema.json"),
		sg.Fn(downloadSchema, "kinds/Location.v1alpha1.schema.json"),
		sg.Fn(downloadSchema, "kinds/Resource.v1alpha1.schema.json"),
		sg.Fn(downloadSchema, "kinds/System.v1alpha1.schema.json"),
		sg.Fn(downloadSchema, "kinds/User.v1alpha1.schema.json"),
	)
	return nil
}

func downloadSchema(ctx context.Context, path string) error {
	const version = "v1.12.1"
	url := fmt.Sprintf(
		"https://raw.githubusercontent.com/backstage/backstage/%s/packages/catalog-model/src/schema/%s",
		version,
		path,
	)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	out := sg.FromGitRoot("cmd", "backstage", "internal", "schema", filepath.Base(path))
	if err := os.WriteFile(out, data, 0o600); err != nil {
		return err
	}
	sg.Logger(ctx).Println("wrote schema", out)
	return nil
}
