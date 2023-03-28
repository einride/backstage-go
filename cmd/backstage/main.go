package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"go.einride.tech/backstage/catalog"
)

func main() {
	if err := newBackstageCommand().Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const authConfigFile = "backstage/auth.json"

func newCatalogClient() (*catalog.Client, error) {
	authFilepath, err := xdg.ConfigFile(authConfigFile)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(authFilepath); err != nil {
		return nil, err
	}
	data, err := os.ReadFile(authFilepath)
	if err != nil {
		return nil, err
	}
	var authFileContent authFile
	if err := json.Unmarshal(data, &authFileContent); err != nil {
		return nil, err
	}
	return catalog.NewClient(
		catalog.WithBaseURL(authFileContent.BaseURL),
		catalog.WithToken(authFileContent.Token),
	), nil
}

func newBackstageCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "backstage"
	cmd.Short = "Backstage CLI"
	cmd.AddCommand(newAuthCommand())
	cmd.AddCommand(newCatalogCommand())
	return cmd
}

func newAuthCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "auth"
	cmd.Short = "Authenticate with a Backstage instance."
	cmd.AddCommand(newLoginCommand())
	return cmd
}

type authFile struct {
	BaseURL string `json:"baseUrl"`
	Token   string `json:"token"`
}

func newLoginCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "login"
	cmd.Short = "Login to a Backstage instance"
	baseURL := cmd.Flags().String("base-url", "", "backend base URL to login with")
	_ = cmd.MarkFlagRequired("base-url")
	token := cmd.PersistentFlags().String("token", "", "bearer token to use for authentication")
	_ = cmd.MarkFlagRequired("token")
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		authFilepath, err := xdg.ConfigFile(authConfigFile)
		if err != nil {
			return err
		}
		tokenData, err := json.MarshalIndent(&authFile{
			BaseURL: *baseURL,
			Token:   *token,
		}, "", "  ")
		if err != nil {
			return err
		}
		if err := os.WriteFile(authFilepath, tokenData, 0o600); err != nil {
			return err
		}
		cmd.Println()
		cmd.Println("Logged in.")
		return nil
	}
	return cmd
}

func newCatalogCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "catalog"
	cmd.Short = "Work with the Backstage catalog"
	cmd.AddCommand(newEntitiesCommand())
	return cmd
}

func newEntitiesCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "entities"
	cmd.Short = "Read entities from the Backstage catalog"
	cmd.AddCommand(newEntitiesListCommand())
	cmd.AddCommand(newEntitiesGetByUIDCommand())
	cmd.AddCommand(newEntitiesGetByNameCommand())
	cmd.AddCommand(newEntitiesDeleteByUIDCommand())
	return cmd
}

func newEntitiesListCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "list"
	cmd.Short = "List entities in the catalog"
	filters := cmd.Flags().StringArray("filter", nil, "select only a subset of all entities")
	fields := cmd.Flags().StringSlice("fields", nil, "select only parts of each entity")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := newCatalogClient()
		if err != nil {
			return err
		}
		var nextPageToken string
		for {
			response, err := client.ListEntities(cmd.Context(), &catalog.ListEntitiesRequest{
				Filters: *filters,
				Fields:  *fields,
				Limit:   100,
				After:   nextPageToken,
			})
			if err != nil {
				return err
			}
			for _, entity := range response.Entities {
				printRawJSON(cmd, entity.Raw)
			}
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		return nil
	}
	return cmd
}

func newEntitiesGetByNameCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "get-by-name"
	cmd.Short = "Get an entity by its kind, namespace and name"
	kind := cmd.Flags().String("kind", "", "kind of the entity to get")
	_ = cmd.MarkFlagRequired("kind")
	namespace := cmd.Flags().String("namespace", "default", "namespace of the entity to get")
	name := cmd.Flags().String("name", "", "name of the entity to get")
	_ = cmd.MarkFlagRequired("name")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := newCatalogClient()
		if err != nil {
			return err
		}
		entity, err := client.GetEntityByName(cmd.Context(), &catalog.GetEntityByNameRequest{
			Kind:      *kind,
			Namespace: *namespace,
			Name:      *name,
		})
		if err != nil {
			return err
		}
		printRawJSON(cmd, entity.Raw)
		return nil
	}
	return cmd
}

func newEntitiesGetByUIDCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "get-by-uid"
	cmd.Short = "Get an entity by its UID"
	uid := cmd.Flags().String("uid", "", "UID of the entity to get")
	_ = cmd.MarkFlagRequired("uid")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := newCatalogClient()
		if err != nil {
			return err
		}
		entity, err := client.GetEntityByUID(cmd.Context(), &catalog.GetEntityByUIDRequest{
			UID: *uid,
		})
		if err != nil {
			return err
		}
		printRawJSON(cmd, entity.Raw)
		return nil
	}
	return cmd
}

func newEntitiesDeleteByUIDCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "delete"
	cmd.Short = "Delete an entity by its UID"
	uid := cmd.Flags().String("uid", "", "UID of the entity to get")
	_ = cmd.MarkFlagRequired("uid")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := newCatalogClient()
		if err != nil {
			return err
		}
		return client.DeleteEntityByUID(cmd.Context(), &catalog.DeleteEntityByUIDRequest{
			UID: *uid,
		})
	}
	return cmd
}

func printRawJSON(cmd *cobra.Command, raw json.RawMessage) {
	var indented bytes.Buffer
	indented.Grow(len(raw) * 2)
	_ = json.Indent(&indented, raw, "", " ")
	cmd.Println(indented.String())
}

func newCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	return cmd
}
