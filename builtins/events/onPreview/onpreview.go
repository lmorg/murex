package onpreview

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/cache"
	"github.com/lmorg/murex/utils/lists"
	"github.com/lmorg/murex/utils/readline"
)

const eventType = "onPreview"

func init() {
	event := newOnPreview()
	events.AddEventType(eventType, event, nil)
	shell.EventsPreview = event.callback
}

// Interrupt is a JSONable structure passed to the murex function
type Interrupt struct {
	Name        string
	Operation   string
	PreviewItem string
	CmdLine     string
	Width       int
}

type previewEvent struct {
	Key     string
	Block   []rune
	FileRef *ref.File
}

type previewEvents struct {
	events []previewEvent
	//mutex  sync.Mutex
}

func newOnPreview() *previewEvents {
	return new(previewEvents)
}

// Add a command to the onPrompt
func (evt *previewEvents) Add(name, interrupt string, block []rune, fileRef *ref.File) error {
	if err := isValidInterrupt(interrupt); err != nil {
		return err
	}

	//evt.mutex.Lock()

	key := events.CompileInterruptKey(interrupt, name)
	event := previewEvent{
		Key:     key,
		Block:   block,
		FileRef: fileRef,
	}

	i := evt.exists(key)
	if i == doesNotExist {
		evt.events = append(evt.events, event)
		sort.SliceStable(evt.events, func(i, j int) bool {
			return evt.events[i].Key < evt.events[j].Key
		})
	} else {
		evt.events[i] = event
	}

	//evt.mutex.Unlock()

	return nil
}

func (evt *previewEvents) Remove(key string) error {
	//evt.mutex.Lock()
	//defer evt.mutex.Unlock()

	i := evt.exists(key)
	if i != doesNotExist {
		events, err := lists.RemoveOrdered(evt.events, i)
		if err != nil {
			return fmt.Errorf("unable to remove %s: %s", key, err.Error())
		}
		evt.events = events
		return nil
	}

	var success bool
	for _, interrupt := range interrupts {
		newKey := events.CompileInterruptKey(interrupt, key)
		i = evt.exists(newKey)
		if i != doesNotExist {
			events, err := lists.RemoveOrdered(evt.events, i)
			if err != nil {
				return fmt.Errorf("unable to remove %s: %s", newKey, err.Error())
			}
			evt.events = events
			success = true
		}
	}

	if success {
		return nil
	}
	return fmt.Errorf("no %s event found called `%s`", eventType, key)
}

const (
	metaCacheTTL = "CacheTTL"
)

func (evt *previewEvents) callback(
	ctx context.Context, interrupt string, // event
	previewItem string, cmdLine []rune, // meta
	previousLines []string, size *readline.PreviewSizeT, callback readline.PreviewFuncCallbackT, // render
) {
	isValidInterruptDebug(interrupt)

	//evt.mutex.Lock()

	var (
		b, e           []byte
		interruptValue Interrupt
		stdout, stderr stdio.Io
		err            error
		meta           any
		metaMap        map[string]any
		ttl            int
		ok             bool
	)

	for i := range evt.events {
		key := events.GetInterruptFromKey(evt.events[i].Key)
		if key.Interrupt == interrupt {
			dur := time.After(2 * time.Second)

			hash := cache.CreateHash(previewItem, []string{evt.events[i].Key, string(cmdLine)}, evt.events[i].Block)
			if cache.Read(cache.PREVIEW_EVENT, hash, &b) {
				goto callback
			}

			interruptValue = Interrupt{
				Name:        key.Name,
				Operation:   interrupt,
				CmdLine:     string(cmdLine),
				PreviewItem: previewItem,
				Width:       size.Width,
			}
			stdout, stderr = streams.NewStdin(), streams.NewStdin()
			meta = map[string]any{
				metaCacheTTL: cacheTTL,
			}

			meta, err = events.Callback(
				evt.events[i].Key, interruptValue, // event
				evt.events[i].Block, evt.events[i].FileRef, // script
				stdout, stderr, // pipes
				meta, // meta
				true, // background
			)
			b, _ = stdout.ReadAll()
			e, _ = stderr.ReadAll()
			b = append(b, e...)
			if err != nil {
				b = append([]byte(err.Error()), b...)
			}

			metaMap, ok = meta.(map[string]any)
			if !ok {
				b = append([]byte(fmt.Sprintf(
					"!!! error decoding event meta variable: value is %T, expecting a map: %v !!!\n", meta, meta,
				)), b...)
				goto callback
			}
			ttl, ok = metaMap[metaCacheTTL].(int)
			if !ok {
				b = append([]byte(fmt.Sprintf(
					"!!! error decoding event meta variable: value is %T, expecting an %s: %v !!!\n", metaMap[metaCacheTTL], types.Integer, metaMap[metaCacheTTL],
				)), b...)
				goto callback
			}

			cache.Write(cache.PREVIEW_EVENT, hash, b, cache.Seconds(ttl))

		callback:
			lines, err := shell.PreviewParseAppendEvent(previousLines, b, size, key.Name)

			select {
			case <-ctx.Done():
				return
			default:
				callback(lines, -1, err)
				previousLines = lines
				select {
				case <-dur:
					shell.Prompt.ForceHintTextUpdate(fmt.Sprintf("Slow running event completed: %s", key.Name))
				default:
					continue
				}
			}
		}
	}

	//evt.mutex.Unlock()
}

const doesNotExist = -1

func (evt *previewEvents) exists(key string) int {
	//evt.mutex.Lock()

	for i := range evt.events {
		if evt.events[i].Key == key {
			return i
		}
	}

	//evt.mutex.Unlock()

	return doesNotExist
}

func (evt *previewEvents) Dump() map[string]events.DumpT {
	dump := make(map[string]events.DumpT)

	//evt.mutex.Lock()

	for i := range evt.events {
		dump[evt.events[i].Key] = events.DumpT{
			Interrupt: events.GetInterruptFromKey(evt.events[i].Key).Interrupt,
			Block:     string(evt.events[i].Block),
			FileRef:   evt.events[i].FileRef,
		}
	}

	//evt.mutex.Unlock()

	return dump
}
