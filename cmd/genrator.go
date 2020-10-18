package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xr-hui/nox/pkg/template"
)

var (
	// Used for flags
	templateFlag string

	generatorCmd = &cobra.Command{
		Use:   "generate",
		Short: "Auto generate codes follow template settings",
		Run: func(cmd *cobra.Command, args []string) {
			template.NewTemplate(templateFlag, args)
		},
	}
)

func init() {
	generatorCmd.Flags().StringVarP(&templateFlag, "template", "t", "", "Specific which template should be generated.")
	rootCmd.AddCommand(generatorCmd)
}
