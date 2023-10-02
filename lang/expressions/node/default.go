package node

func InitialiseDefaultTheme() {
	err := DefaultTheme.CompileTheme()
	if err != nil {
		panic(err)
	}
}

var DefaultTheme = ThemeT{
	Command:      "{BOLD}",
	CmdModifier:  "{GREEN}",
	Parameter:    "",
	Glob:         "{GREEN}",
	Number:       "{CYAN}",
	Bareword:     "",
	Boolean:      "{GREEN}",
	Null:         "{GREEN}",
	Variable:     "{GREEN}",
	Macro:        "{GREEN}",
	Escape:       "{GREEN}",
	QuotedString: "{BLUE}",
	ArrayItem:    "{YELLOW}",
	ObjectKey:    "{BLUE}",
	ObjectValue:  "{YELLOW}",
	Operator:     "{MAGENTA}",
	Pipe:         "{MAGENTA}",
	Braces: []string{
		"{CYAN}", "{YELLOW}", "{BLUE}", "{MAGENTA}",
	},
	Comment: "{BG-GREEN}",
	//Comment:    "->",
	//EndComment: "<-",
	Error: "{RED}",
}
