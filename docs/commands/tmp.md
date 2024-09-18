# Create Temporary File (`tmp`)

> Create a temporary file and write to it

## Description

`tmp` creates a temporary file, writes the contents of stdin to it then returns
its filename to stdout.

You can optionally specify a file extension, for example if the temporary file
needs to be read by `open` or an editor which uses extensions to define syntax
highlighting.

## Usage

```
<stdin> -> tmp [ file-extension ] -> <stdout>
```

## Examples

```
» out "Hello, world!" -> set tmp

» out $tmp
/var/folders/3t/267q_b0j27d29bnf6pf7m7vm0000gn/T/murex838290600/8ec6936c1ac1c347bf85675eab4a0877-13893

» open $tmp
Hello, world!
```

## Detail

The temporary file name is a base64 encoded md5 hash of the time plus Murex
function ID with Murex process ID appended:

```go
package io

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

func init() {
	lang.DefineMethod("tmp", cmdTempFile, types.Any, types.String)
}

func cmdTempFile(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)

	ext, _ := p.Parameters.String(0)
	if ext != "" {
		ext = "." + ext
	}

	fileId := time.Now().String() + ":" + strconv.Itoa(int(p.Id))

	h := md5.New()
	_, err := h.Write([]byte(fileId))
	if err != nil {
		return err
	}

	name := consts.TempDir + hex.EncodeToString(h.Sum(nil)) + "-" + strconv.Itoa(os.Getpid()) + ext

	file, err := os.Create(name)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, p.Stdin)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write([]byte(name))
	return err
}
```

This should should provide enough distance to run `tmp` in parallel....should
you ever want to.

`tmp` files are also located inside a unique per-process Murex temp directory
which itself is located in the appropriate temp directory for the host OS (eg
`$TMPDIR` on macOS).

## Synonyms

* `tmp`


## See Also

* [Create Named Pipe (`pipe`)](../commands/pipe.md):
  Manage Murex named pipes
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Truncate File (`>`)](../parser/file-truncate.md):
  Writes stdin to disk - overwriting contents if file already exists
* [`>>` Append File](../parser/file-append.md):
  Writes stdin to disk - appending contents if file already exists

<hr/>

This document was generated from [builtins/core/io/tmp_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/tmp_doc.yaml).