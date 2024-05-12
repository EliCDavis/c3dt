package c3dt

type EnumValueType string

const (
	EnumValueTypeInt8   = "INT8"
	EnumValueTypeUInt8  = "UINT8"
	EnumValueTypeInt16  = "INT16"
	EnumValueTypeUInt16 = "UINT16"
	EnumValueTypeInt32  = "INT32"
	EnumValueTypeUInt32 = "UINT32"
	EnumValueTypeInt64  = "INT64"
	EnumValueTypeUInt64 = "UINT64"
)

type Enum struct {
	RootProperty

	// The name of the enum, e.g. for display purposes.
	Name string `json:"name,omitempty"`

	// The description of the enum
	Description string `json:"description,omitempty"`

	// The type of the integer enum value.
	ValueType EnumValueType `json:"valueType,omitempty"`

	// An array of enum values. Duplicate names or duplicate integer values are
	// not allowed.
	Values []EnumValue `json:"values"`
}
