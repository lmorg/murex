package consts

// Global project-wide constants

const (
	// NamePipeProcName is the GoFunction name used for murex builtin that outputs from a named pipe
	NamedPipeProcName = "<named-piped>"

	// CmdExec is the GoFunction name used for murex builtin that executes external processes without a TTY
	CmdExec = "exec"

	// CmdExec is the GoFunction name used for murex builtin that executes external processes with a TTY
	CmdPty = "pty"
)
