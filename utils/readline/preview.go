package readline

import (
	"fmt"
	"os"
	"strings"
)

func getPreviewWidth(width int) (preview, forward int) {
	switch {
	case width < 80:
		return 0, 0
	case width < 90:
		preview = 40
	default:
		preview = 80
	}

	forward = width - preview
	forward -= 2
	return
}

type PreviewSizeT struct {
	Height  int
	Width   int
	Forward int
}

type previewCacheT struct {
	pos   int
	len   int
	lines []string
	size  *PreviewSizeT
}

func getPreviewXY() (*PreviewSizeT, error) {
	width, height, err := GetSize(int(primary.Fd()))
	if err != nil {
		return nil, err
	}

	if height == 0 {
		height = 25
	}

	if width == 0 {
		width = 80
	}

	preview, forward := getPreviewWidth(width)
	size := &PreviewSizeT{
		Height:  height / 3,
		Width:   preview,
		Forward: forward,
	}

	return size, nil
}

func (rl *Instance) writePreview(item string) {
	if rl.previewCache != nil {
		// refresh screen if preview written previously and this one empty
		defer func() {
			if rl.previewCache == nil {
				rl.screenRefresh()
			}
		}()
	}

	if rl.ShowPreviews && rl.tcr != nil && rl.tcr.Preview != nil {
		size, err := getPreviewXY()
		if err != nil || size.Height < 8 || size.Width < 40 {
			rl.previewCache = nil
			return
		}

		item = strings.ReplaceAll(item, "\\", "")
		item = strings.TrimSpace(item)

		lines, pos, err := rl.tcr.Preview(rl.line, item, rl.PreviewImages, size)
		if len(lines) == 0 || err != nil {
			rl.previewCache = nil
			return
		}
		err = previewDraw(lines[pos:], size)
		if err != nil {
			rl.previewCache = nil
			return
		}

		rl.previewCache = &previewCacheT{
			pos:   pos,
			len:   size.Height,
			lines: lines,
			size:  size,
		}

		return
	}

	rl.previewCache = nil
}

func (rl *Instance) screenRefresh() {
	if rl.ScreenRefresh == nil {
		return
	}

	/*print := func(s string) {
		_, _ = os.Stdout.WriteString(s)
	}*/

	old := primary
	primary = os.Stdout

	rl.ScreenRefresh()
	/*print(rl.prompt + string(rl.line))
	rl.writeHintText(false)
	rl.writeTabCompletion(false)
	print("\r\n")*/

	primary = old
}

const (
	curHome       = "\x1b[H"
	curPosSave    = "\x1b[s"
	curPosRestore = "\x1b[u"
)

func previewDraw(preview []string, size *PreviewSizeT) error {
	print := func(s string) {
		_, _ = os.Stdout.WriteString(s)
	}

	pf := fmt.Sprintf("│%%-%ds│\r\n", size.Width)

	print(curPosSave + curHome)
	defer func() {
		print(curPosRestore)
	}()

	//moveCursorForwards(size.Forward)
	print(fmt.Sprintf("\x1b[%dC", size.Forward))
	hr := strings.Repeat("─", size.Width)
	print("╭" + hr + "╮\r\n")

	for i := 0; i <= size.Height; i++ {
		//moveCursorForwards(size.Forward)
		print(fmt.Sprintf("\x1b[%dC", size.Forward))

		if i >= len(preview) {
			blank := strings.Repeat(" ", size.Width)
			print("│" + blank + "│\r\n")
			continue
		}

		print(fmt.Sprintf(pf, preview[i]))
	}

	//moveCursorForwards(size.Forward)
	print(fmt.Sprintf("\x1b[%dC", size.Forward))
	print("╰" + hr + "╯\r\n")

	return nil
}

func (rl *Instance) previewPageUp() {
	if rl.previewCache == nil {
		return
	}

	rl.previewCache.pos -= rl.previewCache.len
	if rl.previewCache.pos < 0 {
		rl.previewCache.pos = 0
	}

	_ = previewDraw(rl.previewCache.lines[rl.previewCache.pos:], rl.previewCache.size)
}

func (rl *Instance) previewPageDown() {
	if rl.previewCache == nil {
		return
	}

	rl.previewCache.pos += rl.previewCache.len
	if rl.previewCache.pos > len(rl.previewCache.lines)-rl.previewCache.len-2 {
		rl.previewCache.pos = len(rl.previewCache.lines) - rl.previewCache.len - 2
		if rl.previewCache.pos < 0 {
			rl.previewCache.pos = 0
		}
	}

	_ = previewDraw(rl.previewCache.lines[rl.previewCache.pos:], rl.previewCache.size)
}
