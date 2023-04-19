package readline

func (rl *Instance) insert(r []rune) {
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
		rl.line.Set(r)

	case rl.line.RunePos() == 0:
		rl.line.Set(append(r, rl.line.Runes()...))

	case rl.line.RunePos() < rl.line.RuneLen():
		value := rl.line.Runes()
		new := append(r, value[rl.line.RunePos():]...)
		new = append(value[:rl.line.RunePos()], new...)
		rl.line.Set(new)

	default:
		rl.line.Set(append(rl.line.Runes(), r...))
	}

	rl.moveCursorByRuneAdjust(len(r))
	rl.echo()

	if rl.modeViMode == vimInsert {
		rl.updateHelpers()
	}
}

func (rl *Instance) backspace() {
	if rl.line.RuneLen() == 0 || rl.line.RunePos() == 0 {
		return
	}

	moveCursorBackwards(1)
	rl.line.SetRunePos(rl.line.RunePos() - 1)
	rl.delete()
}

func (rl *Instance) delete() {
	switch {
	case rl.line.RuneLen() == 0:
		return

	case rl.line.RunePos() == 0:
		rl.line.Set(rl.line.Runes()[1:])
		rl.echo()

	case rl.line.RunePos() > rl.line.RuneLen():
		rl.backspace()

	case rl.line.RunePos() == rl.line.RuneLen():
		rl.line.Set(rl.line.Runes()[:rl.line.RunePos()])
		rl.echo()

	default:
		rl.line.Set(append(rl.line.Runes()[:rl.line.RunePos()], rl.line.Runes()[rl.line.RunePos()+1:]...))
		rl.echo()
	}

	rl.updateHelpers()
}
