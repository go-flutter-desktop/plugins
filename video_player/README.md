# video_player

This Go package implements the host-side of the Flutter [video_player](https://github.com/flutter/plugins/tree/master/packages/video_player) plugin.

## Usage

Import as:

```go
import "github.com/go-flutter-desktop/plugins/video_player"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&video_player.VideoPlayerPlugin{}),
```

The plugin uses a third party library to handle video to image decoding,
 [3d0c/gmf](https://github.com/3d0c/gmf), a go
FFmpeg bindings.  
If you have trouble installing the plugin, checkout their [installation](https://github.com/3d0c/gmf#installation) procedure.

## :warning: Disclaimer :warning:

This plugin is available for educational purposes, and the go-flutter team isn't
actively working on it.  
**`Don't use it in production`** nasty bugs can occur
(mostly memory leak).  
The plugin needs a significant rewrite. We are looking for maintainers. Pull Requests are most welcome!

## Issues

Please report issues in the [go-flutter **video_player** issue tracker :warning:](https://github.com/go-flutter-desktop/go-flutter/issues/134).
