package template

// OptionType determines how to interact with user
type OptionType string

// Predeinf option types
const (
	// Input option
	OptionTypeInput = "input"

	// Select option
	OptionTypeSelect = "select"
)

// The select status of select option type
const (
	OptionTypeSelectIndex = 0
)

// Option determines the interaction between the template and the user
type Option struct {
	Name    string
	Type    string
	Title   string
	Example string
	Value   string
}
