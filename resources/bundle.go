package resources

import (
	d "github.com/monarko/fhirgo/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Bundle resource
type Bundle struct {
	Base
	Identifier *d.Identifier   `json:"identifier,omitempty"`
	Type       *d.Code         `json:"type,omitempty"`
	Timestamp  *d.Instant      `json:"timestamp,omitempty"`
	Total      *d.UnsignedInt  `json:"total,omitempty"`
	Link       []d.BundleLink  `json:"link,omitempty"`
	Entry      []d.BundleEntry `json:"entry,omitempty"`
	Signature  *d.Signature    `json:"signature,omitempty"`
}

// Validate returns a check against schema
func (b *Bundle) Validate() (bool, []error) {
	return schema.ValidateResource(*b, "4.0")
}
