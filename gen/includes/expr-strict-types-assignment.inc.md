### Strict Types

Unlike with the standard arithmetic operators (`+`, `-`, `*`, `/`), silent data
casting isn't supported with arithmetic assignments like `+=`, `-=`, `*=` and
`/=`. Not even when `strict-types` is disabled ({{link "read more" "strict-types"}})

You can work around this by using the slightly longer syntax: **variable =
value op value**, for example:

```
» $i = "3"
» $i = $i + "2"
» $i
5
```