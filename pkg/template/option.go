package template

// OptionType determines how to interact with user
type OptionType string

// Option determines the interaction between the template and the user
type Option struct {
	Name    string
	Type    string
	Title   string
	Example string
}
