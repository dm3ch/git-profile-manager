package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/dm3ch/git-profile-manager/gitconfig"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [profile name]",
	Short: "Use specified profile for current repo",
	Args: func(cmd *cobra.Command, args []string) error {
		unset, _ := cmd.Flags().GetBool("unset")

		if unset {
			if len(args) != 0 {
				return errors.New("no arguments required to unset profile")
			}
		} else {
			if len(args) != 1 {
				return errors.New("reqired exact 1 argument that contains profile name")
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		global, _ := cmd.Flags().GetBool("global")
		unset, _ := cmd.Flags().GetBool("unset")

		var configType gitconfig.ConfigType
		if global {
			configType = gitconfig.GlobalConfig
		} else {
			configType = gitconfig.LocalConfig
		}

		configDir := getConfigDirRelativePath()

		var profileName, path string
		var profileExists bool

		if !unset {
			profileName = args[0]
			path = getProfilePath(configDir, profileName)
			profileExists = isFileExist(path)
			if !profileExists {
				fmt.Printf("Profile %s does not exists\n", profileName)
				os.Exit(1)
			}
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

		if !unset {
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
		}
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
	useCmd.Flags().BoolP("global", "g", false, "Set profile for global config")
	useCmd.Flags().BoolP("unset", "u", false, "Just unset currently used profile")
}
