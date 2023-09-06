//go:build opt_mail
// +build opt_mail

package optional

// This is an optional builtin because it duplicates sendmail functionality,
// but it's a handy pipe to have if you wish to send emails from the command
// line using idiomatic murex syntax and controls

import _ "github.com/lmorg/murex/builtins/pipes/mail" // piping data via esmail
