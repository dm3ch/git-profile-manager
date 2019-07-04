package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove [profile name] ([profile name]...)",
	Short:   "Remove git profile",
	Aliases: []string{"rm"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")

		configDir, err := getConfigDirAbsolutePath()
		if err != nil {
			fmt.Println("Can't get configuration directory absolute path:")
			fmt.Println(err)
			os.Exit(1)
		}

		for i := 0; i < len(args); i++ {
			profileName := args[i]

			path := getProfilePath(configDir, profileName)
			profileExists := isFileExist(path)
			if !profileExists {
				fmt.Printf("Profile %s does not exists\n", profileName)
				return
			}

			yes := force
			if !force {
				yes = promptYesNo(fmt.Sprintf("Remove %s profile", profileName))
			}

			if yes {
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
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolP("force", "f", false, "Remove without confirmation")
}
