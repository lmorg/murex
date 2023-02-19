package readline

import (
	"fmt"
	"strings"
)

func getPreviewWidth(width int) (preview, forward int) {
	preview = width / 3
	forward = preview * 2
	preview += width - (preview + forward)
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
		if err != nil || size.Height < 8 || size.Width < 25 {
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

	rl.ScreenRefresh()
	print(rl.prompt + string(rl.line))
	rl.writeHintText(false)
	rl.writeTabCompletion(false)
	print("\r\n")
}

const (
	curHome       = "\x1b[H"
	curPosSave    = "\x1b[s"
	curPosRestore = "\x1b[u"
)

func previewDraw(preview []string, size *PreviewSizeT) error {
	pf := fmt.Sprintf("│%%-%ds│\r\n", size.Width)

	print(curPosSave + curHome)
	defer func() {
		print(curPosRestore)
	}()

	moveCursorForwards(size.Forward)
	hr := strings.Repeat("─", size.Width)
	print("╭" + hr + "╮\r\n")

	for i := 0; i <= size.Height; i++ {
		moveCursorForwards(size.Forward)

		if i >= len(preview) {
			blank := strings.Repeat(" ", size.Width)
			print("│" + blank + "│\r\n")
			continue
		}

		out := fmt.Sprintf(pf, preview[i])
		print(out)
	}

	moveCursorForwards(size.Forward)
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
