package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var rootCmd = &cobra.Command{
	Use:   "git-profile-manager",
	Short: "Allows to manage and switch between multiple git profiles",
	Long: `Git Profile Manager allows to manage and switch between multiple
user profiles in your git configurations`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
