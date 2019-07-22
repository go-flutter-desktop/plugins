package image_picker

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/errors"
)

const channelName = "plugins.flutter.io/image_picker"

const (
	methodCallImage    = `pickImage`
	methodCallVideo    = `pickVideo`
	methodCallRetrieve = `retrieve`

	sourceCamera  = 0
	sourceGallery = 1
)

// ImagePickerPlugin implements flutter.Plugin and handles method calls to
// the plugins.flutter.io/image_picker channel.
type ImagePickerPlugin struct{}

var _ flutter.Plugin = &ImagePickerPlugin{} // compile-time type check

// InitPlugin initializes the path provider plugin.
func (p *ImagePickerPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("pickImage", p.handlePickImage)
	channel.HandleFunc("pickVideo", p.handlePickVideo)
	channel.HandleFunc("retrieve", p.handleRetrieve)
	return nil
}

func (p *ImagePickerPlugin) handlePickImage(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	switch argsMap["source"].(int32) {
	case sourceCamera:
		return nil, errors.New("source camera is not yet supported by image_picker desktop plugin")
	case sourceGallery:
		if argsMap["maxWidth"] != nil || argsMap["maxHeight"] != nil {
			return nil, errors.New("maxWidth and maxHeight are not yet supported by image_picker desktop plugin")
		}
		path, err := fileDialog("Select an image", "image")
		if err != nil {
			return nil, errors.Wrap(err, "failed to pick an image")
		}
		if path == "" {
			return nil, nil
		}
		return path, nil
	}
	return
}

func (p *ImagePickerPlugin) handlePickVideo(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	switch argsMap["source"].(int32) {
	case sourceCamera:
		return nil, errors.New("source camera is not yet supported by image_picker desktop plugin")
	case sourceGallery:
		path, err := fileDialog("Select a video", "video")
		if err != nil {
			return nil, errors.Wrap(err, "failed to pick a video")
		}
		if path == "" {
			return nil, nil
		}
		return path, nil
	}
	return
}

func (p *ImagePickerPlugin) handleRetrieve(arguments interface{}) (reply interface{}, err error) {
	// retrieve is an android-only plugin method
	return nil, errors.New("retrieve is not supported by desktop image_picker plugin")
}
