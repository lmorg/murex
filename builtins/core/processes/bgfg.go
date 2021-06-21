package processes

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["bg"] = cmdBackground
	lang.GoFunctions["fg"] = cmdForeground
}

func cmdBackground(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Null)

	var block []rune

	if p.IsMethod {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		block = []rune(string(b))

	} else {
		block, err = p.Parameters.Block(0)
		if err != nil {
			return mkbg(p)
		}
	}

	p.Background.Set(true)
	p.WaitForTermination <- false
	fork := p.Fork(lang.F_FUNCTION | lang.F_BACKGROUND)
	fork.Name = p.Name
	fork.Parameters = p.Parameters
	go fork.Execute(block)

	return nil
}

func updateTree(p *lang.Process, isBackground bool) {
	pTree := p
	for {
		//if pTree.Parent == nil || pTree.Parent.Id == 0 || pTree.Name.String() == `bg` {
		if pTree.Id == 0 || pTree.Name.String() == `bg` {
			break
		}
		pTree = pTree.Parent
		pTree.Background.Set(isBackground)
	}

	pTree = p
	for {
		//if pTree.Next == nil || pTree.Next.Id == p.Parent.Id {
		if pTree.Next.Id == p.Parent.Id {
			pTree.Background.Set(isBackground)
			break
		}
		pTree = pTree.Next
		pTree.Background.Set(isBackground)
	}
}
