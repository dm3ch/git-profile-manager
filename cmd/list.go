package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List git profiles",
	Run: func(cmd *cobra.Command, args []string) {
		configDir, err := GetConfigDirAbsolutePath()
		if err != nil {
			fmt.Println("Can't get configuration directory absolute path:")
			fmt.Println(err)
			os.Exit(1)
		}

		profiles, err := filepath.Glob(getProfilePath(configDir, "*"))
		if err != nil {
			fmt.Println("Can't list profiles:")
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Existing profiles:")
		for _, profile := range profiles {
			profileName := filepath.Base(profile)
			profileName = profileName[:len(profileName)-len(profileExtention)-1]
			fmt.Println(profileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
