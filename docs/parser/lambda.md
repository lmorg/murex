# `[{ Lambda }]`

> Iterate through structured data

## Description

Lambdas, in Murex, are a concise way of performing various actions against
structured data. They're a convenience tool to range over arrays and objects,
similar to `foreach` and `formap`.

### Etiquette

The intention of lambdas isn't to be used liberally within shell scripts but
rather than allow for more efficient one liners when performing quick, often
one-off, tasks in the interactive terminal. The terse syntax for lambdas
combined with it's adaptive functionality allow for flexibility when using,
albeit at the potential cost of readability.

For shell scripts where error handling, readability, and maintainability are a
concern, more conventional iteration blocks like `foreach` and `formap` are
recommended.

The reason lambda variables (known as **meta values** are single characters
also falls in line with this vision: they're terse to make one-liners
convenient and to discourage shell script usage.

### Technical

Code running inside a lambda inherit a special variable, named **meta values**,
which hold the state for each iteration (see section below).

Lambdas will adapt its return value depending on the nature of the code it's
executing.

#### Filter

If your code returns a boolean data type, eg `[{$.v =~ "foobar"}]`, then the
lambda will filter that list or map, only returning values that match `true`,
discarding the others.

#### Update

If the **meta value** `v` (value) is updated then the lambda output reflects
that change.

The **meta value** `k` (key) works similarly for maps / objects too. However
updating it in arrays and lists currently does nothing.

This cannot be used in combination with **output**.

#### Output

If stdout isn't empty, then stdout is outputted rather than the object being
filtered and/or updated.

This usage most closely resembles `foreach` and `formap` except that the data
type of stdout is not preserved.



## Examples

### Filter

Filtering a map:

```
» %{hello: world, foo: bar} -> [{$.v == "world"}]
{
    "hello": "world"
}
```

Filtering an array:

In this example we return the days of the week excluding todays day (in this
example, today is Friday)

```
» %[Monday..Sunday] -> [{$.v != datetime(--in {now} --out {py}%A)}]
[
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Saturday",
    "Sunday"
]
```

### Update

Updating a map:

```
» %{hello: world, foo: bar} -> [{$.v == "world" -> if {$.v = "Earth"}}]
{
    "foo": "bar",
    "hello": "Earth"
}
```

Updating an array:

```
» %[monday..friday] -> [{$.v | tr [:lower:] [:upper:] | set $.v}]
[
    "MONDAY",
    "TUESDAY",
    "WEDNESDAY",
    "THURSDAY",
    "FRIDAY"
]
```

### Output

Output from a map:

```
» %{hello: world, foo: bar} -> [{$.v == "world" && out "Key '$.k' contains '$.v'"}]
Key 'hello' contains 'world'
```

Output from an array:

```
» %[Monday..Sunday] -> [{ $.v =~ "^S" && out "$.v is the weekend"}]
Saturday is the weekend
Sunday is the weekend
```

### Foreach

Here we are using a lambda just as a terser way of writing a standard `foreach`
loop:

```
» %[Monday..Sunday] -> [{$.v =~ "^S" && $count+=1}]; echo "$count days being with an 'S'"
Error in `expr` (0,22): [json marshaller] no data returned
2 days being with an 'S'
```

However this is a contrived example. The more idiomatic way to write the above
would be (and notice it doesn't produce any empty array error too):

```
» %[Monday..Sunday] -> foreach day { $day =~ "^S" && $count+=1}; echo "$count days being with an 'S'"
2 days being with an 'S'
```

...or even just using `regexp`, since the check is just a simple regexp match:

```
» %[Monday..Sunday] -> regexp m/^S/ -> count
2
```

## Detail

### Meta values

Meta values are a JSON object stored as the variable `$.`. The meta variable
will get overwritten by any other block which invokes meta values. So if you
wish to persist meta values across blocks you will need to reassign `$.`, eg

```
%[1..3] -> foreach {
    meta_parent = $.
    %[7..9] -> foreach {
        out "$(meta_parent.i): $.i"
    }
}
```

The following meta values are defined:

* `i`: iteration number (counts from one)
* `k`: key name (for arrays / lists this will count from zero)
* `v`: item value of map / object or array / list

## See Also

* [Alter Data Structure (`alter`)](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [Date And Time Conversion (`datetime`)](../commands/datetime.md):
  A date and/or time conversion tool (like `printf` but for date and time values)
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [Loop While (`while`)](../commands/while.md):
  Loop until condition false
* [Regex Operations (`regexp`)](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [`%[]` Array Builder](../parser/create-array.md):
  Quickly generate arrays
* [`%{}` Object Builder](../parser/create-object.md):
  Quickly generate objects (dictionaries / maps)

<hr/>

This document was generated from [gen/parser/lambda_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/lambda_doc.yaml).