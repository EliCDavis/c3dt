package c3dt

// https://github.com/CesiumGS/3d-tiles/blob/main/specification/schema/tile.schema.json

type TileRefinement string

const (
	AddTileRefinement     = "ADD"
	TileRefinementReplace = "REPLACE"
)

// A tile in a 3D Tiles tileset.
type Tile struct {
	// Specifies if additive or replacement refinement is used when traversing
	// the tileset for rendering. This property is required for the root tile
	// of a tileset; it is optional for all other tiles. The default is to
	// inherit from the parent tile.
	Refine TileRefinement `json:"refine"`

	// A floating-point 4x4 affine transformation matrix, stored in
	// column-major order, that transforms the tile's content--i.e., its
	// features as well as content.boundingVolume, boundingVolume, and
	// viewerRequestVolume--from the tile's local coordinate system to the
	// parent tile's coordinate system, or, in the case of a root tile, from
	// the tile's local coordinate system to the tileset's coordinate system.
	// `transform` does not apply to any volume property when the volume is a
	// region, defined in EPSG:4979 coordinates. `transform` scales the
	// `geometricError` by the maximum scaling factor from the matrix.
	Transform []float64 `json:"transform,omitempty"`

	// Metadata about the tile's content and a link to the content. When this
	// is omitted the tile is just used for culling. When this is defined, then
	// `contents` shall be undefined.
	Content *Content `json:"content,omitempty"`

	// An array of contents. When this is defined, then `content` shall be undefined.
	Contents []Content `json:"contents,omitempty"`

	// An array of objects that define child tiles. Each child tile content is
	// fully enclosed by its parent tile's bounding volume and, generally, has
	// a geometricError less than its parent tile's geometricError. For leaf
	// tiles, there are no children, and this property may not be defined.
	Children []Tile `json:"children,omitempty"`

	// The error, in meters, introduced if this tile is rendered and its
	// children are not. At runtime, the geometric error is used to compute
	// screen space error (SSE), i.e., the error measured in pixels.
	GeometricError float64 `json:"geometricError"`

	// The bounding volume that encloses the tile.
	BoundingVolume BoundingVolume `json:"boundingVolume"`

	// Metadata that is associated with this content
	Metadata *MetadataEntity `json:"metadata,omitempty"`

	// Optional bounding volume that defines the volume the viewer shall be
	// inside of before the tile's content will be requested and before the
	// tile will be refined based on geometricError.
	ViewerRequestVolume BoundingVolume `json:"viewerRequestVolume,omitempty"`
}
