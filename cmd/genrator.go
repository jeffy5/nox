package cmd

import "github.com/spf13/cobra"

var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "Auto generate codes follow template settings",
}

func init() {
	rootCmd.AddCommand(generatorCmd)
}
