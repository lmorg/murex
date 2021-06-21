package lang

import "fmt"

func testStates(p *Process) {
	for _, name := range p.testState {
		p.Tests.mutex.Lock()
		block := p.Tests.stateBlocks[name]
		p.Tests.mutex.Unlock()
		if len(block) == 0 {
			p.Tests.AddResult(&TestProperties{Name: name}, p, TestError, "No test state defined with that name")
			continue
		}

		fork := p.Fork(F_PARENT_VARTABLE | F_BACKGROUND | F_NO_STDIN | F_CREATE_STDOUT | F_CREATE_STDERR)
		fork.Name.Set(fmt.Sprintf("<state_%s> (%s)", name, p.Name.String()))
		_, err := fork.Execute(block)
		if err != nil {
			p.Tests.AddResult(&TestProperties{Name: name}, p, TestError, err.Error())
		}

		stdout, err := fork.Stdout.ReadAll()
		if err != nil {
			p.Tests.AddResult(&TestProperties{Name: name}, p, TestError, "state stdout: "+err.Error())
		} else {
			p.Tests.AddResult(&TestProperties{Name: name}, p, TestState, "state stdout: "+string(stdout))
		}

		stderr, err := fork.Stderr.ReadAll()
		if err != nil {
			p.Tests.AddResult(&TestProperties{Name: name}, p, TestError, "state stderr: "+err.Error())
		} else {
			p.Tests.AddResult(&TestProperties{Name: name}, p, TestState, "state stderr: "+string(stderr))

		}
	}
}
