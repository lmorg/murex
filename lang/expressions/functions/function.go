package functions

type FunctionT struct {
	Command    []rune
	Parameters [][]rune
	NamedPipes []string
	Properties Property
	LineN      int
	ColumnN    int
	Raw        []rune
}

type Property int

const (
	P_NEW_CHAIN = 1 << iota
	P_METHOD
	P_PIPE_OUT
	P_PIPE_ERR
	P_LOGIC_AND
	P_LOGIC_OR
)

func (prop Property) NewChain() bool { return prop&P_NEW_CHAIN != 0 }
func (prop Property) Method() bool   { return prop&P_METHOD != 0 }
func (prop Property) PipeOut() bool  { return prop&P_PIPE_OUT != 0 }
func (prop Property) PipeErr() bool  { return prop&P_PIPE_ERR != 0 }
func (prop Property) LogicAnd() bool { return prop&P_LOGIC_AND != 0 }
func (prop Property) LogicOr() bool  { return prop&P_LOGIC_OR != 0 }
