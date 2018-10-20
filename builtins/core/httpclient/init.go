package httpclient

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["get"] = cmdGet
	proc.GoFunctions["post"] = cmdPost
	proc.GoFunctions["getfile"] = cmdGetFile

	proc.InitConf.Define("http", "user-agent", config.Properties{
		Description: "User agent string for `get` and `getfile`.",
		Default:     config.AppName + "/" + config.Version,
		DataType:    types.String,
	})

	proc.InitConf.Define("http", "timeout", config.Properties{
		Description: "Timeout in seconds for `get` and `getfile`.",
		Default:     10,
		DataType:    types.Integer,
	})

	proc.InitConf.Define("http", "insecure", config.Properties{
		Description: "Ignore certificate errors.",
		Default:     false,
		DataType:    types.Boolean,
	})

	proc.InitConf.Define("http", "redirect", config.Properties{
		Description: "Automatically follow redirects.",
		Default:     true,
		DataType:    types.Boolean,
	})

	proc.InitConf.Define("http", "default-https", config.Properties{
		Description: "If true then when no protocol is specified (`http://` not `https://`) then default to `https://`.",
		Default:     false,
		DataType:    types.Boolean,
	})

	proc.InitConf.Define("http", "cookies", config.Properties{
		Description: "Defined cookies to send, ordered by domain.",
		Default: metaDomains{
			"example.com":     metaValues{"name": "value"},
			"www.example.com": metaValues{"name": "value"},
		},
		DataType: types.Json,
	})

	proc.InitConf.Define("http", "headers", config.Properties{
		Description: "Defined HTTP request headers to send, ordered by domain.",
		Default: metaDomains{
			"example.com":     metaValues{"name": "value"},
			"www.example.com": metaValues{"name": "value"},
		},
		DataType: types.Json,
	})
}
