package cmd

import (
	"github.com/dm3ch/git-profile-manager/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "version",
	Short: "Print the version number of Git Profile",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Git Profile")
		cmd.Println("Version:", version.VersionNumber)
		cmd.Println("Commit hash:", version.VersionCommitHash)
		cmd.Println("Build date:", version.VersionBuildDate)
	},
}
