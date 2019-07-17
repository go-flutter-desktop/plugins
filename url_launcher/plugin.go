package url_launcher

import (
	"fmt"

	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
)

const channelName = "plugins.flutter.io/url_launcher"

type UrlLauncherPlugin struct{}

var _ flutter.Plugin = &UrlLauncherPlugin{} // compile-time type check

func (p *UrlLauncherPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	fmt.Println("InitPlugin")
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("launch", p.launch)
	channel.HandleFunc("canLaunch", p.canLaunch)

	// Ignored: The plugins doesn't handle WebView.
	// This call will not do anything, because there is no WebView to close.
	channel.HandleFunc("closeWebView", func(_ interface{}) (interface{}, error) { return nil, nil })
	return nil
}

func (p *UrlLauncherPlugin) launch(arguments interface{}) (reply interface{}, err error) {
	var url string
	var useWebView bool
	var useSafariVC bool
	var enableJavaScript bool
	var enableDomStorage bool
	var universalLinksOnly bool

	argsMap := arguments.(map[interface{}]interface{})
	url = argsMap["url"].(string)
	if url == "" {
		return nil, errors.New("url is empty")
	}

	useWebView = argsMap["useWebView"].(bool)
	if useWebView == true {
		fmt.Println("plugins.flutter.io/url_launcher: WebView aren't supported on desktop.")
	}

	useSafariVC = argsMap["useSafariVC"].(bool)
	if useSafariVC == true {
		fmt.Println("plugins.flutter.io/url_launcher: SafariVC aren't supported on desktop.")
	}

	enableJavaScript = argsMap["enableJavaScript"].(bool)
	if enableJavaScript == true {
		fmt.Println("plugins.flutter.io/url_launcher: enableJavaScript aren't supported on desktop.")
	}

	enableDomStorage = argsMap["enableDomStorage"].(bool)
	if enableDomStorage == true {
		fmt.Println("plugins.flutter.io/url_launcher: enableDomStorage aren't supported on desktop.")
	}

	universalLinksOnly = argsMap["universalLinksOnly"].(bool)
	if universalLinksOnly == true {
		fmt.Println("plugins.flutter.io/url_launcher: universalLinksOnly aren't supported on desktop.")
	}

	browser.OpenURL(url)

	return nil, nil
}

func (p *UrlLauncherPlugin) canLaunch(arguments interface{}) (reply interface{}, err error) {
	var url string

	argsMap := arguments.(map[interface{}]interface{})
	url = argsMap["url"].(string)
	if url == "" {
		return false, errors.New("url is empty")
	}

	return true, nil
}
