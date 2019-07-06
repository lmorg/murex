// +build !windows,!plan9

package onfilesystemchange

import (
	"github.com/lmorg/murex/lang/ref"
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/lang"
)

const eventType = "onFileSystemChange"

// Interrupt is a JSONable structure passed to the murex function
type Interrupt struct {
	Path      string
	Operation string
}

func init() {
	evt := newWatch()
	events.AddEventType(eventType, evt)
	go evt.init()
}

type watch struct {
	watcher *fsnotify.Watcher
	error   error
	mutex   sync.Mutex
	paths   map[string]string // map of paths indexed by event name
	blocks  map[string][]rune // map of blocks indexed by path
	refFiles map[string]*ref.File
}

func newWatch() (w *watch) {
	w = new(watch)
	w.watcher, w.error = fsnotify.NewWatcher()
	w.paths = make(map[string]string)
	w.blocks = make(map[string][]rune)
	w.refFiles = make(map[string]*ref.File)

	return
}

// Callback returns the block to execute upon a triggered event
func (evt *watch) findCallbackBlock(path string) (block []rune) {
	evt.mutex.Lock()

	for {
		for len(path) > 1 && path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}

		block = evt.blocks[path]

		if len(block) > 0 {
			break
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
		//debug.Log("path=" + path)
	}

	evt.mutex.Unlock()
	return
}

// Add a path to the watch event list
func (evt *watch) Add(name, path string, block []rune, fileRef *ref.File ) error {
	for len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	pwd, err := os.Getwd()
	if err == nil && path[0] != '/' {
		path = pwd + "/" + path
	}

	err = evt.watcher.Add(path)
	if err == nil {
		evt.mutex.Lock()
		evt.paths[name] = path
		evt.blocks[path] = block
		evt.refFiles[name] = fileRef
		evt.mutex.Unlock()
	}

	return err
}

// Remove a path to the watch event list
func (evt *watch) Remove(name string) error {
	path := evt.paths[name]
	if path == "" {
		return errors.New("No event found for this listener with the name `" + name + "`.")
	}

	err := evt.watcher.Remove(path)
	if err == nil {
		evt.mutex.Lock()
		delete(evt.paths, name)
		delete(evt.blocks, path)
		delete(evt.refFiles, name)
		evt.mutex.Unlock()
	}

	return err
}

func getName(evt *watch, path string) (name string, fileRef *ref.File) {
	var evtPath string
	evt.mutex.Lock()

	for name, evtPath = range evt.paths {
		if path == evtPath {
			fileRef = evt.refFiles[name]
			evt.mutex.Unlock()
			return
		}
	}

	// code shouldn't hit this point anyway
	evt.mutex.Unlock()
	return
}

// Init starts a new watch event loop
func (evt *watch) init() {
	defer evt.watcher.Close()

	for {
		select {
		case event := <-evt.watcher.Events:
			name, module := getName(evt, event.Name)
			events.Callback(
				name,
				Interrupt{
					Path:      event.Name,
					Operation: event.Op.String(),
				},
				evt.findCallbackBlock(event.Name),
				module,
				lang.ShellProcess.Stdout,
			)

		case err := <-evt.watcher.Errors:
			lang.ShellProcess.Stderr.Writeln([]byte("onFileSystemChange event error with watcher: " + err.Error()))
		}
	}
}

// Dump returns all the events in fsWatch
func (evt *watch) Dump() interface{} {
	type jsonable struct {
		Path  string
		Block string
		FileRef *ref.File
	}

	dump := make(map[string]jsonable)

	evt.mutex.Lock()

	for name, path := range evt.paths {
		dump[name] = jsonable{
			Path:  path,
			Block: string(evt.blocks[path]),
			FileRef: evt.refFiles[name],
		}
	}

	evt.mutex.Unlock()

	return dump
}
