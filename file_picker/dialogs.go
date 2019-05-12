package file_picker

import "github.com/gen2brain/dlgs"

type dialog interface {
	File(title, filter string, directory bool) (string, bool, error)
}

type dialogProvider struct {}

func (dialog dialogProvider) File(title, filter string, directory bool) (string, bool, error) {
	return dlgs.File(title, filter, directory)
}