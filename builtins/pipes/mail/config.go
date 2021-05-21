package mail

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	config.InitConf.Define("mail", "sender", config.Properties{
		Description: "Sender email address used when sending an email via the 'mail' pipe (ie FROM). This is used both with authenticated mail and MX sends",
		Default:     fmt.Sprintf("%s@%s", username(), hostname()),
		DataType:    types.String,
		Global:      true,
	})

	config.InitConf.Define("mail", "allow-plain-text", config.Properties{
		Description: "Allow plain text connections when sending direct via MX records (not recommended)",
		Default:     false,
		DataType:    types.Boolean,
		Global:      true,
	})

	config.InitConf.Define("mail", "ignore-verify-certs", config.Properties{
		Description: "Don't verify the server name when making TLS connections when sending direct via MX records (not recommended)",
		Default:     false,
		DataType:    types.Boolean,
		Global:      true,
	})

	config.InitConf.Define("mail", "ports", config.Properties{
		Description: "Port numbers to try, in order of precedence when sending direct via MX records",
		Default:     ports,
		DataType:    types.Json,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  getPorts,
			Write: setPorts,
		},
	})

	config.InitConf.Define("mail", "smtp-auth-host", config.Properties{
		Description: "SMTP hostname (eg 'smtp.google.com') to send authenticated mail from (if unset, murex will attempt to send directly from via MX records and not use auth)",
		Default:     "",
		DataType:    types.String,
		Global:      true,
	})

	config.InitConf.Define("mail", "smtp-auth-port", config.Properties{
		Description: "SMTP port number (eg 587) to send authenticated mail (if unset, murex will attempt to send directly from via MX records and not use auth)",
		Default:     587,
		DataType:    types.Integer,
		Global:      true,
	})

	config.InitConf.Define("mail", "smtp-auth-user", config.Properties{
		Description: "User name when using SMTP user auth (requires 'smtp-host' set too)",
		Default:     "",
		DataType:    types.String,
		Global:      true,
	})

	config.InitConf.Define("mail", "smtp-auth-pass", config.Properties{
		Description: "Password when using SMTP user auth (requires 'smtp-host' set too)",
		Default:     "",
		DataType:    types.String,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  getPass,
			Write: setPass,
		},
	})

	config.InitConf.Define("mail", "smtp-auth-enabled", config.Properties{
		Description: "Enable or disable SMTP auth",
		Default:     false,
		DataType:    types.Boolean,
		Global:      true,
	})
}

var ports = []int{
	587,
	2525,
	465,
	25,
}

func getPorts() (interface{}, error) {
	return ports, nil
}

func setPorts(v interface{}) error {
	switch v.(type) {
	case string:
		return json.Unmarshal([]byte(v.(string)), &ports)

	default:
		return fmt.Errorf("Invalid data-type. Expecting a %s encoded string", types.Json)
	}
}

func allowPlainText() bool {
	v, err := lang.ShellProcess.Config.Get("mail", "allow-plain-text", types.Boolean)
	if err != nil {
		return false
	}

	return v.(bool)
}

func allowInsecure() bool {
	v, err := lang.ShellProcess.Config.Get("mail", "ignore-verify-certs", types.Boolean)
	if err != nil {
		return false
	}

	return v.(bool)
}

func senderAddr() string {
	v, err := lang.ShellProcess.Config.Get("mail", "sender", types.String)
	if err != nil {
		return fmt.Sprintf("%s@%s", username(), hostname())
	}

	s := strings.TrimSpace(v.(string))
	if len(s) == 0 {
		return fmt.Sprintf("%s@%s", username(), hostname())
	}

	return s
}

func smtpAuthHost() string {
	v, err := lang.ShellProcess.Config.Get("mail", "smtp-auth-host", types.String)
	if err != nil {
		return ""
	}

	return v.(string)
}

func smtpAuthPort() int {
	v, err := lang.ShellProcess.Config.Get("mail", "smtp-auth-port", types.Integer)
	if err != nil {
		return 587
	}

	return v.(int)
}

var (
	smtpPassword string
)

func getPass() (interface{}, error) {
	if smtpPassword == "" {
		return "unset", nil
	}
	return "redacted", nil
}

func setPass(v interface{}) error {
	switch v.(type) {
	case string:
		smtpPassword = v.(string)
		return nil

	default:
		return fmt.Errorf("Invalid data-type. Expecting a %s", types.String)
	}
}

func getSmtpAuth() (smtpAuth smtp.Auth, enabled bool) {
	var (
		user string
		pass string
		host string
	)

	v, err := lang.ShellProcess.Config.Get("mail", "smtp-auth-user", types.String)
	if err != nil {
		user = ""
	} else {
		user = v.(string)
	}

	/*v, err = lang.ShellProcess.Config.Get("mail", "smtp-auth-pass", types.String)
	if err != nil {
		pass = ""
	} else {
		pass = v.(string)
	}*/
	pass = smtpPassword

	v, err = lang.ShellProcess.Config.Get("mail", "smtp-auth-enabled", types.Boolean)
	if err != nil {
		enabled = false
	} else {
		enabled = v.(bool)
	}

	host = smtpAuthHost()
	if host == "" {
		enabled = false
	}

	if enabled {
		smtpAuth = smtp.PlainAuth("", user, pass, host)
	}

	return
}
