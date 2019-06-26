package cmd

import (
	"log"
	"os"

	"github.com/dm3ch/git-profile-manager/config"
	"github.com/hashicorp/logutils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	cfgStorage *config.Config //nolint:gochecknoglobals
	cfgFile    string         //nolint:gochecknoglobals
	isDebug    bool           //nolint:gochecknoglobals

	rootCmd = &cobra.Command{ //nolint:gochecknoglobals
		Use:   "git-profile-manager",
		Short: "Allows to add and switch between multiple user profiles in your git repositories",
		Long: `Git Profile Manager allows to add and switch between multiple
user profiles in your git repositories.`,
	}
)

func Init() {
	cobra.OnInitialize(initLogs, initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "~/.gitprofile", "config file")
	rootCmd.PersistentFlags().BoolVarP(&isDebug, "debug", "d", false, "show debug log")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(delCmd)
	rootCmd.AddCommand(useCmd)
	rootCmd.AddCommand(versionCmd)
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func initLogs() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}

	if isDebug {
		filter.MinLevel = logutils.LogLevel("DEBUG")
	}

	log.SetOutput(filter)
}

func initConfig() {
	cfgFile, _ = homedir.Expand(cfgFile)
	cfgStorage = config.NewConfig()

	err := cfgStorage.Load(cfgFile)
	if err != nil {
		log.Println("[ERROR] Cannot load json from", cfgFile, err)
		os.Exit(1)
	}
}
