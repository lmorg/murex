### Type Safety

Because shells are historically untyped, you cannot always guarantee that a
numeric-looking value isn't a string. To solve this problem, by default Murex
assumes anything that looks like a number is a number when performing addition.

```
» str = "2"
» int = 3
» $str + $int
1
```

For occasions when type safety is more important than the convenience of silent
data casting, you can disable the above behaviour via `config` ({{link "read more" "strict-types"}}):

```
» config set proc strict-types true
» $str + $int
Error in `expr` (0,1): cannot Add with string types
                     > Expression: $str + $int
                     >           : ^
                     > Character : 1
                     > Symbol    : Scalar
                     > Value     : '$str'
```