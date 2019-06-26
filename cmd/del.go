package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:     "del [profile] [key]",
	Aliases: []string{"rm"},
	Short:   "Delete an entry or a profile",
	Long: `Delete an entry from a profile or an entire profile.
Provide a "key" argument to remove only one key from a profile.`,
	Example: `  git-profile del my-profile -> to delete an entire profile
  git-profile del my-profile user.name -> to delete only user.name`,
	Args: cobra.RangeArgs(1, 2),
	Run:  delRun,
}

func delRun(cmd *cobra.Command, args []string) {
	profile := args[0]

	if len(args) == 2 {
		key := args[1]
		if ok := cfgStorage.RemoveValue(profile, key); !ok {
			cmd.Printf("There is no profile with `%s` name", profile)
			os.Exit(0)
		}

		if err := cfgStorage.Save(cfgFile); err != nil {
			cmd.Printf("Can't remove `%s` from `%s` profile.\n", key, profile)
			cmd.Printf(err.Error())
			os.Exit(1)
		}

		cmd.Printf("Successfully removed `%s` from `%s` profile.", key, profile)
		os.Exit(0)
	}

	cfgStorage.RemoveProfile(profile)
	if err := cfgStorage.Save(cfgFile); err != nil {
		cmd.Printf("Can't remove %s profile.\n", profile)
		cmd.Printf(err.Error())
	}
	cmd.Printf("Successfully removed `%s` profile.", profile)
}
