package c3dt

// https://github.com/CesiumGS/3d-tiles/blob/main/specification/schema/metadataEntity.schema.json

// An object containing a reference to a class from a metadata schema, and property values that conform to the properties of that class
type MetadataEntity struct {
	// The class that property values conform to. The value shall be a class ID
	// declared in the `classes` dictionary of the metadata schema
	Class string `json:"class"`

	// A dictionary, where each key corresponds to a property ID in the class'
	// `properties` dictionary and each value contains the property values. The
	// type of the value shall match the property definition: For `BOOLEAN` use
	// `true` or `false`. For `STRING` use a JSON string. For numeric types use
	// a JSON number. For `ENUM` use a valid enum `name`, not an integer value.
	// For `ARRAY`, `VECN`, and `MATN` types use a JSON array containing values
	// matching the `componentType`. Required properties shall be included in
	// this dictionary.
	Properties map[string]any `json:"properties"`
}
