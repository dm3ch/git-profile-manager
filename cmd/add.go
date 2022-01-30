package cmd

import (
	"fmt"
	"os"

	"github.com/dm3ch/git-profile-manager/profile"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [profile name]",
	Short: "Add git profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profile := new(profile.Profile)
		profile.Name = args[0]
		profile.User.Name, _ = cmd.Flags().GetString("name")
		profile.User.Email, _ = cmd.Flags().GetString("email")
		profile.User.SigningKey, _ = cmd.Flags().GetString("signingkey")
		force, _ := cmd.Flags().GetBool("force")

		configDir, err := getConfigDirAbsolutePath()
		if err != nil {
			fmt.Println("Can't get configuration directory absolute path:")
			fmt.Println(err)
			os.Exit(1)
		}

		if profile.User.Name == "" && profile.User.Email == "" && profile.User.SigningKey == "" {
			promptGitUser(&profile.User)
		}

		path := getProfilePath(configDir, profile.Name)
		profileExists := isFileExist(path)

		if profileExists && !force {
			force = promptYesNo(fmt.Sprintf("Override existing %s profile", profile.Name))
		}

		if !profileExists || force {
			err = profile.Save(path)
			if err != nil {
				fmt.Printf("Profile %s save failed:\n", profile.Name)
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Printf("Profile %s added successfully\n", profile.Name)
			}
		} else {
			fmt.Printf("Profile %s wasn't added\n", profile.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("name", "n", "", "Git user.name")
	addCmd.Flags().StringP("email", "e", "", "Git user.email")
	addCmd.Flags().StringP("signingkey", "s", "", "Git user.signingkey")
	addCmd.Flags().BoolP("force", "f", false, "Override exitsting profile")
}
