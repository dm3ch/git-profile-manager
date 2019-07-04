package cmd

import (
	"fmt"

	"github.com/dm3ch/git-profile-manager/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version.VersionNumber)
		fmt.Println("Commit hash:", version.VersionCommitHash)
		fmt.Println("Build date:", version.VersionBuildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
