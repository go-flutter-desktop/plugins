# url_launcher

This Go package implements the host-side of the Flutter [url_launcher](https://github.com/flutter/plugins/tree/master/packages/url_launcher) plugin.

## Usage

Import as:

```go
import "github.com/go-flutter-desktop/plugins/url_launcher"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/blob/68868301742b864b719b31ae51c7ec4b3b642d1a/example/simpleDemo/main.go#L53):

```go
flutter.AddPlugin(&url_launcher.UrlLauncherPlugin{}),
```

Change the values of the Vendor and Application names to a custom and unique
string, so it doesn't conflict with other organizations.

## Issues

Please report issues at the [go-flutter issue tracker](https://github.com/go-flutter-desktop/go-flutter/issues/).
