package c3dt

// Metadata about the entire tileset
type Asset struct {
	// The 3D Tiles version. The version defines the JSON schema for the
	// tileset JSON and the base set of tile formats.
	Version string `json:"version"`

	// Application-specific version of this tileset, e.g., for when an existing
	// tileset is updated.
	TilesetVersion *string `json:"tilesetVersion,omitempty"`
}
