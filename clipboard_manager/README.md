# clipboard_manager

This Go package implements the host-side of the Flutter [clipboard_manager](https://github.com/anuranBarman/ClipboardManager) plugin.

## Usage

Import as:

```go
import "github.com/go-flutter-desktop/plugins/clipboard_manager"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&clipboard_manager.ClipboardManagerPlugin{}),
```

## Issues

Please report issues at the [go-flutter issue tracker](https://github.com/go-flutter-desktop/go-flutter/issues/).
