package ansititle

import (
	"bytes"
	"regexp"
)

// rxAnsiEscape matches ANSI escape sequences that occupy zero visible
// columns on screen. Stripping these before measuring is required for
// any prompt or hint that uses richer output than plain SGR colors.
//
// Alternatives, in order:
//
//  1. OSC (Operating System Command): ESC ] ... (BEL | ESC \)
//     Hyperlinks (OSC 8), window titles (OSC 0/1/2), CWD reporting
//     (OSC 7), iTerm shell-integration marks (OSC 1337), notifications
//     (OSC 9), etc. Both string terminators (0x07 BEL and ESC \) are
//     accepted; the payload may not contain ESC or BEL.
//
//  2. CSI (Control Sequence Introducer): ESC [ params intermediates final
//     Covers SGR (final byte 'm'), cursor moves, clears, scroll regions
//     and friends. ECMA-48 conformant byte ranges: params 0x30-0x3F,
//     intermediates 0x20-0x2F, final 0x40-0x7E.
//
//  3. Single-byte Fe escapes: ESC X where X is 0x40-0x5F
//     Picks up DEC save/restore cursor (ESC 7 / ESC 8 are 0x37/0x38 so
//     handled separately below), index (ESC D), next line (ESC E),
//     reverse index (ESC M), single shifts, etc. Excludes the CSI ([)
//     and OSC (]) introducers, which are consumed by the earlier
//     alternatives via leftmost-first matching.
//
//  4. DEC private cursor save/restore: ESC 7, ESC 8
//     Numeric finals fall outside the Fe range above.
var rxAnsiEscape = regexp.MustCompile(
	`\x1b[\]_][^\x07\x1b]*(?:\x07|\x1b\\)` +
		`|\x1b\[[\x30-\x3f]*[\x20-\x2f]*[\x40-\x7e]` +
		`|\x1b[\x40-\x5f]` +
		`|\x1b[78]`,
)

func sanitize(b []byte) []byte {
	b = rxAnsiEscape.ReplaceAll(b, nil)
	b = bytes.ReplaceAll(b, []byte{'\r'}, nil)
	// replace all control characters with space
	for i := range b {
		if b[i] < 32 || b[i] == 127 {
			b[i] = 32
		}
	}

	return b
}
