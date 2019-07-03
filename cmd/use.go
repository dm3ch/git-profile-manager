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
		global, _ := cmd.Flags().GetBool("global")

		var configType gitconfig.ConfigType
		if global {
			configType = gitconfig.GlobalConfig
		} else {
			configType = gitconfig.LocalConfig
		}

		configDir := getConfigDirRelativePath()

		path := getProfilePath(configDir, profileName)
		profileExists := isFileExist(path)
		if !profileExists {
			fmt.Printf("Profile %s does not exists\n", profileName)
			os.Exit(1)
		}

		out, err := gitconfig.Get(configType, "profile.path")
		if err == nil && out != "" {
			out, err = gitconfig.UnsetAll(configType, "include.path", out[:len(out)-1])
			if err != nil {
				fmt.Printf("git config command error:\n Output: %s\n", out)
				fmt.Println(err)
				os.Exit(1)
			}
		}

		out, err = gitconfig.Add(configType, "include.path", path)
		if err != nil {
			fmt.Printf("git config command error:\n Output: %s\n", out)
			fmt.Println(err)
			os.Exit(1)
		}

		out, err = gitconfig.ReplaceAll(configType, "profile.path", path)
		if err != nil {
			fmt.Printf("git config command error:\n Output: %s\n", out)
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
	useCmd.Flags().BoolP("global", "g", false, "Set profile for global config")
}
