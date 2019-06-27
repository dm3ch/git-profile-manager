package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfigDir = "~/.git-profile-manager"
	envConfigDir     = "GPM_CONFIG_DIR"
)

var rootCmd = &cobra.Command{
	Use:   "git-profile-manager",
	Short: "Allows to manage and switch between multiple git profiles",
	Long: `Git Profile Manager allows to manage and switch between multiple
user profiles in your git configurations`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		configDir := viper.GetString("configDir")

		configDir, err := homedir.Expand(configDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = CreateDirIfNotExist(configDir)
		if err != nil {
			fmt.Println("Can't create config directory")
			fmt.Println(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("configDir", "C", defaultConfigDir,
		fmt.Sprintf("Configuration directory (Could be also set via %s environment variable)", envConfigDir))
	_ = viper.BindPFlag("configDir", rootCmd.PersistentFlags().Lookup("configDir"))
	_ = viper.BindEnv("configDir", envConfigDir)
	viper.SetDefault("configDir", defaultConfigDir)
}

// Create directory if it doesn't exists
func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		return err
	}

	return nil
}
