package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:     "nox [sub]",
	Version: "1.0",
	Long: `A modern development tool.

What nox tool can help you?

Auto generate code:
	You can define your custom tmeplate, release your hands to deal with repeat works.`,
}

// Execute to start program
func Execute() {
	rootCmd.Execute()
}
