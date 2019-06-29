package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [profile name]",
	Short: "Remove git profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]
		force, _ := cmd.Flags().GetBool("force")

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
			return
		}

		if !force {
			force = promptYesNo(fmt.Sprintf("Remove %s profile", profileName))
		}

		if force {
			err = os.Remove(path)

			if err != nil {
				fmt.Printf("Profile %s remove failed:\n", profileName)
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Printf("Profile %s removed successfully\n", profileName)
			}
		} else {
			fmt.Printf("Profile %s wasn't removed\n", profileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolP("force", "f", false, "Remove without confirmation")
}
