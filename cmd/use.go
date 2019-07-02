package cmd

import (
	"fmt"
	"os"

	"github.com/dm3ch/git-profile-manager/gitconfig"
	"github.com/spf13/cobra"
)

type include struct {
	Paths []string `ini:"path,omitempty,allowshadow"`
}

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

		incSection, err := gitConfig.NewSection("include")
		if err != nil {
			fmt.Println("Can't get profile section:")
			fmt.Println(err)
			os.Exit(1)
		}

		inc := new(include)
		err = incSection.MapTo(inc)
		if err != nil {
			fmt.Println("Can't map include section:")
			fmt.Println(err)
			os.Exit(1)
		}

		// inc.Paths = append(inc.Paths[:1], inc.Paths[1+1:]...)

		err = incSection.ReflectFrom(inc)
		if err != nil {
			fmt.Println("Can't reflect include section:")
			fmt.Println(err)
			os.Exit(1)
		}

		// _, err = section.NewKey("path", path)
		// if err != nil {
		// 	fmt.Println("Can't set include.path key:")
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		// section, err = gitConfig.NewSection("profile")
		// if err != nil {
		// 	fmt.Println("Can't get profile section:")
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		// key, err := section.Key("path", path)
		// if err != nil {
		// 	fmt.Println("Can't set include.path key:")
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		gitconfig.SaveLocalConfig(gitConfig)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
