package c3dt

// A basis for storing extensions and extras.
type RootProperty struct {
	Extensions Extension `json:"extensions,omitempty"`
	Extras     Extras    `json:"extras,omitempty"`
}
