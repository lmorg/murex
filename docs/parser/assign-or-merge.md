# `<~` Assign Or Merge

> Merges the right hand value to a variable on the left hand side (expression)

## Description

The **Assign Or Merge** operator merges your data from a the right hand side
into the variable on the left hand side.

If the variable doesn't exist, then it is created.

This operator is only available in expressions.



## Examples

### Appending to an array

Lets say you have a directory hierarchy that looks like:

```
» tree
.
├── a
│   ├── )rPsD8Dt5EtaC4*Yyn0q
│   ├── B[E3P2@gyzl2oSfvFs5(
│   ├── WNYBb>B{Y:9oBNq~eVn{
│   ├── W~e5bLBkGkv 2sr<XTj:
│   └── lgCVRC.PkUkh(!epI(ls
├── b
│   ├── ]^[;og5$x'%Zp* TY(NR
│   ├── kKcuV<9@pBrFr@"O\j?%
│   ├── wX'\>V`4P=K}FaxE^Hra
│   ├── yXjB#Cu'{%iLtsDCkKU%
│   └── |oK7e25Dz7z&ys.?2(]E
├── c
│   ├── )Kb!TOQ]\9J6 &<Y\2qj
│   ├── -X-Dm,m[JU0FZ#b0+fe+
│   ├── Lw2"`S<ag{EnJ=YI8A\W
│   ├── W4RUF_D.z,%M|OFsLB_A
│   └── z@meR3m7h(~V4m7(V{N
├── d
│   ├── %"6Tn]&w@Uas*Gi5$?Q0
│   ├── F}Ly:]zGTk}4]V+L=Wk+
│   ├── z%;lf^2n0r'p0Fy?f[$j
│   ├── {Iz}*#HCR_@H.KyA3=xy
│   └── ~2hUs'_NftfpH`?>Bqpt
├── e
│   ├── :#'G'Rs~^~A)g,k29Er1
│   ├── =N-KR9!lh"H FjCp@sP%
│   └── ?,E XTt%TGD4vrvR@qXw
└── f
    ├── !B#v!iYSBmi<i6[mdlL'
    ├── _@$*?WgS0KozEnmHV*gW
    ├── eT8?OgIK@4zSHTz0$m<O
    └── |c[c-8S.;X$&UzI@jp!X

7 directories, 27 files
```

...and in this example you want to list the files in only directories that are
vowels, you can use the **Assign Or Merge** operator to append to the list on
each iteration for `foreach`:

```
» %[a,e] -> foreach d { files <~ g($d/*) }
» $files
[
    "a/)rPsD8Dt5EtaC4*Yyn0q",
    "a/B[E3P2@gyzl2oSfvFs5(",
    "a/WNYBb\u003eB{Y:9oBNq~eVn{",
    "a/W~e5bLBkGkv 2sr\u003cXTj:",
    "a/lgCVRC.PkUkh(!epI(ls",
    "e/:#'G'Rs~^~A)g,k29Er1",
    "e/=N-KR9!lh\"H FjCp@sP%",
    "e/?,E XTt%TGD4vrvR@qXw"
]
```

> Please note that you can also do this with [standard globbing](/docs/commands/g.md), for example:
> ```
> » files = g([ae]/*)
> ```

## Detail

**Assign Or Merge** uses the same underlying algorithm as the [`alter -m`](/docs/commands/alter.md)
builtin.

## See Also

* [Alter Data Structure (`alter`)](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [Globbing (`g`)](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [Operators And Tokens](../user-guide/operators-and-tokens.md):
  A table of all supported operators and tokens
* [`%[]` Array Builder](../parser/create-array.md):
  Quickly generate arrays

<hr/>

This document was generated from [gen/expr/assign-merge-op_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/expr/assign-merge-op_doc.yaml).