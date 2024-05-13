//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package onfilesystemchange

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
)

const eventType = "onFileSystemChange"

// Interrupt is a JSON structure passed to the murex function
type Interrupt struct {
	Path      string
	Operation string
}

func init() {
	evt, err := newWatch()
	events.AddEventType(eventType, evt, err)
	go evt.init()
}

type watch struct {
	watcher *fsnotify.Watcher
	mutex   sync.Mutex
	paths   map[string]string // map of paths indexed by event name
	source  map[string]source // map of blocks indexed by path
}

type source struct {
	name    string
	block   []rune
	fileRef *ref.File
}

func newWatch() (w *watch, err error) {
	w = new(watch)
	w.watcher, err = fsnotify.NewWatcher()
	w.paths = make(map[string]string)
	w.source = make(map[string]source)
	return
}

// Callback returns the block to execute upon a triggered event
func (evt *watch) findCallbackBlock(path string) (source, error) {
	evt.mutex.Lock()

	for {
		for len(path) > 1 && path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}

		source := evt.source[path]
		if len(source.block) > 0 {
			evt.mutex.Unlock()
			return source, nil
		}

		split := strings.Split(path, "/")
		switch len(split) {
		case 0:
			path = "/"
		case 1:
			path = strings.Join(split, "/")
		default:
			path = strings.Join(split[:len(split)-1], "/")
		}
	}

	evt.mutex.Unlock()
	return source{}, fmt.Errorf("cannot locate source for event '%s'. This is probably a bug in murex, please report to https://github.com/lmorg/murex/issues", path)
}

// Add a path to the watch event list
func (evt *watch) Add(name, path string, block []rune, fileRef *ref.File) error {
	if len(path) == 0 {
		return errors.New("no path to watch supplied")
	}

	for len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	pwd, err := os.Getwd()
	if err == nil && path[0] != '/' {
		path = pwd + "/" + path
	}

	path = filepath.Clean(path)

	err = evt.watcher.Add(path)
	if err == nil {
		evt.mutex.Lock()
		evt.paths[name] = path
		evt.source[path] = source{
			name:    name,
			block:   block,
			fileRef: fileRef,
		}
		evt.mutex.Unlock()
	}

	return err
}

// Remove a path to the watch event list
func (evt *watch) Remove(name string) error {
	path := evt.paths[name]
	if path == "" {
		return fmt.Errorf("no event found for this listener with the name '%s'", name)
	}

	err := evt.watcher.Remove(path)
	if err == nil {
		evt.mutex.Lock()
		delete(evt.paths, name)
		delete(evt.source, path)
		evt.mutex.Unlock()
	}

	return err
}

// Init starts a new watch event loop
func (evt *watch) init() {
	defer evt.watcher.Close()

	for {
		select {
		case event := <-evt.watcher.Events:
			source, err := evt.findCallbackBlock(event.Name)
			if err != nil {
				lang.ShellProcess.Stderr.Writeln([]byte("onFileSystemChange event error: " + err.Error()))
				continue
			}
			
			_, err = events.Callback(
				source.name,
				Interrupt{
					Path:      event.Name,
					Operation: strings.ToLower(event.Op.String()),
				},
				source.block,
				source.fileRef,
				lang.ShellProcess.Stdout,
				lang.ShellProcess.Stderr,
				nil,
				true,
			)
			if err != nil {
				lang.ShellProcess.Stderr.Writeln([]byte("onFileSystemChange event error: " + err.Error()))
				continue
			}

		case err := <-evt.watcher.Errors:
			lang.ShellProcess.Stderr.Writeln([]byte("onFileSystemChange watcher error: " + err.Error()))
		}
	}
}

// Dump returns all the events in fsWatch
func (evt *watch) Dump() map[string]events.DumpT {
	dump := make(map[string]events.DumpT)

	evt.mutex.Lock()

	for name, path := range evt.paths {
		dump[name] = events.DumpT{
			Interrupt: path,
			Block:     string(evt.source[path].block),
			FileRef:   evt.source[path].fileRef,
		}
	}

	evt.mutex.Unlock()

	return dump
}
