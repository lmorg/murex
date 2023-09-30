package readline

import (
	"fmt"
	"strings"
)

type previewModeT int

const (
	previewModeClosed       previewModeT = 0
	previewModeOpen         previewModeT = 1
	previewModeAutocomplete previewModeT = 2
)

const previewPromptHSpace = 3

func getPreviewWidth(width int) (preview, forward int) {
	/*switch {
	case width < 5:
		return 0, 0
	case width < 85:
		preview = width - 4
	case width < 105:
		preview = 80
	case width < 120+5:
		preview = width - 4
	default:
		preview = 120
	}*/
	preview = width - 3

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

func (rl *Instance) getPreviewXY() (*PreviewSizeT, error) {
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
		Height:  height - rl.MaxTabCompleterRows - 10, // hintText, multi-line prompts, etc
		Width:   preview,
		Forward: forward,
	}

	return size, nil
}

func (rl *Instance) writePreviewStr() string {
	if rl.previewMode > previewModeClosed && rl.tcr != nil && rl.tcr.Preview != nil {
		size, err := rl.getPreviewXY()
		if err != nil || size.Height < 8 || size.Width < 40 {
			rl.previewCache = nil
			return ""
		}

		item := rl.previewItem
		item = strings.ReplaceAll(item, "\\", "")
		item = strings.TrimSpace(item)

		lines, pos, err := rl.tcr.Preview(rl.line.Runes(), item, rl.PreviewImages, size)

		if err != nil {
			rl.ForceHintTextUpdate(err.Error())
		}
		output, err := rl.previewDrawStr(lines[pos:], size)
		if err != nil {
			rl.previewCache = nil
			return output
		}

		rl.previewCache = &previewCacheT{
			pos:   pos,
			len:   size.Height,
			lines: lines,
			size:  size,
		}

		return output
	}

	rl.previewCache = nil

	return ""
}

const (
	curHome       = "\x1b[H"
	curPosSave    = "\x1b[s"
	curPosRestore = "\x1b[u"
)

func (rl *Instance) previewDrawStr(preview []string, size *PreviewSizeT) (string, error) {
	var output string

	pf := fmt.Sprintf("│%%-%ds│\r\n", size.Width)

	output += curHome

	output += fmt.Sprintf(cursorForwf, size.Forward)
	hr := strings.Repeat("─", size.Width)
	output += "╭" + hr + "╮\r\n"

	for i := 0; i <= size.Height; i++ {
		output += fmt.Sprintf(cursorForwf, size.Forward)

		if i >= len(preview) {
			blank := strings.Repeat(" ", size.Width)
			output += "│" + blank + "│\r\n"
			continue
		}

		output += fmt.Sprintf(pf, preview[i])
	}

	output += fmt.Sprintf(cursorForwf, size.Forward)
	output += "╰" + hr + "╯\r\n"

	output += rl.previewMoveToPromptStr(size)
	return output, nil
}

func (rl *Instance) previewMoveToPromptStr(size *PreviewSizeT) string {
	output := curHome
	output += moveCursorDownStr(size.Height + previewPromptHSpace)
	output += rl.moveCursorFromStartToLinePosStr()
	return output
}

func (rl *Instance) previewPageUpStr() string {
	if rl.previewCache == nil {
		return ""
	}

	rl.previewCache.pos -= rl.previewCache.len
	if rl.previewCache.pos < 0 {
		rl.previewCache.pos = 0
	}

	output, _ := rl.previewDrawStr(rl.previewCache.lines[rl.previewCache.pos:], rl.previewCache.size)
	return output
}

func (rl *Instance) previewPageDownStr() string {
	if rl.previewCache == nil {
		return ""
	}

	rl.previewCache.pos += rl.previewCache.len
	if rl.previewCache.pos > len(rl.previewCache.lines)-rl.previewCache.len-2 {
		rl.previewCache.pos = len(rl.previewCache.lines) - rl.previewCache.len - 2
		if rl.previewCache.pos < 0 {
			rl.previewCache.pos = 0
		}
	}

	output, _ := rl.previewDrawStr(rl.previewCache.lines[rl.previewCache.pos:], rl.previewCache.size)
	return output
}

func (rl *Instance) clearPreviewStr() string {
	var output string

	if rl.previewMode > previewModeClosed {
		output = seqRestoreBuffer
		output += rl.echoStr()
		rl.previewMode = previewModeClosed
	}

	return output
}
