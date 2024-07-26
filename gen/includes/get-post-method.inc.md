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