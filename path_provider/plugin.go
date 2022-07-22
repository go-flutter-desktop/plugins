package path_provider

import (
	"path/filepath"

	"github.com/adrg/xdg"

	"github.com/pkg/errors"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

var channelNames = []string{
	"plugins.flutter.io/path_provider",
	"plugins.flutter.io/path_provider_macos",
}

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
}

var _ flutter.Plugin = &PathProviderPlugin{} // compile-time type check

// InitPlugin initializes the path provider plugin.
func (p *PathProviderPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	if p.VendorName == "" {
		// returned immediately because this is likely a programming error
		return errors.New("PathProviderPlugin.VendorName must be set")
	}
	if p.ApplicationName == "" {
		// returned immediately because this is likely a programming error
		return errors.New("PathProviderPlugin.ApplicationName must be set")
	}

	for _, channelName := range channelNames {
		channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
		channel.HandleFunc("getTemporaryDirectory", p.handleTempDir)
		channel.HandleFunc("getApplicationSupportDirectory", p.handleAppSupportDir)
		channel.HandleFunc("getLibraryDirectory", p.handleLibraryDir) // MacOS only
		channel.HandleFunc("getApplicationDocumentsDirectory", p.handleAppDocumentsDir)
		channel.HandleFunc("getStorageDirectory", p.returnError)           // Android only
		channel.HandleFunc("getExternalCacheDirectories", p.returnError)   // Android only
		channel.HandleFunc("getExternalStorageDirectories", p.returnError) // Android only
		channel.HandleFunc("getDownloadsDirectory", p.handleDownloadsDir)
	}

	return nil
}

func (p *PathProviderPlugin) returnError(arguments interface{}) (reply interface{}, err error) {
	return nil, errors.New("This channel is not supported")
}

func (p *PathProviderPlugin) handleTempDir(arguments interface{}) (reply interface{}, err error) {
	return filepath.Join(xdg.CacheHome, p.VendorName, p.ApplicationName), nil
}

func (p *PathProviderPlugin) handleAppSupportDir(arguments interface{}) (reply interface{}, err error) {
	return filepath.Join(xdg.DataHome, p.VendorName, p.ApplicationName), nil
}

// handleLibraryDir is MacOS only and therefore hardcoded, as it is not specified in the XDG specifications
func (p *PathProviderPlugin) handleLibraryDir(arguments interface{}) (reply interface{}, err error) {
	return "/Library/", nil
}

func (p *PathProviderPlugin) handleAppDocumentsDir(arguments interface{}) (reply interface{}, err error) {
	return filepath.Join(xdg.ConfigHome, p.VendorName, p.ApplicationName), nil
}

// handleDownloadsDir is from the flutter plugin side MacOS only
// (https://github.com/flutter/plugins/blob/8819b219c5ca83a000ae482b9a51b7f1f421845b/packages/path_provider/path_provider_platform_interface/lib/src/method_channel_path_provider.dart#L82)
// but should work out of the box once the restriction is not longer there
func (p *PathProviderPlugin) handleDownloadsDir(arguments interface{}) (reply interface{}, err error) {
	return xdg.UserDirs.Download, nil
}
