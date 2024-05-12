package c3dt

// A 3D Tiles tileset.
type Tileset struct {
	// Metadata about the entire tileset.
	Asset Asset `json:"asset"`

	// The error, in meters, introduced if this tileset is not rendered. At
	// runtime, the geometric error is used to compute screen space error
	// (SSE), i.e., the error measured in pixels.
	GeometricError float64 `json:"geometricError"`

	// The root tile.
	Root Tile `json:"root"`

	// Names of 3D Tiles extensions used somewhere in this tileset.
	ExtensionsUsed []string `json:"extensionsUsed,omitempty"`

	// Names of 3D Tiles extensions required to properly load this tileset.
	// Each element of this array shall also be contained in `extensionsUsed`
	ExtensionsRequired []string `json:"extensionsRequired,omitempty"`

	// A metadata entity that is associated with this tileset.
	Metadata *MetadataEntity `json:"metadata,omitempty"`

	// An array of groups that tile content may belong to. Each element of this
	// array is a metadata entity that describes the group. The tile content
	// `group` property is an index into this array.
	Groups []Group `json:"groups,omitempty"`

	// An object containing statistics about metadata entities.
	Statistics *Statistics `json:"statistics,omitempty"`

	// The URI (or IRI) of the external schema file. When this is defined, then
	// `schema` shall be undefined.
	SchemaURI string `json:"schemaUri,omitempty"`

	// An object defining the structure of metadata classes and enums. When
	// this is defined, then `schemaUri` shall be undefined.
	Schema *Schema `json:"schema,omitempty"`
}
