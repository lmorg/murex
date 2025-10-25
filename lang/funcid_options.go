package lang

type OptFidList func(*Process) bool

func FidWithBackground(background bool) OptFidList {
	return func(p *Process) bool {
		return p.Background.Get() == background
	}
}

func FidWithStopped(stopped bool) OptFidList {
	return func(p *Process) bool {
		select {
		case <-p.HasStopped:
			return true && stopped
		default:
			return false && stopped
		}
	}
}

func FidWithIsChildOf(fid uint32, isChild bool) OptFidList {
	return func(p *Process) bool {
		pp := p.Parent
		for pp.Id != ShellProcess.Id {
			if p.Id == fid {
				return true && isChild
			}
			pp = pp.Parent
		}
		return false && isChild
	}
}
