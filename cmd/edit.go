package cmd

import (
	"fmt"
	"os"

	"github.com/dm3ch/git-profile-manager/editor"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [profile name]",
	Short: "Edit git profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]

		configDir, err := getConfigDirAbsolutePath()
		if err != nil {
			fmt.Println("Can't get configuration directory absolute path:")
			fmt.Println(err)
			os.Exit(1)
		}

		path := getProfilePath(configDir, profileName)
		profileExists := isFileExist(path)
		if !profileExists {
			fmt.Printf("Profile %s does not exists\n", profileName)
			os.Exit(1)
		}

		err = editor.NewDefaultEditor(nil).Launch(path)
		if err != nil {
			fmt.Printf("Error while editing profile %s file:\n", profileName)
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
