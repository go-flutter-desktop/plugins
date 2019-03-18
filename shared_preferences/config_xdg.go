// +build !windows,!darwin

package shared_preferences

import (
	"os"
	"path/filepath"
)

// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

var userSettingFolder string

func init() {
	if os.Getenv("XDG_CONFIG_HOME") != "" {
		userSettingFolder = os.Getenv("XDG_CONFIG_HOME")
	} else {
		userSettingFolder = filepath.Join(os.Getenv("HOME"), ".config")
	}
}
