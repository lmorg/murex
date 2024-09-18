### Creating JSON files

First we need to create two JSON files.

Here the syntax is pretty similar however with the Nushell code we are using
the command `save` whereas the Murex equivalent uses the `|>` operator, which
is like the Bash `>` redirection pipe.

Nushell:

```nu
[0,1,2,3] | range 0..3 | save a.json
[4,5,6,7] | range 0..3 | save b.json
```

Murex:

```mx
%[0..3] |> a.json
%[4..7] |> b.json
```

This will create two files that look like

```
# a.json
[
  0,
  1,
  2,
  3
]
# b.json
[
  4,
  5,
  6,
  7
]
```

> Murex's JSON files are minified by default. You can make them human readable
> with the `pretty` command, eg
> 
> ```
> %[0..3] | pretty |> a.json
> ```

### Embedding arrays in an array

To then embed two the files together, with Nushell you reuse the `echo` command
whereas in Murex you would use the array builder:

Nushell:

```nu
echo (open a.json) (open b.json) | save c.json
```

Murex:

```mx
%[ ${open a.json} ${open b.json} ]> c.json
```

### Merging and flattening two arrays

With Murex merging (or "flattening") the two arrays can be done with a single
operator `~>`.

Nushell:

```
echo (open a.json) (open b.json) | flatten | save c.json
```

Murex:

```
open(a.json) ~> open(b.json) |> c.json
```

### Nushell references

The Nushell script used here can be found at https://github.com/nushell/nu_scripts/blob/main/sourced/cool-oneliners/file_cat.nu

