package gitconfig

import (
	"errors"
	"os/exec"
)

type ConfigType int

const (
	LocalConfig ConfigType = iota
	GlobalConfig
	SystemConfig
	MergedConfig
)

func GitExec(command ...string) (string, error) {
	out, err := exec.Command("git", command...).CombinedOutput()

	return string(out), err
}

func Exec(configType ConfigType, command ...string) (string, error) {
	var args []string

	switch configType {
	case LocalConfig:
		args = append([]string{"config", "--local"}, command...)
	case GlobalConfig:
		args = append([]string{"config", "--global"}, command...)
	case SystemConfig:
		args = append([]string{"config", "--system"}, command...)
	case MergedConfig:
		args = append([]string{"config"}, command...)
	default:
		return "", errors.New("can't recognize ConfigType")
	}

	return GitExec(args...)
}

func ReplaceAll(configType ConfigType, key, value string) (string, error) {
	return Exec(configType, "--replace-all", key, value)
}

func Add(configType ConfigType, key, value string) (string, error) {
	return Exec(configType, "--add", key, value)
}

func UnsetAll(configType ConfigType, key, value string) (string, error) {
	if value != "" {
		return Exec(configType, "--unset-all", key, value)
	}

	return Exec(configType, "--unset-all", key)
}

func Get(configType ConfigType, key string) (string, error) {
	return Exec(configType, "--get", key)
}
