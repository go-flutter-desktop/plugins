package file_picker

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/errors"
)

const channelName = "plugins.flutter.io/file_picker"

type FilePickerPlugin struct {
	VendorName      string
	ApplicationName string
}

var _ flutter.Plugin = &FilePickerPlugin{} // compile-time type check

func (p *FilePickerPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	err := p.guard()
	if err != nil {
		return err
	}

	dialogProvider := dialogProvider{}

	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("openDirectory", p.filePicker(dialogProvider, true))
	channel.HandleFunc("openFile", p.filePicker(dialogProvider, false))

	return nil
}

func (p *FilePickerPlugin) guard() error {
	if p.VendorName == "" {
		return errors.New("FilePickerPlugin.VendorName must be set")
	}
	if p.ApplicationName == "" {
		return errors.New("FilePickerPlugin.ApplicationName must be set")
	}

	return nil
}

func (p *FilePickerPlugin) filePicker(dialog dialog, isDirectory bool) func(arguments interface{}) (reply interface{}, err error) {
	return func(arguments interface{}) (reply interface{}, err error) {
		decodedArgs, ok := arguments.(map[interface{}]interface{})
		if !ok {
			return nil, errors.New("arguments must be encoded in JSON format")
		}
		title, ok := decodedArgs["title"].(string)
		if !ok {
			return nil, errors.New("arguments requires a title parameter with type string")
		}

		directory, _, err := dialog.File(title, "*", isDirectory)
		if err != nil {
			return nil, errors.Wrap(err, "failed to open dialog picker")
		}

		return directory, nil
	}
}
