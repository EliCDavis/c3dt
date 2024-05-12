package c3dt

// https://github.com/CesiumGS/3d-tiles/blob/main/specification/schema/content.schema.json

type Content struct {
	// An optional bounding volume that tightly encloses tile content.
	// tile.boundingVolume provides spatial coherence and
	// tile.content.boundingVolume enables tight view frustum culling. When
	// this is omitted, tile.boundingVolume is used.
	BoundingVolume *BoundingVolume `json:"boundingVolume,omitempty"`

	// A uri that points to tile content. When the uri is relative, it is
	// relative to the referring tileset JSON file.
	URI string `json:"uri"`

	// The group this content belongs to. The value is an index into the array
	// of `groups` that is defined for the containing tileset.
	Group *int `json:"group,omitempty"`

	// Metadata that is associated with this content
	Metadata *MetadataEntity `json:"metadata,omitempty"`
}
