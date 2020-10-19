package cmd

import (
	"fmt"

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
			tmpl, err := template.NewTemplate(templateFlag, args)
			if err != nil {
				fmt.Printf("New template failed, %s.\n", err)
				return
			}

			err = tmpl.Generate()
			if err != nil {
				fmt.Printf("Generate failed, %s.\n", err)
			}
		},
	}
)

func init() {
	generatorCmd.Flags().StringVarP(&templateFlag, "template", "t", "", "Specific which template should be generated.")
	rootCmd.AddCommand(generatorCmd)
}
