package catalog

import "encoding/json"

// An Entity in the software catalog.
type Entity struct {
	// Raw entity JSON message.
	Raw json.RawMessage
}
