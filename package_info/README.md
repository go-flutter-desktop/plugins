# package_info

This Go package implements the host-side of the Flutter [package_info](https://github.com/flutter/plugins/tree/master/packages/package_info) plugin.

## Usage

Import as:

```go
import "github.com/go-flutter-desktop/plugins/package_info"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&package_info.PackageInfoPlugin{}),
```

This plugin requires go-flutter `v0.30.0`, and the latest version of hover.

## Issues

Please report issues at the [go-flutter issue tracker](https://github.com/go-flutter-desktop/go-flutter/issues/).
