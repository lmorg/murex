package lang

/*
import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi/codes"
	"github.com/lmorg/murex/utils/escape"
	"github.com/lmorg/murex/utils/home"
	"github.com/lmorg/murex/utils/lists"
	"github.com/lmorg/murex/utils/readline"
)

var (
	rxTokenIndex   = regexp.MustCompile(`(.*?)\[(.*?)\]`)
	rxTokenElement = regexp.MustCompile(`(.*?)\[\[(.*?)\]\]`)
	rxTokenRange   = regexp.MustCompile(`(.*?)\[(.*?)\]([bt8ernsi]*)`)
	rlMutex        sync.Mutex
)

const (
	errEmptyArray = "array '@%s' is empty"
	errEmptyRange = "range '@%s' is empty"
)

// PPConfigDefaults returns the following config settings:
//   - strictArrays
//   - expandGlob
func PPConfigDefaults(p *Process) (bool, bool) {
	strictArrays, err := p.Config.Get("proc", "strict-arrays", "bool")
	if err != nil {
		strictArrays = true
	}

	expandGlob, err := p.Config.Get("shell", "expand-glob", "bool")
	if err != nil {
		expandGlob = true
	}
	expandGlob = expandGlob.(bool) && p.Scope.Id == 0 && !lists.Match(GetNoGlobCmds(), p.Name.String())

	return strictArrays.(bool), expandGlob.(bool)
}

// ParseParameters is an internal function to parse parameters
func ParseParameters(p *Process, paramT *parameters.Parameters) error {
	var namedPipeIsParam bool
	params := []string{}

	strictArrays, expandGlob := PPConfigDefaults(p)

	for i := range paramT.Tokens {
		params = append(params, "")

		var tCount bool
		for j := range paramT.Tokens[i] {
			switch paramT.Tokens[i][j].Type {
			case parameters.TokenTypeNil:
				// do nothing

			case parameters.TokenTypeNamedPipe:
				if !namedPipeIsParam {
					continue
				}
				paramT.Tokens[i][j].Type = parameters.TokenTypeValue
				fallthrough

			case parameters.TokenTypeValue:
				params[len(params)-1] += paramT.Tokens[i][j].Key
				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeGlob:
				if !expandGlob || p.Parent.Id != ShellProcess.Id || p.Background.Get() || !Interactive {
					params[len(params)-1] += paramT.Tokens[i][j].Key

				} else {
					match, globErr := filepath.Glob(paramT.Tokens[i][j].Key)
					glob, err := autoGlobPrompt(p.Name.String(), paramT.Tokens[i][j].Key, match)
					if err != nil {
						return err
					}
					if glob {
						if globErr != nil {
							return fmt.Errorf("invalid glob: '%s'\n%s", paramT.Tokens[i][j].Key, err.Error())
						}
						if len(match) == 0 {
							return fmt.Errorf("glob returned zero results.\nglob: '%s'", paramT.Tokens[i][j].Key)
						}
						if !tCount {
							params = params[:len(params)-1]
						}
						params = append(params, match...)
					} else {
						params[len(params)-1] += paramT.Tokens[i][j].Key
					}
				}
				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeVarString:
				s, err := p.Variables.GetString(paramT.Tokens[i][j].Key)
				if err != nil {
					return err
				}
				s = utils.CrLfTrimString(s)
				params[len(params)-1] += s
				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeVarBlockString:
				fork := p.Fork(F_NO_STDIN | F_CREATE_STDOUT | F_PARENT_VARTABLE)
				exitNum, err := fork.Execute([]rune(paramT.Tokens[i][j].Key))
				if err != nil {
					return fmt.Errorf("subshell failed: %s", err.Error())
				}
				if exitNum > 0 && p.RunMode.IsStrict() {
					return fmt.Errorf("subshell exit status %d", exitNum)
				}
				b, err := fork.Stdout.ReadAll()
				if err != nil {
					return err
				}

				b = utils.CrLfTrim(b)

				params[len(params)-1] += string(b)
				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeVarArray:
				data, err := p.Variables.GetString(paramT.Tokens[i][j].Key)
				if err != nil {
					return err
				}

				if data == "" {
					if strictArrays {
						return fmt.Errorf(errEmptyArray, paramT.Tokens[i][j].Key)
					} else {
						continue
					}
				}

				var array []string

				variable := streams.NewStdin()
				variable.SetDataType(p.Variables.GetDataType(paramT.Tokens[i][j].Key))
				variable.Write([]byte(data))

				variable.ReadArray(p.Context, func(b []byte) {
					array = append(array, string(b))
				})

				if len(array) == 0 && strictArrays {
					return fmt.Errorf(errEmptyArray, paramT.Tokens[i][j].Key)
				}

				if !tCount {
					params = params[:len(params)-1]
				}

				params = append(params, array...)

				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeVarBlockArray:
				var array []string

				fork := p.Fork(F_NO_STDIN | F_CREATE_STDOUT | F_PARENT_VARTABLE)
				exitNum, err := fork.Execute([]rune(paramT.Tokens[i][j].Key))
				if err != nil {
					return fmt.Errorf("subshell failed: %s", err.Error())
				}
				if exitNum > 0 && p.RunMode.IsStrict() {
					return fmt.Errorf("subshell exit status %d", exitNum)
				}
				err = fork.Stdout.ReadArray(p.Context, func(b []byte) {
					array = append(array, string(b))
				})
				if err != nil {
					return err
				}

				if len(array) == 0 && strictArrays {
					return fmt.Errorf(errEmptyArray, "{"+paramT.Tokens[i][j].Key+"}")
				}

				if !tCount {
					params = params[:len(params)-1]
				}

				params = append(params, array...)

				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeVarIndex:
				//debug.Log("parameters.TokenTypeVarIndex:", paramT.Tokens[i][j].Key)
				match := rxTokenIndex.FindStringSubmatch(paramT.Tokens[i][j].Key)
				if len(match) != 3 {
					params[len(params)-1] = paramT.Tokens[i][j].Key
					tCount = true
					continue
				}

				block := []rune("$" + match[1] + "->[" + match[2] + "]")
				fork := p.Fork(F_NO_STDIN | F_CREATE_STDOUT | F_PARENT_VARTABLE)
				fork.Execute(block)
				b, err := fork.Stdout.ReadAll()
				if err != nil {
					return err
				}

				b = utils.CrLfTrim(b)

				params[len(params)-1] += string(b)
				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeVarElement:
				//debug.Log("parameters.TokenTypeVarIndex:", paramT.Tokens[i][j].Key)
				match := rxTokenElement.FindStringSubmatch(paramT.Tokens[i][j].Key)
				if len(match) != 3 {
					params[len(params)-1] = paramT.Tokens[i][j].Key
					tCount = true
					continue
				}

				block := []rune("$" + match[1] + "->[[" + match[2] + "]]")
				fork := p.Fork(F_NO_STDIN | F_CREATE_STDOUT | F_PARENT_VARTABLE)
				fork.Execute(block)
				b, err := fork.Stdout.ReadAll()
				if err != nil {
					return err
				}

				b = utils.CrLfTrim(b)

				params[len(params)-1] += string(b)
				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeVarRange:
				//debug.Log("parameters.TokenTypeVarRange:", paramT.Tokens[i][j].Key)
				match := rxTokenRange.FindStringSubmatch(paramT.Tokens[i][j].Key)
				//debug.Json("parameters.TokenTypeVarRange:", match)

				var flags string

				switch len(match) {
				case 3:
					// do nothing
				case 4:
					flags = match[3]
				default:
					params[len(params)-1] = paramT.Tokens[i][j].Key
					tCount = true
					continue
				}

				var array []string
				block := []rune("$" + match[1] + "-> @[" + match[2] + "]" + flags)
				debug.Log(string(block))
				fork := p.Fork(F_NO_STDIN | F_CREATE_STDOUT | F_PARENT_VARTABLE)
				fork.Execute(block)
				fork.Stdout.ReadArray(p.Context, func(b []byte) {
					array = append(array, string(b))
				})

				if len(array) == 0 && strictArrays {
					return fmt.Errorf(errEmptyRange, paramT.Tokens[i][j].Key)
				}

				if !tCount {
					params = params[:len(params)-1]
				}

				params = append(params, array...)

				tCount = true
				namedPipeIsParam = true

			case parameters.TokenTypeVarTilde:
				if len(paramT.Tokens[i][j].Key) == 0 {
					params[len(params)-1] += home.MyDir
				} else {
					params[len(params)-1] += (paramT.Tokens[i][j].Key)
				}
				tCount = true
				namedPipeIsParam = true

			default:
				err := fmt.Errorf(
					`unexpected parameter token type (%d) in parsed parameters. Param[%d][%d] == "%s"`,
					paramT.Tokens[i][j].Type, i, j, paramT.Tokens[i][j].Key,
				)
				return err
			}
		}

		if !tCount {
			params = params[:len(params)-1]
		}

	}

	paramT.DefineParsed(params)

	return nil
}

func autoGlobPrompt(cmd string, before string, match []string) (bool, error) {
	rlMutex.Lock()
	defer rlMutex.Unlock() // performance doesn't matter here

	rl := readline.NewInstance()
	prompt := fmt.Sprintf("(%s) Do you wish to expand '%s'? [yN]: ", cmd, before)
	rl.SetPrompt(prompt)
	rl.HintText = func(_ []rune, _ int) []rune { return autoGlobPromptHintText(rl, match) }
	rl.History = new(readline.NullHistory)

	for {
		line, err := rl.Readline()
		if err != nil {
			return false, err
		}

		switch strings.ToLower(line) {
		case "y", "yes":
			return true, nil
		case "", "n", "no":
			return false, nil
		}
	}
}

const (
	warningNoGlobMatch = "Warning: no files match that pattern"
	globExpandsTo      = "Glob expands to: "
)

func autoGlobPromptHintText(rl *readline.Instance, match []string) []rune {
	if len(match) == 0 {
		rl.HintFormatting = codes.FgRed
		return []rune(warningNoGlobMatch)
	}

	slice := make([]string, len(match))
	copy(slice, match)
	escape.CommandLine(slice)
	after := globExpandsTo + strings.Join(slice, ", ")
	return []rune(after)
}
*/
