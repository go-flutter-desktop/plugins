// +build !darwin,!linux,!windows

package image_picker

import (
	"github.com/pkg/errors"
)

func fileDialog(title string, fileType string) (string, error) {
	return "", errors.New("platform unsupported")
}
