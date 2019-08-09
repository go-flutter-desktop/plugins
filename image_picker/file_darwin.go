package image_picker

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func fileDialog(title string, fileType string) (string, error) {
	osascript, err := exec.LookPath("osascript")
	if err != nil {
		return "", err
	}

	var filters string
	switch fileType {
	case "image":
		filters = `"PNG", "public.png", "JPEG", "jpg", "public.jpeg"`
	case "video":
		filters = `"MOV","mov","MP4","mp4","AVI","avi","MKV","mkv"`
	default:
		return "", errors.New("unsupported fileType")
	}

	output, err := exec.Command(osascript, "-e", `choose file of type {`+filters+`} with prompt "`+title+`"`).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Printf("go-flutter/plugins/image_picker: file dialog exited with code %d and output `%s`\n", exitError.ExitCode(), string(output))
			return "", nil // user probably canceled or closed the selection window
		}
		return "", errors.Wrap(err, "failed to open file dialog")
	}

	trimmedOutput := strings.TrimSpace(string(output))

	pathParts := strings.Split(trimmedOutput, ":")
	path := string(filepath.Separator) + filepath.Join(pathParts[1:]...)
	return path, nil
}
