package cmd

import (
	"fmt"

	"github.com/dm3ch/git-profile-manager/gitconfig"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current [output template]",
	Short: "Show current git config keys",
	Long:  "Command output git config keys in a format specified by output template (Go template) that was passed.",
	Example: `# Show current use.email
	$ git-profile-manager current "{{ .user.email }}"
	test@test.com
	
	# Show current name and email 
	$ git-profile-manager current "{{ user.name }} ({{ user.email }})"
	Test Name (test@test.com)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tpl := args[0]
		fmt.Println(templateRender(tpl))
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}

func getConfigValue(key string) string {
	out, err := gitconfig.Get(gitconfig.MergedConfig, key)
	if err != nil {
		return ""
	}

	return out[:len(out)-1]
}

//nolint: gocyclo, gocognit, cyclop, gocritic
func templateRender(tpl string) string {
	phStartPos := -1
	phEndPos := 0
	keyStartPos := -1
	keyEndPos := -1
	result := ""

	for index := 0; index < len(tpl); index++ {
		if index != 0 && tpl[index] == '{' && tpl[index-1] == '{' {
			phStartPos = index - 1

			continue
		}

		if phStartPos != -1 && keyStartPos == -1 && tpl[index] != ' ' {
			keyStartPos = index
		}

		if keyStartPos != -1 && keyEndPos == -1 && tpl[index] == ' ' {
			keyEndPos = index
		}

		if index != 0 && tpl[index] == '}' && tpl[index-1] == '}' {
			if phStartPos != -1 {
				result += tpl[phEndPos:phStartPos]

				if keyEndPos == -1 {
					keyEndPos = index - 1
				}

				result += getConfigValue(tpl[keyStartPos:keyEndPos])
				phEndPos = index + 1
				phStartPos = -1
				keyStartPos = -1
				keyEndPos = -1
			}
		}
	}

	result += tpl[phEndPos:]

	return result
}
