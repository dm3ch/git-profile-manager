package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dm3ch/git-profile-manager/profile"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	profileExtention = "profile"
)

// Create directory if it doesn't exists
func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		return err
	}

	return nil
}

// Get configuration directory
func GetConfigDirAbsolutePath() (string, error) {
	configDir := viper.GetString("configDir")
	configDir, err := homedir.Expand(configDir)
	if err != nil {
		return "", err
	}

	return configDir, nil
}

// Prompts string with label
func prompt(label string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", label)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Prompt failed:")
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

// Prompts yes or now answer
func promptYesNo(label string) bool {
	answer := prompt(label + " [y/N]")
	return (answer == "y" || answer == "Y")
}

// Get profile file path
func getProfilePath(configDir, profileName string) string {
	return filepath.Join(configDir, profileName+"."+profileExtention)
}

// Check if file exists
func isFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
