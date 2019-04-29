# Git Profile switcher

[![CircleCI](https://circleci.com/gh/Dm3Ch/git-profile-manager.svg?style=svg)](https://circleci.com/gh/Dm3Ch/git-profile-manager)
[![Go Report Card](https://goreportcard.com/badge/github.com/Dm3Ch/git-profile)](https://goreportcard.com/report/github.com/Dm3Ch/git-profile-manager)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/dotzero/git-profile/blob/master/LICENSE)

Git Profile allows to add and switch between multiple user profiles in your git repositories.

## Installation

If you are OSX user, you can use [Homebrew](http://brew.sh/):

```bash
brew install dm3ch/tap/git-profile-manager
```

### Prebuilt binaries

Download the binary from the [releases](https://github.com/dotzero/git-profile/releases) page and place it in `$PATH` directory.

### Building from source

If your operating system does not have a binary release, but does run Go, you can build from source.

Make sure that you have Go version 1.11 or greater and that your `GOPATH` env variable is set (I recommend setting it to `~/go` if you don't have one).

```bash
go get -u github.com/dm3ch/git-profile-manager
```

The binary will then be installed to `$GOPATH/bin` (or your `$GOBIN`).

## Usage

Add an entry to a profile

```bash
git profile add home user.name dotzero
git profile add home user.email "mail@dotzero.ru"
git profile add home user.signingkey AAAAAAAA
```

List of available profiles

```bash
git profile list
```

Apply the profile to current git repository

```bash
git profile use home

# Under the hood it runs following commands:
# git git config --local user.name dotzero
# git git config --local user.email "mail@dotzero.ru"
# git git config --local user.signingkey AAAAAAAA
```

## License

http://www.opensource.org/licenses/mit-license.php
