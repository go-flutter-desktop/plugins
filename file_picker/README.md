# dialog_picker

This Go package opens the native directory/file picker.

## Usage

Import as:

```go
import file_picker
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/blob/68868301742b864b719b31ae51c7ec4b3b642d1a/example/simpleDemo/main.go#L53):

```go
flutter.AddPlugin(&path_provider.DialogPickerPlugin{
	VendorName:      "myOrganizationOrUsername",
	ApplicationName: "myApplicationName",
}),
```

Change the values of the Vendor and Application names to a custom and unique
string, so it doesn't conflict with other organizations.

```dart
    static const platform = const MethodChannel('plugins.flutter.io/file_picker');

    try {
      var result = await platform.invokeMethod(
          "openDirectory", <String, dynamic>{"title": "Open a Folder"});
      debugPrint(result);
    } catch (err) {
      debugPrint("Error" + err.toString());
    }
    
    try {
      var result = await platform.invokeMethod(
          "openFile", <String, dynamic>{"title": "Open a File"});
      debugPrint(result);
    } catch (err) {
      debugPrint("Error" + err.toString());
    }
```

Add the above to your application code.

## Running the Tests

1. Perform a [manual install](https://github.com/go-flutter-desktop/go-flutter/wiki/Manual-install-and-usage) of go-flutter in this directory
2. `export LD_LIBRARY_PATH=$PWD`
3. `go test`

## Issues

Please report issues at the [go-flutter issue tracker](https://github.com/go-flutter-desktop/go-flutter/issues/).
