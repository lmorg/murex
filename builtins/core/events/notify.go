package events

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"os"
	"strings"
	"sync"
)

var w *watch = newWatch()

func init() {
	proc.GoFunctions["event"] = cmdEvent
	go w.Watch()
}

type watch struct {
	watcher  *fsnotify.Watcher
	error    error
	mutex    sync.Mutex
	cbBlocks map[string][]rune
}

func newWatch() (w *watch) {
	w = new(watch)
	w.watcher, w.error = fsnotify.NewWatcher()
	w.cbBlocks = make(map[string][]rune)

	return
}

// Callback returns the block to execute upon a triggered event
func (w *watch) Callback(path string) (block []rune) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	for {
		for len(path) > 1 && path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}

		block = w.cbBlocks[path]

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
	}

	return
}

// Add a path to the watch event list
func (w *watch) Add(path string, block []rune) (err error) {
	for len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	pwd, err := os.Getwd()
	if err == nil && path[0] != '/' {
		path = pwd + "/" + path
	}

	w.mutex.Lock()
	defer w.mutex.Unlock()

	err = w.watcher.Add(path)
	if err == nil {
		w.cbBlocks[path] = block
	}

	return
}

// Watch starts a new watch event loop
func (w *watch) Watch() {
	defer w.watcher.Close()
	type j struct {
		Object      string
		Event       fsnotify.Op
		Description string
	}

	for {
		select {
		case event := <-w.watcher.Events:
			debug.Log("Event:", event)

			block := w.Callback(event.Name)
			stdin := streams.NewStdin()
			json, err := utils.JsonMarshal(&j{
				Object:      event.Name,
				Event:       event.Op,
				Description: event.Op.String(),
			}, false)
			if err != nil {
				ansi.Stderrln(ansi.FgRed, "error building Event input: "+err.Error())
				continue
			}

			_, err = stdin.Write(json)
			if err != nil {
				ansi.Stderrln(ansi.FgRed, "error writing Event input: "+err.Error())
				continue
			}
			stdin.Close()

			_, err = lang.ProcessNewBlock(block, stdin, proc.ShellProcess.Stdout, proc.ShellProcess.Stderr, proc.ShellProcess)
			if err != nil {
				ansi.Stderrln(ansi.FgRed, "error compiling Event callback: "+err.Error())
				continue
			}

			//if Event.Op&fsnotify.Write == fsnotify.Write {
			//	debug.Log("modified file:", Event.Name)
			//}

		case err := <-w.watcher.Errors:
			ansi.Stderrln(ansi.FgRed, "error in watcher: "+err.Error())
		}
	}
}

func cmdEvent(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if w.error != nil {
		return errors.New("watcher failed to initialise with the following error: " + w.error.Error())
	}

	if p.Parameters.Len() < 2 {
		return errors.New("No paths selected to watch")
	}

	block, err := p.Parameters.Block(p.Parameters.Len() - 1)
	if err != nil {
		return err
	}

	var errs string
	for _, f := range p.Parameters.StringArray()[0 : p.Parameters.Len()-1] {
		err := w.Add(f, block)
		if err != nil {
			errs += " {path: " + f + ", err: " + err.Error() + "}"
		}
	}

	if errs != "" {
		err = errors.New(errs)
	}

	return err
}

// Dump returns all the events in w
func (w watch) Dump() (dump map[string]string) {
	dump = make(map[string]string)
	w.mutex.Lock()
	defer w.mutex.Unlock()
	for s := range w.cbBlocks {
		dump[s] = string(w.cbBlocks[s])
	}
	return
}

// DumpEvents is used for `runtime` to output all the saved events
func DumpEvents() map[string]string {
	return w.Dump()
}
