package httpclient

import (
	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["get"] = cmdGet
	lang.GoFunctions["post"] = cmdPost
	lang.GoFunctions["getfile"] = cmdGetFile

	config.InitConf.Define("http", "user-agent", config.Properties{
		Description: "User agent string for `get` and `getfile`.",
		Default:     app.Name + "/" + app.Version,
		DataType:    types.String,
	})

	config.InitConf.Define("http", "timeout", config.Properties{
		Description: "Timeout in seconds for `get` and `getfile`.",
		Default:     10,
		DataType:    types.Integer,
	})

	config.InitConf.Define("http", "insecure", config.Properties{
		Description: "Ignore certificate errors.",
		Default:     false,
		DataType:    types.Boolean,
	})

	config.InitConf.Define("http", "redirect", config.Properties{
		Description: "Automatically follow redirects.",
		Default:     true,
		DataType:    types.Boolean,
	})

	config.InitConf.Define("http", "default-https", config.Properties{
		Description: "If true then when no protocol is specified (`http://` nor `https://`) then default to `https://`.",
		Default:     false,
		DataType:    types.Boolean,
	})

	config.InitConf.Define("http", "cookies", config.Properties{
		Description: "Defined cookies to send, ordered by domain.",
		Default: metaDomains{
			"example.com":     metaValues{"name": "value"},
			"www.example.com": metaValues{"name": "value"},
		},
		DataType: types.Json,
	})

	config.InitConf.Define("http", "headers", config.Properties{
		Description: "Defined HTTP request headers to send, ordered by domain.",
		Default: metaDomains{
			"example.com":     metaValues{"name": "value"},
			"www.example.com": metaValues{"name": "value"},
		},
		DataType: types.Json,
	})
}
