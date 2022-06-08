package parameters

import "sync"

// Parameters is the parameter object
type Parameters struct {
	mutex  sync.RWMutex
	params []string
	Tokens [][]ParamToken
}

// ParamToken holds information on each parameter token before it is parsed into a string
type ParamToken struct {
	Type int
	Key  string
}

// Type IDs for tokenised parameters (see ParamToken struct)
const (
	// TokenTypeNil: no parameter set or uninitialised parameter.
	TokenTypeNil = iota

	// TokenTypeNamedPipe: key is a named pipe. eg command <stdio>
	TokenTypeNamedPipe

	// TokenTypeValue: key is not a variable. eg `command "just a parameter"`
	TokenTypeValue

	// TokenTypeGlob: key is a value that supports globbing
	TokenTypeGlob

	// TokenTypeVarString: key is a variable. Expand as a single string. eg `command $variable`
	TokenTypeVarString

	// TokenTypeVarBlockString: key is a code block. Expand as a single string. eg `command ${ command }`
	TokenTypeVarBlockString

	// TokenTypeVarArray: key is an array. Expand as multiple parameters. eg `command @files`
	TokenTypeVarArray

	// TokenTypeVarBlockArray: key is a code block. Expand as multiple parameters. eg `command @{ command }`
	TokenTypeVarBlockArray

	// TokenTypeVarIndex: key is an array or map. Return only specific indexes. eg `command $variable[index]`
	TokenTypeVarIndex

	// TokenTypeVarElement: key is an array or map. Return only specific elements in a nested structure. eg `command $variable[index]`
	TokenTypeVarElement

	// TokenTypeVarRange: key is an array. Return only a range. eg `command @variable[start..end]r`
	TokenTypeVarRange

	// TokenTypeVarTilde: key is a user name. Return home directory. eg `command ~user`
	TokenTypeVarTilde
)
