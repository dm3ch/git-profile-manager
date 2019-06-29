package cmd

import (
	"fmt"
	"os"

	"github.com/dm3ch/git-profile-manager/gitconfig"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [profile name]",
	Short: "Use specified profile for current repo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]
		configDir := getConfigDirRelativePath()

		path := getProfilePath(configDir, profileName)
		profileExists := isFileExist(path)
		if !profileExists {
			fmt.Printf("Profile %s does not exists\n", profileName)
			os.Exit(1)
		}

		gitConfig, err := gitconfig.LoadLocalConfig()
		if err != nil {
			fmt.Println("Can't load git repo config:")
			fmt.Println(err)
			os.Exit(1)
		}

		section, err := gitConfig.NewSection("include")
		if err != nil {
			fmt.Println("Can't get profile section:")
			fmt.Println(err)
			os.Exit(1)
		}

		section.NewKey("path", path)

		gitconfig.SaveLocalConfig(gitConfig)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
