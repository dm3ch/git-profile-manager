package gitconfig

import "github.com/mitchellh/go-homedir"

func getGlobalConfigPath() (string, error) {
	return homedir.Expand("~/.gitconfig")
}

func getLocalConfigPath() (string, error) {
	return ".git/config", nil
}
