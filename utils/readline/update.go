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
	case rl.line.Len() == 0:
		rl.line.Value = r
	case rl.pos == 0:
		rl.line.Value = append(r, rl.line.Value...)
	case rl.pos < rl.line.Len():
		// TODO: this isn't unicode safe
		r := append(r, rl.line.Value[rl.pos:]...)
		rl.line.Value = append(rl.line.Value[:rl.pos], r...)
	default:
		rl.line.Value = append(rl.line.Value, r...)
	}

	rl.moveCursorByAdjust(len(r))
	rl.echo()

	if rl.modeViMode == vimInsert {
		rl.updateHelpers()
	}
}

func (rl *Instance) backspace() {
	if rl.line.Len() == 0 || rl.pos == 0 {
		return
	}

	moveCursorBackwards(1)
	rl.pos--
	rl.delete()
}

func (rl *Instance) delete() {
	switch {
	case rl.line.Len() == 0:
		return
	case rl.pos == 0:
		// TODO: this isn't unicode safe
		rl.line.Value = rl.line.Value[1:]
		rl.echo()
	case rl.pos > rl.line.Len():
		rl.backspace()
	case rl.pos == rl.line.Len():
		// TODO: this isn't unicode safe
		rl.line.Value = rl.line.Value[:rl.pos]
		rl.echo()
	default:
		// TODO: this isn't unicode safe
		rl.line.Value = append(rl.line.Value[:rl.pos], rl.line.Value[rl.pos+1:]...)
		rl.echo()
	}

	rl.updateHelpers()
}
