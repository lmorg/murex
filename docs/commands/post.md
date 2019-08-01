# _murex_ Language Guide

## Command Reference: `post`

> HTTP POST request with a JSON-parsable return

### Description

Fetches a page from a URL via HTTP/S POST request.

### Usage

    post url -> <stdout>
    
    <stdin> -> post url content-type -> <stdout>

### Examples

    » post google.com -> [ Status ] 
    {
        "Code": 405,
        "Message": "Method Not Allowed"
    }

### Detail

#### JSON return

`POST` returns a JSON object with the following fields:

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
    
The concept behind this is it provides and easier path for scripting eg pulling
specific fields via the index, `[`, function.

#### `post` as a method

Running `post` as a method will transmit the contents of STDIN as part of the
body of the HTTP POST request. When run as a method you have to include a second
parameter specifying the Content-Type MIME.

#### Configurable options

`post` has a number of behavioral options which can be configured via _murex_'s
standard `config` tool:

    config: -> [ http ]
    
To change a default, for example the user agent string:

    config: set http user-agent "bob"
    post: google.com
    
This enables sane, repeatable and readable defaults. Read the documents on
`config` for more details about it's usage and the rational behind the command.

### See Also

* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`get`](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [`getfile`](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.
* [config](../commands/config.md):
  
* [element](../commands/element.md):
  