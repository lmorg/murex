# Open File (`open`)

> Open a file with a preferred handler

## Description

`open` is a smart tool for reading files:

1. It will read a file from disk or a HTTP(S) endpoints
2. Detect the file type via file extension or HTTP header `Content-Type`
3. It intelligently writes to stdout
  - If stdout is a TTY it will perform any transformations to render to the
    terminal (eg using inlining images)
  - If stdout is a pipe then it will write a byte stream with the relevant
    data-type
4. If there are no open handlers then it will fallback to the systems default.
   eg `open` (on macOS, Linux), `open-xdg` (X11), etc.

## Usage

```
open filename[.gz]|uri -> <stdout>
```

## Examples

```
» open https://api.github.com/repos/lmorg/murex/issues -> foreach issue { out "$issue[number]: $issue[title]" }
```

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

There is a little bit of additional logic to determine the Murex data-type to
use should the MIME type not appear in `config`, as seen in the following code:

```go
package lang

import (
	"path"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

var rxMimePrefix = regexp.MustCompile(`^([-0-9a-zA-Z]+)/.*$`)

// MimeToMurex gets the murex data type for a corresponding MIME
func MimeToMurex(mimeType string) string {
	return mimeToMurex(normalizeMime(mimeType))
}
func mimeToMurex(mime string) string {
	// Check suffix
	for suffix := range mimeSuffixes {
		if strings.HasSuffix(mime, suffix) {
			return mimeSuffixes[suffix]
		}
	}

	// Find a direct match
	dt := mimes[mime]
	if dt != "" {
		return dt
	}

	// No matches found. Fall back to prefix
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
		return types.Generic

	default:
		// Mime type not recognized so lets just make it a generic
		return types.Generic
	}
}

// RequestToMurex attempts to pick a data type for the content of
// a downloaded file using the value of the Content-Type header
// or, if not, the filename. If neither can be used to infer a
// usable data type, types.Generic is returned.
func RequestMetadataToMurex(contentType string, filename string) string {
	contentType = normalizeMime(contentType)

	if !(contentType == "" || contentType == "text/plain") {
		foundType := mimeToMurex(contentType)

		// MimeToMurex will return Generic when it can't find
		// any other matches, but in that case we still want
		// to check the filename ext if possible
		if foundType != types.Generic {
			return foundType
		}
	}

	if filename != "" {
		ext := strings.ToLower(strings.TrimPrefix(path.Ext(filename), "."))
		knownType, foundType := GetFileExts()[ext]

		if foundType {
			return knownType
		}
	}

	return types.Generic
}

// NormalizeMime prepares a mime type for processing by removing charset
// information, trimming leading/trailing spaces, and making it lower case
func normalizeMime(rawMimeType string) string {
	mime := strings.Split(rawMimeType, ";")[0]
	mime = strings.TrimSpace(mime)
	mime = strings.ToLower(mime)

	return mime
}

// MurexToMime returns the default MIME for a given Murex data type.
// The intended use case for this is for GET and POST requests where the body
// is STDIN.
func MurexToMime(dataType string) string {
	return defaultMimes[dataType]
}
```

### HTTP User Agent

`open`'s user agent is the same as `get` and `post` and is configurable via
`config` under they app **http**

```
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
```

### Open Flags

If the `open` builtin falls back to using the systems default (like `open-xdg`)
then the only thing that gets passed is the path being opened. If the path is
stdin then a temporary file will be created. If you want to pass command line
flags to `open-xdg` (for example), then you need to call that command directly.
In the case of macOS and some Linux systems, that might look like:

```
exec open --flags filename
```

## Synonyms

* `open`


## See Also

* [Define Handlers For "`open`" (`openagent`)](../commands/openagent.md):
  Creates a handler function for `open`
* [Download File (`getfile`)](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Execute Shell Function or Builtin (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [Get Request (`get`)](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Post Request (`post`)](../commands/post.md):
  HTTP POST request with a JSON-parsable return
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [`*` (generic)](../types/generic.md):
  generic (primitive)

<hr/>

This document was generated from [builtins/core/open/open_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/open/open_doc.yaml).