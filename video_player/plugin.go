package video_player

import (
	"errors"
	"fmt"
	"time"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "flutter.io/videoPlayer"

// VideoPlayerPlugin implements flutter.Plugin and handles the host side of
// the official Dart Video Player plugin for Flutter.
// VideoPlayerPlugin contains multiple players.
type VideoPlayerPlugin struct {
	textureRegistry *flutter.TextureRegistry
	messenger       plugin.BinaryMessenger
	videoPlayers    map[int32]*player
}

var _ flutter.Plugin = &VideoPlayerPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *VideoPlayerPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.messenger = messenger
	p.videoPlayers = make(map[int32]*player)
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("init", func(_ interface{}) (interface{}, error) { return nil, nil })
	channel.HandleFunc("create", p.create)
	channel.HandleFunc("play", p.play)
	channel.HandleFunc("pause", p.pause)
	channel.HandleFunc("position", p.position)
	channel.HandleFunc("dispose", p.dispose)
	channel.CatchAllHandleFunc(warning)
	return nil
}

func warning(methodCall interface{}) (interface{}, error) {
	method := methodCall.(plugin.MethodCall)
	fmt.Println("go-flutter/plugins/video_player   WARNING   MethodCall to '",
		method.Method, "' isn't supported by the Golang video_player",
		"\n -- please refer to: github.com/go-flutter-desktop/go-flutter/issues/134",
		"for more information.")
	return nil, nil
}

// InitPluginTexture is used to create and manage backend textures
func (p *VideoPlayerPlugin) InitPluginTexture(registry *flutter.TextureRegistry) error {
	p.textureRegistry = registry
	return nil
}

func (p *VideoPlayerPlugin) create(arguments interface{}) (reply interface{}, err error) {
	args := arguments.(map[interface{}]interface{})

	if _, ok := args["asset"]; ok {
		return nil, errors.New("only online video and relative path videos are supported")
	}
	texture := p.textureRegistry.NewTexture()

	player := &player{
		uri:     args["uri"].(string),
		texture: texture,
	}

	eventChannel := plugin.NewEventChannel(p.messenger, fmt.Sprintf("flutter.io/videoPlayer/videoEvents%d", texture.ID), plugin.StandardMethodCodec{})
	eventChannel.Handle(player)

	p.videoPlayers[int32(texture.ID)] = player
	texture.Register(player.textureHanler)

	return map[interface{}]interface{}{
		"textureId": texture.ID,
	}, nil
}

func (p *VideoPlayerPlugin) play(arguments interface{}) (reply interface{}, err error) {
	args := arguments.(map[interface{}]interface{})
	p.videoPlayers[args["textureId"].(int32)].play()
	return nil, nil
}

func (p *VideoPlayerPlugin) dispose(arguments interface{}) (reply interface{}, err error) {
	args := arguments.(map[interface{}]interface{})
	p.videoPlayers[args["textureId"].(int32)].dispose()
	return nil, nil
}

func (p *VideoPlayerPlugin) pause(arguments interface{}) (reply interface{}, err error) {
	args := arguments.(map[interface{}]interface{})
	player := p.videoPlayers[args["textureId"].(int32)]
	player.videoBuffer.Pause()
	return nil, nil
}

func (p *VideoPlayerPlugin) position(arguments interface{}) (reply interface{}, err error) {
	args := arguments.(map[interface{}]interface{})
	videoPlayer := p.videoPlayers[args["textureId"].(int32)]

	return int64(videoPlayer.currentTime * 1000), nil
}

// player correspond to one instance of a player with his associated texture
// handler, eventChannel, methodChannel handler
type player struct {
	uri         string
	videoBuffer *ffmpegVideo
	texture     flutter.Texture
	eventSink   *plugin.EventSink

	// Keep the frame in-sync
	newFrame chan bool

	isStreaming bool
	currentTime float64
}

func (p *player) OnListen(arguments interface{}, sink *plugin.EventSink) { // flutter.EventChannel interface
	p.eventSink = sink

	p.newFrame = make(chan bool, 2)
	p.videoBuffer = &ffmpegVideo{}

	bufferSize := 10 // in frames

	err := p.videoBuffer.Init(p.uri, bufferSize)
	if err != nil {
		sink.Error("VideoError", fmt.Sprintf("Video player had error: %v", err), nil)
	}

	vWidth, vHeight := p.videoBuffer.Bounds()
	sink.Success(map[interface{}]interface{}{
		"event":    "initialized",
		"duration": int64(p.videoBuffer.Duration() * 1000),
		"width":    int32(vWidth),
		"height":   int32(vHeight),
	})
	sink.EndOfStream() // !not a complete implementation
}
func (p *player) OnCancel(arguments interface{}) {} // flutter.EventChannel interface

func (p *player) dispose() {
	p.videoBuffer.Cancel()
	p.texture.UnRegister()
	for len(p.videoBuffer.Frames) > 0 {
		pixels := <-p.videoBuffer.Frames // get the frame, ! Block the main thread !
		pixels.Free()
	}
}

func (p *player) play() {
	if p.isStreaming {
		p.videoBuffer.UnPause()
		return
	}
	p.isStreaming = true

	consumer := func() {
		imagePerSec := p.videoBuffer.GetFrameRate()

		for p.videoBuffer.HasFrameAvailable() {
			p.videoBuffer.WaitUnPause()
			time.Sleep(time.Duration(imagePerSec*1000) * time.Millisecond)
			p.videoBuffer.WaitUnPause()
			p.newFrame <- true
			p.texture.FrameAvailable() // trigger p.textureHanler (display new frame)
		}
	}

	// on the pending frames, consume image in the channel
	go func() {
		p.videoBuffer.Stream(consumer)
		close(p.newFrame)
		p.videoBuffer.Free()
	}()

}

func (p *player) textureHanler(width, height int) (bool, *flutter.PixelBuffer) {
	if p.videoBuffer.Closed() {
		return false, nil
	}

	// Sync frames
	select {
	case <-p.newFrame:
	default:
		// Drop this frame, the event doesn't come from this plugin
		return false, nil
	}

	vWidth, vHeight := p.videoBuffer.Bounds()
	pixels := <-p.videoBuffer.Frames // get the frame, ! Block the main thread !
	p.currentTime = pixels.Time()
	defer pixels.Free()
	return true, &flutter.PixelBuffer{ // send the image to the scene
		Pix:    pixels.Data(),
		Width:  vWidth,
		Height: vHeight,
	}
}
