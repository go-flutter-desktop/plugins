package url_launcher

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/errors"
)

var channelNames = []string{
	"plugins.flutter.io/url_launcher",
	"plugins.flutter.io/url_launcher_macos",
	"plugins.flutter.io/url_launcher_windows",
	"plugins.flutter.io/url_launcher_linux",
}

// ImagePickerPlugin implements flutter.Plugin and handles method calls to
// the plugins.flutter.io/url_launcher channel.
type UrlLauncherPlugin struct{}

var _ flutter.Plugin = &UrlLauncherPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *UrlLauncherPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	for _, channelName := range channelNames {
		channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
		channel.HandleFunc("launch", p.launch)
		channel.HandleFunc("canLaunch", p.canLaunch)

		// Ignored: The plugins doesn't handle WebView.
		// This call will not do anything, because there is no WebView to close.
		channel.HandleFunc("closeWebView", func(_ interface{}) (interface{}, error) { return nil, nil })
	}
	return nil
}

func (p *UrlLauncherPlugin) launch(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})

	url := argsMap["url"].(string)
	if url == "" {
		return nil, errors.New("url is empty")
	}

	useWebView, ok := argsMap["useWebView"].(bool)
	if ok && useWebView == true {
		fmt.Println("go-flutter-desktop/plugins/url_launcher: WebView aren't supported on desktop.")
	}

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = errors.New("Unsupported platform")
	}

	return err == nil, err
}

func (p *UrlLauncherPlugin) canLaunch(arguments interface{}) (reply interface{}, err error) {
	var url string

	argsMap := arguments.(map[interface{}]interface{})
	url = argsMap["url"].(string)
	if url == "" {
		return false, nil
	}

	return true, nil
}
