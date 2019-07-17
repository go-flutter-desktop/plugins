# path_provider

This Go package implements the host-side of the Flutter [path_provider](https://github.com/flutter/plugins/tree/master/packages/path_provider) plugin.

## Usage

Import as:

```go
import "github.com/go-flutter-desktop/plugins/path_provider"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&path_provider.PathProviderPlugin{
	VendorName:      "myOrganizationOrUsername",
	ApplicationName: "myApplicationName",
}),
```

Change the values of the Vendor and Application names to a custom and unique
string, so it doesn't conflict with other organizations.

## Issues

Please report issues at the [go-flutter issue tracker](https://github.com/go-flutter-desktop/go-flutter/issues/).
