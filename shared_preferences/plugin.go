package shared_preferences

import (
	"fmt"
	"os"
	"path/filepath"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const channelName = "plugins.flutter.io/shared_preferences"

// SharedPreferencesPlugin implements flutter.Plugin and handles method calls to
// the plugins.flutter.io/shared_preferences channel. Preferences are stored
// using leveldb in the users' home directory config location.
type SharedPreferencesPlugin struct {
	// VendorName must be set to a nonempty value. Use company name or a domain
	// that you own. Note that the value must be valid as a cross-platform directory name.
	VendorName string
	// ApplicationName must be set to a nonempty value. Use the unique name for
	// this application. Note that the value must be valid as a cross-platform
	// directory name.
	ApplicationName string

	userConfigFolder string
	db               *leveldb.DB
	codec            plugin.StandardMessageCodec
}

var _ flutter.Plugin = &SharedPreferencesPlugin{} // compile-time type check

// InitPlugin initializes the shared preferences plugin.
func (p *SharedPreferencesPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	fmt.Println("The go-flutter 'shared_preferences' plugin has been deprecated in favor of the official pure Dart implementation. Please see https://github.com/go-flutter-desktop/plugins for more information")
	if p.VendorName == "" {
		// returned immediately because this is likely a programming error
		return errors.New("SharedPreferencesPlugin.VendorName must be set")
	}
	if p.ApplicationName == "" {
		// returned immediately because this is likely a programming error
		return errors.New("SharedPreferencesPlugin.ApplicationName must be set")
	}

	var err error
	p.userConfigFolder, err = os.UserConfigDir()
	if err != nil {
		return errors.Wrap(err, "failed to resolve user config dir")
	}

	// TODO: move into a getDB call which initializes on first use, lower startup latency.
	p.db, err = leveldb.OpenFile(filepath.Join(p.userConfigFolder, p.VendorName, p.ApplicationName, "shared_preferences.leveldb"), nil)
	if err != nil {
		// TODO: when moved into getDB: error shouldn't kill the plugin and thereby the whole app,
		return errors.Wrap(err, "failed to open leveldb for shared_preferences")
	}

	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("setBool", p.handleSet)
	channel.HandleFunc("setDouble", p.handleSet)
	channel.HandleFunc("setInt", p.handleSet)
	channel.HandleFunc("setString", p.handleSet)
	channel.HandleFunc("setStringList", p.handleSet)
	channel.HandleFunc("commit", p.handleCommit)
	channel.HandleFunc("getAll", p.handleGetAll)
	channel.HandleFunc("remove", p.handleRemove)
	channel.HandleFunc("clear", p.handleClear)

	return nil
}

var defaultWriteOptions = &opt.WriteOptions{
	Sync: true,
}

func (p *SharedPreferencesPlugin) handleSet(arguments interface{}) (reply interface{}, err error) {
	key := []byte(arguments.(map[interface{}]interface{})["key"].(string))
	value, err := p.codec.EncodeMessage(arguments.(map[interface{}]interface{})["value"])
	if err != nil {
		return false, errors.Wrap(err, "failed to encode value")
	}
	err = p.db.Put(key, value, defaultWriteOptions)
	if err != nil {
		return false, errors.Wrap(err, "failed to put key/value into db")
	}
	return true, nil
}

func (p *SharedPreferencesPlugin) handleCommit(arguments interface{}) (reply interface{}, err error) {
	// We've been committing the whole time.
	return true, nil
}

func (p *SharedPreferencesPlugin) handleGetAll(arguments interface{}) (reply interface{}, err error) {
	var values = make(map[interface{}]interface{})
	iter := p.db.NewIterator(nil, nil)
	for iter.Next() {
		value, err := p.codec.DecodeMessage(iter.Value())
		if err != nil {
			return nil, errors.Wrap(err, "failed to get value from db")
		}
		values[string(iter.Key())] = value
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		return nil, errors.Wrap(err, "failed to iterate over key/values in db")
	}
	return values, nil
}

func (p *SharedPreferencesPlugin) handleRemove(arguments interface{}) (reply interface{}, err error) {
	key := arguments.(map[interface{}]interface{})["key"].(string)
	err = p.db.Delete([]byte(key), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete key/value from db")
	}
	return true, nil
}

func (p *SharedPreferencesPlugin) handleClear(arguments interface{}) (reply interface{}, err error) {
	iter := p.db.NewIterator(nil, nil)
	for iter.Next() {
		err = p.db.Delete(iter.Key(), defaultWriteOptions)
		if err != nil {
			return nil, errors.Wrap(err, "failed to delete key/value from db")
		}
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		return nil, errors.Wrap(err, "failed to iterate over key/values in db")
	}
	return true, nil
}
