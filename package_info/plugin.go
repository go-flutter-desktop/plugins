package package_info

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "plugins.flutter.io/package_info"

// PackageInfoPlugin implements flutter.Plugin and handles method calls to
// the plugins.flutter.io/package_info channel.
type PackageInfoPlugin struct{}

var _ flutter.Plugin = &PackageInfoPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *PackageInfoPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getAll", p.handlePackageInfo)
	return nil
}

func (p *PackageInfoPlugin) handlePackageInfo(arguments interface{}) (reply interface{}, err error) {
	return map[interface{}]interface{}{
		"appName":     flutter.ProjectName,
		"packageName": flutter.OrganizationName + "." + flutter.ProjectName,
		"version":     flutter.ProjectVersion,
		"buildNumber": flutter.ProjectVersion,
	}, nil
}
