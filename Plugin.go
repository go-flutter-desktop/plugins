package plugins

import "github.com/pkg/errors"

type Plugin struct {
	// VendorName must be set to a nonempty value. Use company name or a domain
	// that you own. Note that the value must be valid as a cross-platform directory name.
	VendorName string
	// ApplicationName must be set to a nonempty value. Use the unique name for
	// this application. Note that the value must be valid as a cross-platform
	// directory name.
	ApplicationName string
}

func (plugin *Plugin) Guard() error {
	if plugin.VendorName == "" {
		// returned immediately because this is likely a programming error
		return errors.New("PathProviderPlugin.VendorName must be set")
	}
	if plugin.ApplicationName == "" {
		// returned immediately because this is likely a programming error
		return errors.New("PathProviderPlugin.ApplicationName must be set")
	}

	return nil
}