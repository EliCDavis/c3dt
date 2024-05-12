package c3dt

// https://github.com/CesiumGS/3d-tiles/blob/main/specification/schema/Schema/enum.value.schema.json

// An enum value.
type EnumValue struct {
	RootProperty

	// The name of the enum value.
	Name string `json:"name"`

	// The description of the enum value.
	Description string `json:"description,omitempty"`

	// The integer enum value.
	Value int `json:"value"`
}
