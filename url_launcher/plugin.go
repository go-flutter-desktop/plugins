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
	return nil
}

func (p *UrlLauncherPlugin) launch(arguments interface{}) (reply interface{}, err error) {
	fmt.Println("launch")
	var url string

	argsMap := arguments.(map[interface{}]interface{})
	url = argsMap["url"].(string)
	if url == "" {
		return nil, errors.New("url is empty")
	}
	browser.OpenURL(url)
	//ignore 'useSafariVC' , 'useWebView' , 'enableJavaScript' , 'enableDomStorage' , 'universalLinksOnly'

	return nil, nil
}

func (p *UrlLauncherPlugin) canLaunch(arguments interface{}) (reply interface{}, err error) {
	fmt.Println("canLaunch")
	var url string

	argsMap := arguments.(map[interface{}]interface{})
	url = argsMap["url"].(string)
	if url == "" {
		return false, errors.New("url is empty")
	}
	//windows can accept mailto: sms: tel: notes: , ignore this check condition strings.HasPrefix(url, "?:")

	return true, nil
}
