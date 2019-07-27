package clipboard_manager

import (
	"runtime"

	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/errors"
    "github.com/atotto/clipboard"
)

const channelName = "clipboard_manager"

// ClipboardManagerPlugin implements flutter.Plugin and handles method calls to
// the clipboard_manager channel.
type ClipboardManagerPlugin struct{}

var _ flutter.Plugin = &ClipboardManagerPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *ClipboardManagerPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("copyToClipBoard", p.copyToClipBoard)
	return nil
}

func (p *ClipboardManagerPlugin) copyToClipBoard(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})

	text := argsMap["text"].(string)

	switch runtime.GOOS {
	case "linux":
        fallthrough
	case "windows":
        fallthrough
	case "darwin":
        err = clipboard.WriteAll(text)
	default:
		err = errors.New("Unsupported platform")
	}

	return nil, err
}
