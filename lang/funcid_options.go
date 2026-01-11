package lang

type OptFidList func(*Process) bool

func FidWithIsBackground(include bool) OptFidList {
	return func(p *Process) bool {
		return p.Background.Get() == include
	}
}

func FidWithIsStopped(include bool) OptFidList {
	return func(p *Process) bool {
		select {
		case <-p.HasStopped:
			return include
		default:
			return !include
		}
	}
}

func FidWithIsChildOf(fid uint32, include bool) OptFidList {
	return func(p *Process) bool {

		proc := p.Parent

		for {
			if proc.Id == fid {
				return include
			}

			proc = proc.Parent

			if proc.Id == ShellProcess.Id {
				return !include
			}
		}
	}
}

func FidWithIsFork(include bool) OptFidList {
	return func(p *Process) bool {
		return p.IsFork == include
	}
}
