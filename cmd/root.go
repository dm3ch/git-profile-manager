package cmd

import (
	"fmt"
	"os"

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
	Long:  "Git Profile Manager allows to manage and switch between multiple user profiles in your git configurations",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		configDir, err := GetConfigDirAbsolutePath()
		if err != nil {
			fmt.Println("Can't get configuration directory absolute path:")
			fmt.Println(err)
			os.Exit(1)
		}

		err = CreateDirIfNotExist(configDir)
		if err != nil {
			fmt.Println("Can't create config directory:")
			fmt.Println(err)
			os.Exit(1)
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
