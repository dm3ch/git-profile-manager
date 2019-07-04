package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/dm3ch/git-profile-manager/gitconfig"
	"github.com/spf13/cobra"
)

type includeType int

const (
	gitIncludePath      = "include.path"
	gitIncludeIfDirPath = "includeif.gitdir:%s.path"
	gitProfilePath      = "profile.path"

	localInclude includeType = iota
	globalInclude
	dirInclude
)

var useCmd = &cobra.Command{
	Use:   "use [profile name]",
	Short: "Use specified profile",
	Args: func(cmd *cobra.Command, args []string) error {
		unset, _ := cmd.Flags().GetBool("unset")

		if unset {
			if len(args) != 0 {
				return errors.New("no arguments required to unset profile")
			}
		} else {
			if len(args) != 1 {
				return errors.New("reqired exact 1 argument that contains profile name")
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		global, _ := cmd.Flags().GetBool("global")
		unset, _ := cmd.Flags().GetBool("unset")
		dir, _ := cmd.Flags().GetString("dir")

		if dir != "" && global {
			fmt.Println("--global and --dir options can't be used together")
			os.Exit(1)
		}

		if dir[len(dir)-1] != '/' {
			dir += "/"
		}

		var profilePath string
		if !unset {
			profileName := args[0]
			profilePath = getProfilePath(getConfigDirRelativePath(), profileName)
			profileExists := isFileExist(profilePath)
			if !profileExists {
				fmt.Printf("Profile %s does not exists\n", profileName)
				os.Exit(1)
			}
		}

		var incType includeType
		if dir != "" {
			incType = dirInclude
		} else if global {
			incType = globalInclude
		} else {
			incType = localInclude
		}

		err := unsetProfileWrapper(incType, dir)
		if !unset {
			err = setProfileWrapper(incType, dir, profilePath)
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
	useCmd.Flags().BoolP("global", "g", false, "Set profile for global config")
	useCmd.Flags().BoolP("unset", "u", false, "Just unset currently used profile")
	useCmd.Flags().StringP("dir", "d", "", "Set profile for all repositories inside specified directory")
}

func unsetProfile(configType gitconfig.ConfigType, unsetProfilePath bool, key, value string) error {
	out, err := gitconfig.UnsetAll(configType, key, value)
	if err != nil {
		return fmt.Errorf("git config command error:\n Output: %s", out)
	}

	if unsetProfilePath {
		out, err = gitconfig.UnsetAll(configType, gitProfilePath, "")
		if err != nil {
			return fmt.Errorf("git config command error:\n Output: %s", out)
		}
	}

	return nil
}

func unsetProfileWrapper(incType includeType, dir string) error {
	var key, value string
	var unsetProfilePath bool
	var err error
	var configType gitconfig.ConfigType

	switch incType {
	case dirInclude:
		configType = gitconfig.GlobalConfig
		unsetProfilePath = false
		key = fmt.Sprintf(gitIncludeIfDirPath, dir)
		value = ""
	case globalInclude:
		configType = gitconfig.GlobalConfig
		unsetProfilePath = true
		key = gitIncludePath
		value, err = gitconfig.Get(configType, gitProfilePath)
		if err != nil || value == "" {
			return nil
		}
		value = value[:len(value)-1]
	case localInclude:
		configType = gitconfig.LocalConfig
		unsetProfilePath = true
		key = gitIncludePath
		value, err = gitconfig.Get(configType, gitProfilePath)
		if err != nil || value == "" {
			return nil
		}
		value = value[:len(value)-1]
	}

	return unsetProfile(configType, unsetProfilePath, key, value)
}

func setProfile(configType gitconfig.ConfigType, setProfilePath bool, key, value string) error {
	out, err := gitconfig.Add(configType, key, value)
	if err != nil {
		return fmt.Errorf("git config command error:\n Output: %s", out)
	}

	if setProfilePath {
		out, err = gitconfig.ReplaceAll(configType, gitProfilePath, value)
		if err != nil {
			return fmt.Errorf("git config command error:\n Output: %s", out)
		}
	}

	return nil
}

func setProfileWrapper(incType includeType, dir, profilePath string) error {
	var key string
	var setProfilePath bool
	var configType gitconfig.ConfigType

	switch incType {
	case dirInclude:
		configType = gitconfig.GlobalConfig
		setProfilePath = false
		key = fmt.Sprintf(gitIncludeIfDirPath, dir)
	case globalInclude:
		configType = gitconfig.GlobalConfig
		setProfilePath = true
		key = gitIncludePath
	case localInclude:
		configType = gitconfig.LocalConfig
		setProfilePath = true
		key = gitIncludePath
	}

	return setProfile(configType, setProfilePath, key, profilePath)
}
