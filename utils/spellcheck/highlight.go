package spellcheck

import (
	"fmt"
	"unicode"

	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/inject"
)

type highlightT struct {
	start string
	end   string
}

func (hl *highlightT) Start() []rune { return []rune(ansi.ExpandConsts(hl.start)) }
func (hl *highlightT) End() []rune   { return []rune(ansi.ExpandConsts(hl.end)) }

var highlight *highlightT

func init() {
	highlight = &highlightT{
		start: "{UNDERLINE}",
		end:   "{UNDEROFF}",
	}
}

func highlighter(line *string, word []rune, highlight *highlightT) {
	var (
		i     int
		rLast rune
		r     = []rune(*line)
	)

	defer func() { *line = string(r) }()

	fmt.Printf("'%s', '%s'\n", string(word), *line)

	for ; i < len(r); i++ {
		if r[i] == word[0] {
			if r[i] == word[0] &&
				!(unicode.IsLetter(rLast) ||
					unicode.IsDigit(rLast) ||
					unicode.IsMark(rLast)) {

				// first character of word found

				rStart := i
				for j := 1; j < len(word); j++ {
					i++
					/*fmt.Printf("r[%d]'%s' word[%d]'%s'\n",
					i,
					string([]rune{r[i]}),
					j,
					string([]rune{word[j]}))*/
					switch {
					case i == len(r):
						// end of line reached
						return

					case r[i] != word[j]:
						// word not matched
						break

					case j+1 == len(word):
						// entire word match found
						if i+1 < len(r) && (unicode.IsLetter(r[i+1]) || unicode.IsMark(r[i+1]) || unicode.IsDigit(r[i+1])) {
							// word is substring of a larger word
							break
						} else {
							// add highlight
							var err error
							r, err = inject.Rune(r, highlight.End(), i+1)
							if err != nil {
								return
							}
							r, err = inject.Rune(r, highlight.Start(), rStart)
							if err != nil {
								return
							}
							break
						}

					default:
						continue
					}
				}
			}
		}
		rLast = r[i]
	}
}
