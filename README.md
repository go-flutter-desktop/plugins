# Flutter desktop plugins

This repo contains [go-flutter](https://github.com/go-flutter-desktop/go-flutter) implementations for popular flutter plugins.

## Issues

Please report issues at the [go-flutter issue tracker](https://github.com/go-flutter-desktop/go-flutter/issues/).

## From [flutter/plugins](https://github.com/flutter/plugins)

Some plugins like `shared_preferences` have been implemented in pure Dart and don't need any work done on the go-flutter side.  
You just need to add `shared_preferences_{linux,macos,windows}` to your dependencies in the `pubspec.yaml`.  
(The respective go-flutter plugins have been deprecated and will eventually be removed, because they sometimes interfere with the official implementation and cause confusion).

- [image_picker](image_picker) - Select an image or video from storage. ([pub.dev](https://pub.dev/packages/image_picker))
- [package_info](package_info) - Provides information about an application package. ([pub.dev](https://pub.dev/packages/package_info))
- [url_launcher](url_launcher) - Flutter plugin for launching a URL. ([pub.dev](https://pub.dev/packages/url_launcher))
- [video_player](video_player) - Flutter plugin for playing back video on a Widget surface. ([pub.dev](https://pub.dev/packages/video_player)) (:warning: work-in-progress, needs rewrite)


## From the community
- [file_picker](https://github.com/miguelpruivo/flutter_file_picker) - Select single or multiple file paths using the native file explorer. ([pub.dev](https://pub.dev/packages/file_picker))
- [sqlite](https://github.com/boltomli/go-flutter-plugin-sqlite) - Flutter plugin for SQLite. ([pub.dev](https://pub.dev/packages/sqflite))
- [platform_device_id](https://github.com/BestBurning/platform_device_id) - Query device identifier. ([pub.dev](https://pub.dev/packages/platform_device_id))
- [shutdown_platform](https://github.com/BestBurning/shutdown_platform) - Shutdown the machine. ([pub.dev](https://pub.dev/packages/shutdown_platform))
- [fast_rsa](https://github.com/jerson/flutter-rsa) - RSA for flutter made with golang for fast performance. ([pub.dev](https://pub.dev/packages/fast_rsa))
- [desktop_cursor](https://github.com/Luukdegram/desktop_cursor) - A Flutter desktop plugin to set the shape of the cursor. ([pub.dev](https://pub.dev/packages/desktop_cursor))
- [title_bar](https://github.com/zephylac/title_bar) - Support custom title bar color [go-flutter#177](https://github.com/go-flutter-desktop/go-flutter/issues/177). **Only for osx**
- [clipboard_manager](https://github.com/djpnewton/go_flutter_clipboard_manager) - Flutter plugin for copying text to the clipboard. ([pub.dev](https://pub.dev/packages/clipboard_manager))
- [open_file](https://github.com/jld3103/go-flutter-open_file) - Flutter plugin for opening a file or URI using the default application on the platform. ([pub.dev](https://pub.dev/packages/open_file))
- [firebase_remote_config](https://github.com/jWinterDay/firebase_remote_config) - Flutter plugin for reading firebase remote config ([pub.dev](https://pub.dev/packages/firebase_remote_config))
- [warble](https://github.com/jslater89/warble) - Play audio from assets, files, or in-memory buffers. ([pub.dev](https://pub.dev/packages/warble))
- [systray](https://github.com/sonr-io/systray) - Support for systray menu for desktop flutter apps
- [flutter_image_compress](https://github.com/OpenFlutter/flutter_image_compress) - Compresses image with fast native libraries 

If you have implemented a plugin that you would like to add to this repository,
feel free to open a PR.
