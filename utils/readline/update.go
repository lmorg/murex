package readline

func (rl *Instance) insertStr(r []rune) string {
	for {
		// I don't really understand why `0` is creeping in at the end of the
		// array but it only happens with unicode characters. Also have a similar
		// annoyance with \r when copy/pasting from iTerm.
		if len(r) > 1 && (r[len(r)-1] == 0 || r[len(r)-1] == '\r') {
			r = r[:len(r)-1]
			continue
		}
		break
	}

	switch {
	case rl.line.RuneLen() == 0:
		rl.line.Set(rl, r)

	case rl.line.RunePos() == 0:
		rl.line.Set(rl, append(r, rl.line.Runes()...))

	case rl.line.RunePos() < rl.line.RuneLen():
		value := rl.line.Runes()
		new := append(r, value[rl.line.RunePos():]...)
		new = append(value[:rl.line.RunePos()], new...)
		rl.line.Set(rl, new)

	default:
		rl.line.Set(rl, append(rl.line.Runes(), r...))
	}

	output := rl.moveCursorByRuneAdjustStr(len(r))
	output += rl.echoStr()

	// TODO: check me
	if rl.modeViMode == vimInsert {
		output += rl._updateHelpers()
	}

	return output
}

func (rl *Instance) backspaceStr() string {
	if rl.line.RuneLen() == 0 || rl.line.RunePos() == 0 {
		return ""
	}

	rl.line.SetRunePos(rl.line.RunePos() - 1)
	return rl.deleteStr()
}

func (rl *Instance) deleteStr() string {
	var output string
	switch {
	case rl.line.RuneLen() == 0:
		return ""

	case rl.line.RunePos() == 0:
		rl.line.Set(rl, rl.line.Runes()[1:])
		output = rl.echoStr()

	case rl.line.RunePos() > rl.line.RuneLen():
		output = rl.backspaceStr()
		return output

	case rl.line.RunePos() == rl.line.RuneLen():
		rl.line.Set(rl, rl.line.Runes()[:rl.line.RunePos()])
		output = rl.echoStr()

	default:
		rl.line.Set(rl, append(rl.line.Runes()[:rl.line.RunePos()], rl.line.Runes()[rl.line.RunePos()+1:]...))
		output = rl.echoStr()
	}

	output += rl._updateHelpers()
	return output
}
