module go.einride.tech/backstage/cmd/backstage

go 1.22

toolchain go1.22.5

require (
	github.com/adrg/xdg v0.5.3
	github.com/santhosh-tekuri/jsonschema v1.2.4
	github.com/spf13/cobra v1.9.1
	go.einride.tech/backstage v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	golang.org/x/sys v0.26.0 // indirect
)

replace go.einride.tech/backstage => ../../
