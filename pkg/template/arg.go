package template

// ArgType determine how to interact with user
type ArgType int

// ArgType related constants
const (
	// For input
	ArgTypeInput = iota

	// For checkbox
	ArgTypeCheckbox
)

// Args describe template file
type Args struct {
	Type  ArgType
	Value string
}
