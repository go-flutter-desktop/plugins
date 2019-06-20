package image_picker

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func fileDialog(title string, fileType string) (string, error) {
	cmd, err := exec.LookPath("zenity")
	if err != nil {
		return "", err
	}

	var filters string
	switch fileType {
	case "image":
		filters = `*.png *.jpg *.jpeg`
	case "video":
		filters = `*.webm *.mpeg *.mkv *.mp4 *.avi *.mov *.flv`
	default:
		return "", errors.New("unsupported fileType")
	}

	output, err := exec.Command(cmd, "--file-selection", "--title", title, "--file-filter="+filters).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Printf("go-flutter/plugins/image_picker: file dialog exited with code %d and output `%s`\n", exitError.ExitCode(), string(output))
			return "", nil // user probably canceled or closed the selection window
		}
		return "", errors.Wrap(err, "failed to open file dialog")
	}

	path := strings.TrimSpace(string(output))
	return path, nil
}
