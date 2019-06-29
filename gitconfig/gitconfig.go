package gitconfig

import (
	"github.com/go-ini/ini"
)

func loadConfig(path string) (*ini.File, error) {
	cfg, err := ini.Load(path)
	return cfg, err
}

func saveConfig(cfg *ini.File, path string) error {
	return cfg.SaveTo(path)
}

func LoadLocalConfig() (*ini.File, error) {
	path, err := getLocalConfigPath()
	if err != nil {
		return nil, err
	}

	return loadConfig(path)
}

func SaveLocalConfig(cfg *ini.File) error {
	path, err := getLocalConfigPath()
	if err != nil {
		return err
	}

	return saveConfig(cfg, path)
}

func LoadGlobalConfig() (*ini.File, error) {
	path, err := getGlobalConfigPath()
	if err != nil {
		return nil, err
	}

	return loadConfig(path)
}

func SaveGlobalConfig(cfg *ini.File) error {
	path, err := getGlobalConfigPath()
	if err != nil {
		return err
	}

	return saveConfig(cfg, path)
}
