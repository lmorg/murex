package parameters

// Parameters is the parameter object
type Parameters struct {
	Params []string
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

	// TokenTypeString: key is a variable. Expand as a single string. eg `command $variable`
	TokenTypeString

	// TokenTypeBlockString: key is a code block. Expand as a single string. eg `command ${ command }`
	TokenTypeBlockString

	// TokenTypeArray: key is an array. Expand as multiple parameters. eg `command @files`
	TokenTypeArray

	// TokenTypeBlockArray: key is a code block. Expand as multiple parameters. eg `command @{ command }`
	TokenTypeBlockArray

	// TokenTypeIndex: key is an array or map. Return only specific indexes. eg `command $variable[index]`
	TokenTypeIndex

	// TokenTypeRange: key is an array. Return only a range. eg `command @variable[start..end]r`
	TokenTypeRange

	// TokenTypeTilde: key is a user name. Return home directory. eg `command ~user`
	TokenTypeTilde
)
