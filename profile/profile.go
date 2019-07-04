package profile

import (
	"github.com/go-ini/ini"
)

type GitUser struct {
	Name       string `ini:"name"`
	Email      string `ini:"email"`
	SigningKey string `ini:"signingkey"`
}

type Profile struct {
	Name string  `ini:"name"`
	User GitUser `ini:"-"`
}

func (profile *Profile) Save(path string) error {
	cfg := ini.Empty()

	profileSection, err := cfg.NewSection("profile")
	if err != nil {
		return err
	}

	err = profileSection.ReflectFrom(&profile)
	if err != nil {
		return err
	}

	userSection, err := cfg.NewSection("user")
	if err != nil {
		return err
	}

	err = userSection.ReflectFrom(&profile.User)
	if err != nil {
		return err
	}

	return cfg.SaveTo(path)
}
