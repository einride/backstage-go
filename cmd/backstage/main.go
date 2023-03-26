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
	cmd := &cobra.Command{
		Use:   "backstage",
		Short: "Backstage CLI",
	}
	cmd.AddCommand(newAuthCommand())
	cmd.AddCommand(newCatalogCommand())
	return cmd
}

func newAuthCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate with a Backstage instance.",
	}
	cmd.AddCommand(newLoginCommand())
	return cmd
}

type authFile struct {
	BaseURL string `json:"baseUrl"`
	Token   string `json:"token"`
}

func newLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to a Backstage instance",
	}
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
	cmd := &cobra.Command{
		Use:   "catalog",
		Short: "Work with the Backstage catalog",
	}
	cmd.AddCommand(newListEntitiesCommand())
	return cmd
}

func newListEntitiesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-entities",
		Short: "List entities in the catalog",
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := newCatalogClient()
		if err != nil {
			return err
		}
		response, err := client.ListEntities(cmd.Context(), &catalog.ListEntitiesRequest{})
		if err != nil {
			return err
		}
		for _, entity := range response.Entities {
			printRawJSON(cmd, entity.Raw)
		}
		return nil
	}
	return cmd
}

func printRawJSON(cmd *cobra.Command, raw json.RawMessage) {
	var indented bytes.Buffer
	indented.Grow(len(raw) * 2)
	_ = json.Indent(&indented, raw, "", " ")
	cmd.Println(indented.String())
}
