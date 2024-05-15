package cmd

import (
	"fmt"

	"github.com/jeremyhahn/go-viper-cobra-skeleton/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the software version",
	Long:  `Displays software build and version details`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Name:\t\t\t%s\n", app.Name)
		fmt.Printf("Git Tag:\t\t%s\n", app.GitTag)
		fmt.Printf("Git Hash:\t\t%s\n", app.GitHash)
		fmt.Printf("Build User:\t\t%s\n", app.BuildUser)
		fmt.Printf("Build Date:\t\t%s\n", app.BuildDate)
		fmt.Printf("Release:\t\t%s\n", app.Release)
		fmt.Printf("Image:\t\t\t%s\n", app.Image)
	},
}
