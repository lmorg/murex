package lang

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/readline/v4"
)

// MurexFuncs is a table of murex functions
type MurexFuncs struct {
	mutex sync.Mutex
	fn    map[string]*murexFuncDetails
}

// MurexFuncDetails is the properties for any given murex function
type murexFuncDetails struct {
	Block      []rune
	Summary    string
	Parameters []MurexFuncParam
	FileRef    *ref.File
}

type MurexFuncParam struct {
	Name        string
	DataType    string
	Description string
	Default     string
	HasDefault  bool
	Optional    bool
}

// NewMurexFuncs creates a new table of murex functions
func NewMurexFuncs() *MurexFuncs {
	mf := new(MurexFuncs)
	mf.fn = make(map[string]*murexFuncDetails)

	return mf
}

func funcSummary(block []rune) string {
	var (
		line1   bool
		comment bool
		summary []rune
	)

	for _, r := range block {
		switch {
		case r == '\r':
			continue

		case r == '\n' && !line1:
			line1 = true

		case r == '\n' && (line1 || comment):
			goto exitParser

		case r == '#':
			comment = true
			line1 = true

		case !line1 && (r == '{' || r == ' ' || r == '\t'):
			continue

		case comment && r == '\t':
			summary = append(summary, ' ', ' ', ' ', ' ')

		case comment:
			summary = append(summary, r)

		case line1 && (r == ' ' || r == '\t'):
			continue

		default:
			return ""
		}
	}

exitParser:
	return strings.TrimSpace(string(summary))
}

const ( // function parameter error messages
	fpeUnexpectedWhiteSpace    = "unexpected whitespace character (chr %d) at %d (%d,%d)"
	fpeUnexpectedNewLine       = "unexpected new line at %d (%d,%d)"
	fpeUnexpectedComma         = "unexpected comma at %d (%d,%d)"
	fpeUnexpectedCharacter     = "unexpected character '%s' (chr %d) at %d (%d,%d)"
	fpeUnexpectedColon         = "unexpected colon ':' (chr %d) at %d (%d,%d)"
	fpeUnexpectedQuotationMark = "unexpected quotation mark '\"' (chr %d) at %d (%d,%d)"
	fpeUnexpectedEndSquare     = "unexpected closing square bracket ']' (chr %d) at %d (%d,%d)"
	fpeEofNameStart            = "missing variable name at %d (%d,%d)"
	fpeEofNameRead             = "variable name not terminated with a colon %d (%d,%d)"
	fpeEofTypeStart            = "missing data type %d (%d,%d)"
	fpeEofDescRead             = "missing closing quotation mark on description %d (%d,%d)"
	fpeEofDefaultRead          = "missing closing square bracket on default %d (%d,%d)"
	fpeParameterNoName         = "parameter %d is missing a name"
	fpeParameterNoDataType     = "parameter %d is missing a data type"
	fpeUnexpectedMandatory     = "mandatory parameters (%d) cannot follow optional parameters (%d)"
)

const ( // function parameter contexts
	fpcNameStart = 0
	fpcNameRead  = iota
	fpcTypeStart
	fpcTypeRead
	fpcDescStart
	fpcDescRead
	fpcDescEnd
	fpcDefaultRead
	fpcDefaultEnd
)

// Parse the function parameter and data type block
func ParseMxFunctionParameters(parameters string) ([]MurexFuncParam, error) {
	/*
		function example (
			name: str [Bob] "User name",
			age:  num [100] "How old are you?"
		) { ... }
	*/

	var (
		context int
		counter int
		x, y    = 0, 1
	)

	mfp := make([]MurexFuncParam, 1)

	for i, r := range parameters {
		x++

		switch r {
		case '\r':
			// do nothing

		case '\n':
			switch context {
			case fpcNameStart, fpcDescEnd, fpcDefaultEnd:
				y++
				x = 1
			default:
				return nil, fmt.Errorf(fpeUnexpectedNewLine, i+1, y, x)
			}

		case ' ', '\t':
			switch context {
			case fpcNameRead:
				return nil, fmt.Errorf(fpeUnexpectedWhiteSpace, r, i+1, y, x)
			case fpcTypeRead:
				context++
			case fpcDescRead:
				mfp[counter].Description += " "
			case fpcDefaultRead:
				mfp[counter].Default += " "
			default:
				// do nothing
				continue
			}

		case ':':
			switch context {
			case fpcNameRead:
				context++
			case fpcDescRead:
				mfp[counter].Description += ":"
			case fpcDefaultRead:
				mfp[counter].Default += ":"
			default:
				return nil, fmt.Errorf(fpeUnexpectedColon, r, i+1, y, x)
			}

		case '"':
			switch context {
			case fpcDefaultRead:
				mfp[counter].Default += "\""
			case fpcDescStart, fpcDescRead:
				context++
			case fpcDefaultEnd:
				context = fpcDescRead
			default:
				return nil, fmt.Errorf(fpeUnexpectedQuotationMark, r, i+1, y, x)
			}

		case '[':
			switch context {
			case fpcDescRead:
				mfp[counter].Description += "["
			case fpcDefaultRead:
				mfp[counter].Default += "["
			case fpcDescStart, fpcDescEnd:
				context = fpcDefaultRead
				mfp[counter].HasDefault = true
			}

		case ']':
			switch context {
			case fpcDescRead:
				mfp[counter].Description += "]"
			case fpcDefaultRead:
				context++
			default:
				return nil, fmt.Errorf(fpeUnexpectedEndSquare, r, i+1, y, x)
			}

		case ',':
			switch context {
			case fpcDescRead:
				mfp[counter].Description += ","
			case fpcDefaultRead:
				mfp[counter].Default += ","
			case fpcNameRead:
				mfp[counter].DataType = types.String
				mfp = append(mfp, MurexFuncParam{})
				counter++
				context = fpcNameStart
			case fpcTypeRead, fpcDescEnd, fpcDefaultEnd:
				mfp = append(mfp, MurexFuncParam{})
				counter++
				context = fpcNameStart
			default:
				return nil, fmt.Errorf(fpeUnexpectedComma, i+1, y, x)
			}

		case '!':
			switch context {
			case fpcNameStart:
				context++
				mfp[counter].Optional = true
			case fpcDescRead:
				mfp[counter].Description += string([]rune{r})
			case fpcDefaultRead:
				mfp[counter].Default += string([]rune{r})
			default:
				return nil, fmt.Errorf(fpeUnexpectedCharacter, string([]rune{r}), r, i+1, y, x)
			}

		default:
			if (r >= 'a' && 'z' >= r) ||
				(r >= 'A' && 'Z' >= r) ||
				(r >= '0' && '9' >= r) ||
				r == '_' || r == '-' {

				switch context {
				case fpcNameStart:
					context++
					fallthrough
				case fpcNameRead:
					mfp[counter].Name += string([]rune{r})
					continue
				case fpcTypeStart:
					context++
					fallthrough
				case fpcTypeRead:
					mfp[counter].DataType += string([]rune{r})
					continue
				case fpcDescRead:
					mfp[counter].Description += string([]rune{r})
					continue
				case fpcDefaultRead:
					mfp[counter].Default += string([]rune{r})
					continue
				}
			}

			switch context {
			case fpcDescRead:
				mfp[counter].Description += string([]rune{r})
			case fpcDefaultRead:
				mfp[counter].Default += string([]rune{r})
			default:
				return nil, fmt.Errorf(fpeUnexpectedCharacter, string([]rune{r}), r, i+1, y, x)
			}
		}
	}

	switch context {
	case fpcNameStart:
		return nil, fmt.Errorf(fpeEofNameStart, len(parameters), y, x)
	case fpcNameRead:
		//return nil, fmt.Errorf(fpeEofNameRead, len(parameters), y, x)
		mfp[counter].DataType = types.String
	case fpcTypeStart:
		return nil, fmt.Errorf(fpeEofTypeStart, len(parameters), y, x)
	case fpcDescRead:
		return nil, fmt.Errorf(fpeEofDescRead, len(parameters), y, x)
	case fpcDefaultRead:
		return nil, fmt.Errorf(fpeEofDefaultRead, len(parameters), y, x)
	}

	var optional bool
	for i := range mfp {
		if mfp[i].Name == "" {
			return nil, fmt.Errorf(fpeParameterNoName, i+1)
		}
		if mfp[i].DataType == "" {
			return nil, fmt.Errorf(fpeParameterNoDataType, i+1)
		}
		if mfp[i].Optional {
			optional = true
		} else if optional {
			return nil, fmt.Errorf(fpeUnexpectedMandatory, i+1, i)
		}
	}

	return mfp, nil
}

func (mfd *murexFuncDetails) castParameters(p *Process) error {
	for i := range mfd.Parameters {
		s, err := p.Parameters.String(i)
		if err != nil {
			if p.Background.Get() {
				return fmt.Errorf("cannot prompt for parameters when a function is run in the background: %s", err.Error())
			}

			if mfd.Parameters[i].Optional {
				if mfd.Parameters[i].HasDefault {
					s = mfd.Parameters[i].Default
					goto convertType
				}
				continue
			}

			s, err = mfd.Parameters[i].promptParameters()
			if err != nil {
				return err
			}
		}

	convertType:
		v, err := types.ConvertGoType(s, mfd.Parameters[i].DataType)
		if err != nil {
			return fmt.Errorf("cannot convert parameter %d '%s' to data type '%s'", i+1, s, mfd.Parameters[i].DataType)
		}

		err = p.Variables.Set(p, mfd.Parameters[i].Name, v, mfd.Parameters[i].DataType)
		if err != nil {
			return fmt.Errorf("cannot set function variable: %s", err.Error())
		}
	}

	return nil
}

func (mfp *MurexFuncParam) promptParameters() (string, error) {
	prompt := mfp.Description
	if prompt == "" {
		prompt = fmt.Sprintf("Please enter a value for '%s'", mfp.Name)
	}
	if len(mfp.Default) > 0 {
		prompt += fmt.Sprintf(" [%s]", mfp.Default)
	}
	rl := readline.NewInstance()
	rl.SetPrompt(prompt + ": ")
	rl.History = new(readline.NullHistory)

	s, err := rl.Readline()
	if err != nil {
		return "", err
	}

	if s == "" {
		s = mfp.Default
	}

	return s, nil
}

// Define creates a function
func (mf *MurexFuncs) Define(name string, parameters []MurexFuncParam, block []rune, fileRef *ref.File) {
	summary := funcSummary(block)

	mf.mutex.Lock()
	mf.fn[name] = &murexFuncDetails{
		Block:      block,
		Parameters: parameters,
		FileRef:    fileRef,
		Summary:    summary,
	}

	mf.mutex.Unlock()
}

// get returns the function's details
func (mf *MurexFuncs) get(name string) *murexFuncDetails {
	mf.mutex.Lock()
	fn := mf.fn[name]
	mf.mutex.Unlock()
	return fn
}

// Exists checks if function already created
func (mf *MurexFuncs) Exists(name string) bool {
	mf.mutex.Lock()
	exists := mf.fn[name] != nil
	mf.mutex.Unlock()
	return exists
}

// Block returns function code
func (mf *MurexFuncs) Block(name string) ([]rune, error) {
	mf.mutex.Lock()
	fn := mf.fn[name]
	mf.mutex.Unlock()

	if fn == nil {
		return nil, errors.New("cannot locate function named `" + name + "`")
	}

	return fn.Block, nil
}

// Summary returns functions summary
func (mf *MurexFuncs) Summary(name string) (string, error) {
	mf.mutex.Lock()
	fn := mf.fn[name]
	mf.mutex.Unlock()

	if fn == nil {
		return "", errors.New("cannot locate function named `" + name + "`")
	}

	return fn.Summary, nil
}

// Undefine deletes function from table
func (mf *MurexFuncs) Undefine(name string) error {
	mf.mutex.Lock()

	if mf.fn[name] == nil {
		mf.mutex.Unlock()
		return errors.New("cannot locate function named `" + name + "`")
	}

	delete(mf.fn, name)
	mf.mutex.Unlock()
	return nil
}

// Dump list all murex functions in table
func (mf *MurexFuncs) Dump() any {
	type funcs struct {
		Summary    string
		Parameters []MurexFuncParam
		Block      string
		FileRef    *ref.File
	}

	dump := make(map[string]funcs)

	mf.mutex.Lock()
	for name, fn := range mf.fn {
		dump[name] = funcs{
			Summary:    fn.Summary,
			Parameters: fn.Parameters,
			Block:      string(fn.Block),
			FileRef:    fn.FileRef,
		}
	}
	mf.mutex.Unlock()

	return dump
}

// UpdateMap is used for auto-completions. It takes an existing map and updates it's values rather than copying data
func (mf *MurexFuncs) UpdateMap(m map[string]bool) {
	for name := range mf.fn {
		m[name] = true
	}
}
