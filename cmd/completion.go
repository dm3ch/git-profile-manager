package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion <shell>",
	Short: "Output shell completion code for the specified shell",
	Long: `Output shell completion code for the specified shell
Only bash and zsh supported now.`,
	Example: `# Enable bash completion:
source <(git-profile-manager completion bash)

# Enable zsh completion:
source <(git-profile-manager completion zsh)`,
	Args: cobra.ExactArgs(1),
	// DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		shell := args[0]

		var err error = nil
		switch shell {
		case "bash":
			err = rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			err = rootCmd.GenZshCompletion(os.Stdout)
		default:
			fmt.Printf("Completion for %s shell is not supported", shell)
		}

		if err != nil {
			fmt.Println("Got error while completion generation:")
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
