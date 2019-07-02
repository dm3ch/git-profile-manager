package cmd

import (
	"fmt"
	"os"
	"os/exec"

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

		out, err := exec.Command("git", "config", "--get", "profile.path").CombinedOutput()
		if err == nil {
			out, err = exec.Command("git", "config", "--unset-all", "include.path", string(out[:len(out)-1])).CombinedOutput()
			if err != nil {
				fmt.Printf("git config command error:\n Output: %s\n", out)
				fmt.Println(err)
				os.Exit(1)
			}
		}

		out, err = exec.Command("git", "config", "--add", "include.path", path).CombinedOutput()
		if err != nil {
			fmt.Printf("git config command error:\n Output: %s\n", out)
			fmt.Println(err)
			os.Exit(1)
		}

		out, err = exec.Command("git", "config", "--replace-all", "profile.path", path).CombinedOutput()
		if err != nil {
			fmt.Printf("git config command error:\n Output: %s\n", out)
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
