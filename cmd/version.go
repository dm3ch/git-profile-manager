package cmd

import (
	"fmt"

	"github.com/dm3ch/git-profile-manager/version"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print tool version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version.VersionNumber)
		fmt.Println("Commit hash:", version.VersionCommitHash)
		fmt.Println("Build date:", version.VersionBuildDate)
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(versionCmd)
}
