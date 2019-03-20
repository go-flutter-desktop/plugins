package path_provider

import (
	"os"
	"path/filepath"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/errors"
)

const channelName = "plugins.flutter.io/path_provider"

// PathProviderPlugin implements flutter.Plugin and handles method calls to
// the plugins.flutter.io/path_provider channel.
type PathProviderPlugin struct {
	// VendorName must be set to a nonempty value. Use company name or a domain
	// that you own. Note that the value must be valid as a cross-platform directory name.
	VendorName string
	// ApplicationName must be set to a nonempty value. Use the unique name for
	// this application. Note that the value must be valid as a cross-platform
	// directory name.
	ApplicationName string

	codec plugin.StandardMessageCodec
}

var _ flutter.Plugin = &PathProviderPlugin{} // compile-time type check

func (p *PathProviderPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	if p.VendorName == "" {
		// returned immediately because this is likely a programming error
		return errors.New("PathProviderPlugin.VendorName must be set")
	}
	if p.ApplicationName == "" {
		// returned immediately because this is likely a programming error
		return errors.New("PathProviderPlugin.ApplicationName must be set")
	}

	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getTemporaryDirectory", p.handleTempDir)
	channel.HandleFunc("getApplicationDocumentsDirectory", p.handleAppDir)
	return nil
}

func (p *PathProviderPlugin) handleTempDir(arguments interface{}) (reply interface{}, err error) {
	cacheDir, err := os.UserCacheDir()
	return filepath.Join(cacheDir, p.VendorName, p.ApplicationName), err
}

func (p *PathProviderPlugin) handleAppDir(arguments interface{}) (reply interface{}, err error) {
	return filepath.Join(userSettingFolder, p.VendorName, p.ApplicationName), nil
}
