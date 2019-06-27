package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/dm3ch/git-profile-manager/profile"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [profile name]",
	Short: "Add new git profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profile := new(profile.Profile)
		profile.Name = args[0]
		profile.User.Name, _ = cmd.Flags().GetString("name")
		profile.User.Email, _ = cmd.Flags().GetString("email")
		profile.User.SigningKey, _ = cmd.Flags().GetString("signingkey")
		force, _ := cmd.Flags().GetBool("force")

		configDir := viper.GetString("configDir")
		configDir, err := homedir.Expand(configDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if profile.User.Name == "" && profile.User.Email == "" && profile.User.SigningKey == "" {
			promptGitUser(&profile.User)
		}

		path := path.Join(configDir, profile.Name+".profile")

		_, err = os.Stat(path)
		profileExists := !os.IsNotExist(err)

		if profileExists && !force {
			answer := prompt("Override existing profile [y/N]")
			force = (answer == "y" || answer == "Y")
		}

		if !profileExists || force {
			err = profile.Save(path, force)
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

// Prompts string with label
func prompt(label string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", label)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Prompt failed")
		fmt.Println(err)
		os.Exit(1)
	}
	return str[:len(str)-1]
}

// Prompts empty GitUser fields
func promptGitUser(user *profile.GitUser) {
	user.Name = prompt("Name")
	user.Email = prompt("Email")
	user.SigningKey = prompt("Signing Key")
}
