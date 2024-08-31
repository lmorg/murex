# Download File (`getfile`)

> Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.

## Description

Fetches a resource from a URL - setting stdout data-type

## Usage

```
getfile url -> <stdout>
```

## Examples

```
getfile google.com 
```

## Detail

This simply fetches a resource (via HTTP GET request) from a URL and returns the
byte stream to stdout. It will set stdout's data-type based on MIME defined in
the `Content-Type` HTTP header.

It is recommended that you only use this command if you're pipelining the output
(eg writing to file or passing on to another function). If you just want to
render the output to the terminal then use `open` which has hooks for smart
terminal rendering.

### As A Method

Running `get`, `post` or `getfile` as a method will transmit the contents of
stdin as part of the body of the HTTP request. When run as a method the
`Content-Type` HTTP header will automatically be set to the default MIME for
the data type from stdin.

This is defined in `config`, pre-defined by sensible defaults from each murex
data type. For example:

```
» config get shell default-mimes -> [json]
application/json
```

You can override this at the global level via **shell default-mimes**, or at
the local level via **http headers**:

```
config set http headers %{
    api.example.com: {
        Content-Type: application/foobar
    }
}
```

### Configurable options

`getfile` has a number of behavioral options which can be configured via
Murex's standard `config` tool:

```
config -> [ http ]
```

To change a default, for example the user agent string:

```
config set http user-agent "bob"
getfile google.com
```

This enables sane, repeatable and readable defaults. Read the documents on
`config` for more details about it's usage and the rational behind the command.

## See Also

* [Get Request (`get`)](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Post Request (`post`)](../commands/post.md):
  HTTP POST request with a JSON-parsable return
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings

<hr/>

This document was generated from [builtins/core/httpclient/get_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/httpclient/get_doc.yaml).