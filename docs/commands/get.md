# `get`

> Makes a standard HTTP request and returns the result as a JSON object

## Description

Fetches a page from a URL via HTTP/S GET request

## Usage

```
get url -> <stdout>

<stdin> -> get url -> <stdout>
```

## Examples

```
» get google.com -> [ Status ]
{
    "Code": 200,
    "Message": "OK"
}
```

## Detail

### JSON return

`get` returns a JSON object with the following fields:

```
{
    "Status": {
        "Code": integer,
        "Message": string,
    },
    "Headers": {
        string [
            string...
        ]
    },
    "Body": string
}
```

The concept behind this is it provides and easier path for scripting eg pulling
specific fields via the index, `[`, function.

### `get` as a method

Running `get` as a method will transmit the contents of STDIN as part of the
body of the HTTP GET request. When run as a method you have to include a second
parameter specifying the Content-Type MIME.

### Configurable options

`get` has a number of behavioral options which can be configured via Murex's
standard `config` tool:

```
config -> [ http ]
```

To change a default, for example the user agent string:

```
config set http user-agent "bob"
get: google.com
```

This enables sane, repeatable and readable defaults. Read the documents on
`config` for more details about it's usage and the rational behind the command.

## See Also

* [`[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [`[[ Element ]]`](../parser/element.md):
  Outputs an element from a nested structure
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`getfile`](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.
* [`post`](../commands/post.md):
  HTTP POST request with a JSON-parsable return

<hr/>

This document was generated from [builtins/core/httpclient/get_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/httpclient/get_doc.yaml).