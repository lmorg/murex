### The `formap` builtin

Arrays are easy though. What about maps? Or dictionaries, objects, hashes... as other languages might refer to them. Key value pairs are always going to be harder to parse because even with file formats like YAML, arrays are line delimited where as splitting key value pairs requires extra grokking.

Murex also has an iterator for such constructs: `formap`:

```
» echo '{"a":1,"b":2,"c":3}' | :json: formap key value { echo "key=$key, value=$value" }
key=a, value=1
key=b, value=2
key=c, value=3
```

Here we are using `:json:` to cast / annotate the STDIN stream for `formap`.

### Lambdas

Lambdas were introduced in version 4.0 of Murex. Lambdas offer shortcuts around common matching problems with structured data. For example:

```
» $example = %{a:1, b:2, c:3}
» @example[{$.val == 2}]
{
    "b": 2
}
```

The way this particular lambda works is that for each element in `$example` the code block `{$.val == 2}` runs with the `$.` variable is assigned with a structure that looks like this:

```
# first iteration
{
    "key": "a",
    "val": 1
}
# second iteration
{
    "key": "b",
    "val": 2
}
# third iteration
{
    "key": "c",
    "val": 3
}
```

Thus `$.val` (eg **2**) would be compared to the number **2**. If the result is true, that element is output. If the result is false then that element is ignored.