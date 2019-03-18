# shared_preferences

This Go package implements the host-side of the Flutter shared_preferences plugin.

## Usage

Import as:

```go
import "github.com/go-flutter-desktop/plugins/shared_preferences"
```

Then add the following option to your go-flutter application options:

```go
flutter.AddPlugin(&shared_preferences.SharedPreferencesPlugin{
	VendorName:      "myOrganizationOrUsername",
	ApplicationName: "myApplicationName",
})
```

Change the values of the Vendor and Application names to a custom and unique
string, so it doesn't conflict with other organizations.

## Issues

Please report issues at the [go-flutter issue tracker](https://github.com/go-flutter-desktop/go-flutter/issues/).
