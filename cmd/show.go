package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [profile name]",
	Short: "Show git profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]

		configDir, err := GetConfigDirAbsolutePath()
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

		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("Can't open profile %s file:\n", profileName)
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		buffer, err := ioutil.ReadAll(file)
		if err != nil {
			file.Close()
			fmt.Printf("Can't read profile %s file:\n", profileName)
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(string(buffer))
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
