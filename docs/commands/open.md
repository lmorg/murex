# `open` - Command Reference

> Open a file with a preferred handler

## Description

`open` is a smart tool for reading files:

1. It will read a file from disk or a HTTP(S) endpoints
2. Detect the file type via file extension or HTTP header `Content-Type`
3. It intelligently writes to STDOUT
  - If STDOUT is a TTY it will perform any transformations to render to the
    terminal (eg using inlining images)
  - If STDOUT is a pipe then it will write a byte stream with the relevant
    data-type

## Usage

    open filename[.gz]|uri -> <stdout>

## Examples

    » open https://api.github.com/repos/lmorg/murex/issues -> foreach issue { out: "$issue[number]: $issue[title]" }

## Detail

### File Extensions

Supported file extensions are listed in `config` under the app and key names of
**shell**, **extensions**.

Unsupported file extensions are defaulted to generic, `*`.

Files with a `.gz` extension are assumed to be gzipped and thus are are
automatically expanded.

### MIME Types

The `Content-Type` HTTP header is compared against a list of MIME types, which
are stored in `config` under the app and key names of **shell**, **mime-types**.

There is a little bit of additional logic to determine the _murex_ data-type to
use should the MIME type not appear in `config`, as seen in the following code:

```go
package lang

import (
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

var rxMimePrefix = regexp.MustCompile(`^([-0-9a-zA-Z]+)/.*$`)

// MimeToMurex gets the murex data type for a corresponding MIME
func MimeToMurex(mimeType string) string {
	mime := strings.Split(mimeType, ";")[0]
	mime = strings.TrimSpace(mime)
	mime = strings.ToLower(mime)

	// Find a direct match. This is only used to pick up edge cases, eg text files used as images.
	dt := mimes[mime]
	if dt != "" {
		return dt
	}

	// No direct match found. Fall back to prefix.
	prefix := rxMimePrefix.FindStringSubmatch(mime)
	if len(prefix) != 2 {
		return types.Generic
	}

	switch prefix[1] {
	case "text", "i-world", "message":
		return types.String

	case "audio", "music", "video", "image", "model":
		return types.Binary

	case "application":
		if strings.HasSuffix(mime, "+json") {
			return types.Json
		}
		return types.Generic

	default:
		// Mime type not recognized so lets just make it a generic.
		return types.Generic
	}

}
```

### HTTP User Agent

`open`'s user agent is the same as `get` and `post` and is configurable via
`config` under they app **http**

    » config -> [http]
    {
        "cookies": {
            "Data-Type": "json",
            "Default": {
                "example.com": {
                    "name": "value"
                },
                "www.example.com": {
                    "name": "value"
                }
            },
            "Description": "Defined cookies to send, ordered by domain.",
            "Dynamic": false,
            "Global": false,
            "Value": {
                "example.com": {
                    "name": "value"
                },
                "www.example.com": {
                    "name": "value"
                }
            }
        },
        "default-https": {
            "Data-Type": "bool",
            "Default": false,
            "Description": "If true then when no protocol is specified (`http://` nor `https://`) then default to `https://`.",
            "Dynamic": false,
            "Global": false,
            "Value": false
        },
        "headers": {
            "Data-Type": "json",
            "Default": {
                "example.com": {
                    "name": "value"
                },
                "www.example.com": {
                    "name": "value"
                }
            },
            "Description": "Defined HTTP request headers to send, ordered by domain.",
            "Dynamic": false,
            "Global": false,
            "Value": {
                "example.com": {
                    "name": "value"
                },
                "www.example.com": {
                    "name": "value"
                }
            }
        },
        "insecure": {
            "Data-Type": "bool",
            "Default": false,
            "Description": "Ignore certificate errors.",
            "Dynamic": false,
            "Global": false,
            "Value": false
        },
        "redirect": {
            "Data-Type": "bool",
            "Default": true,
            "Description": "Automatically follow redirects.",
            "Dynamic": false,
            "Global": false,
            "Value": true
        },
        "timeout": {
            "Data-Type": "int",
            "Default": 10,
            "Description": "Timeout in seconds for `get` and `getfile`.",
            "Dynamic": false,
            "Global": false,
            "Value": 10
        },
        "user-agent": {
            "Data-Type": "str",
            "Default": "murex/1.7.0000 BETA",
            "Description": "User agent string for `get` and `getfile`.",
            "Dynamic": false,
            "Global": false,
            "Value": "murex/1.7.0000 BETA"
        }
    }

## See Also

* [`*` (generic) ](../types/generic.md):
  generic (primitive)
* [`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [`exec`](../commands/exec.md):
  Runs an executable
* [`fexec` ](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [`foreach`](../commands/foreach.md):
  Iterate through an array
* [`get`](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [`getfile`](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.
* [`openagent`](../commands/openagent.md):
  Creates a handler function for `open
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`post`](../commands/post.md):
  HTTP POST request with a JSON-parsable return