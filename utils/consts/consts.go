package consts

// Global project-wide constants

const (
	// NamedPipeProcName is the GoFunction name used for murex builtin that outputs from a named pipe
	NamedPipeProcName = "read-named-pipe"

	// TestTableHeadings is the header line for the `table` test report format
	TestTableHeadings = " Status  Definition Function                                           Line Col.  Message"
)

const (
	EnvTrue  = "true"
	EnvFalse = "false"

	EnvMurexPid   = "MUREX_PID"
	EnvDataType   = "MUREX_DATA_TYPE"
	EnvMethod     = "MUREX_IS_METHOD"
	EnvBackground = "MUREX_IS_BACKGROUND"
)

const IssueTrackerURL = "This is a murex bug. Please raise an issue at https://github.com/lmorg/murex/issues"
