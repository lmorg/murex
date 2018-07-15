package consts

// Global project-wide constants

const (
	// NamedPipeProcName is the GoFunction name used for murex builtin that outputs from a named pipe
	NamedPipeProcName = "<named-piped>"

	// CmdExec is the GoFunction name used for murex builtin that executes external processes without a TTY
	CmdExec = "exec"

	// CmdPty is the GoFunction name used for murex builtin that executes external processes with a TTY
	CmdPty = "pty"

	// TestTableHeadings is the header line for the `table` test report format
	TestTableHeadings = " Status  Definition Function                                           Line Col. Message"
)
