# Strict Types In Expressions

> Expressions can auto-convert types or strictly honour data types

## Automatic Type Conversion

By default, any arithmetic expression on values that look like numbers will
proceed to treat those values as numbers even if their underlying data type is
not numeric. This can be illustrated with the following example:

```
» 1 + "2"
3
```

This is useful because shell processing is generally operating on bytes, which
often are treated as strings. Thus the OS doesn't have any concept of a numeric
verses non-numeric parameter. As far as Linux or Windows knows, all numbers are
strings too. So having Murex auto-convert values reduces your day-to-day
cognitive overhead when working in the shell.

## Strictly Honouring Type Annotations

Sometimes you might be writing Murex-native code and that code needs to care
about type safety. For example, you might need the leading zeros retained in
string when making comparisons. In situations like this, you can enable strict
types via `config`:

```
config set proc strict-types true
```

Using the same example from the previous section but with **strict-types** set
to `true`:

```
» 1 + "2"
Error in `expr` (0,1): cannot Add with string types
                     > Expression: 1 + "2"
                     >           :     ^
                     > Character : 5
                     > Symbol    : QuoteDouble
                     > Value     : '2'
```

## Scope

**strict-types** will be [scoped](/docs/user-guide/scoping.md) to that function and thus
changing this option will not cause accidental side-effects in other functions.

## See Also

* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Language Tour](../Murex/tour.md):
  Getting started with Murex
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex

<hr/>

This document was generated from [gen/user-guide/strict-types_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/strict-types_doc.yaml).