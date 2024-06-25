package readline

import "fmt"

const errCannotRecallWord = "Cannot recall word"

func hkFnRecallWord(rl *Instance, i int) {
	tokens, err := rl.getLastLineTokensBySpace()
	if err != nil {
		recallWordErr(rl, err.Error(), i)
		return
	}

	switch {
	case len(tokens) == 0:
		recallWordErr(rl, "last line contained no words", i)

	case i == -1:
		i = len(tokens)

	case i >= 0:
		if i > len(tokens) {
			recallWordErr(rl, "previous line contained fewer words", i)
			return
		}

	default:
		recallWordErr(rl, "invalid recall value", i)
		return

	}

	word := rTrimWhiteSpace(tokens[i-1])
	output := rl.insertStr([]rune(word + " "))
	print(output)
}

// getLastLineTokensBySpace is a method because we might see value in reusing this
func (rl *Instance) getLastLineTokensBySpace() ([]string, error) {
	line, err := rl.History.GetLine(rl.History.Len() - 1)
	if err != nil {
		return nil, fmt.Errorf("empty history")
	}

	tokens, _, _ := tokeniseSplitSpaces([]rune(line), 0)
	return tokens, nil
}

// recallWordErr is a function so the Go compiler can inline it
func recallWordErr(rl *Instance, msg string, i int) {
	rl.ForceHintTextUpdate(fmt.Sprintf("%s %d: %s", errCannotRecallWord, i, msg))
}
