## Creating Arrays

Arrays can be defined with `%[ ... ]`.

The syntax is a superset of JSON. So that any JSON array can be a Murex array
when prefixed with `%`. For example

```
%["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"]
```

...however while this is readable, this is not convenient when working in the
interactive command line. We want something that's a little more comfortable
for "write many, read once" type environments like a shell REPL.

So Murex makes the parser-defined punctuation optional:

```
%[Monday Tuesday Wednesday Thursday Friday]
```

That's better, however surely the computer already knows what days there are in
a week? I think we can improve this syntax further...

```
%[Monday..Friday]
```

The `..` describes a range. So we are saying "return everyday from Monday to
Friday, inclusive".

```
» %[Monday..Friday]
[
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday"
]
```

It's not just days of the week that can be completed like this. Most forms of
dates, number bases and other familiar sequences can be.

You can also have multiple parts of arrays expanded:

```
» %[[2024..2026]\\, [spring..winter]]
[
    "2024, spring",
    "2024, summer",
    "2024, autumn",
    "2024, winter",
    "2025, spring",
    "2025, summer",
    "2025, autumn",
    "2025, winter",
    "2026, spring",
    "2026, summer",
    "2026, autumn",
    "2026, winter"
]
```

Did you notice how the seasons are lower case this time? Murex respects the
text case of the ranges. For example:

```
» %[MON..WED]
[
    "MON",
    "TUE",
    "WED",
]

» %[mon..wed]
[
    "mon",
    "tue",
    "wed",
]
```

Returning to our multi-part array, perhaps you want **fall** instead of
**autumn**? (here you do need to include a comma)

```
» %[[2024..2026]\\, [spring,summer,fall,winter]]
[
    "2024, spring",
    "2024, summer",
    "2024, fall",
    "2024, winter",
    "2025, spring",
    "2025, summer",
    "2025, fall",
    "2025, winter",
    "2026, spring",
    "2026, summer",
    "2026, fall",
    "2026, winter"
]
```

### Multi-Dimensional Arrays

But what if they were supposed to be nested rather than flattened arrays? Well
that's not a problem either. Just make sure your first nested array is also
prefixed with `%`:

```
» %[%[2024..2026] [Spring..Winter]]
[
    [
        2024,
        2025,
        2026
    ],
    [
        "Spring",
        "Summer",
        "Autumn",
        "Winter"
    ]
]
```

### Streaming Arrays

Ok, that's great, but marshalling a data structure and passing it to functions
requires allocating that entire data structure to memory. Whereas shells best
excel when they're streaming data because it allows processes to perform
concurrent operations across massive data sets.

If that's a problem for you too, then Murex has you sorted: `a`.

The `a` builtin returns a line separated list. And thus can be operated on via
your traditional pipes and UNIX core utilities.

Lets say, for some reason, you wanted to run a job against every IPv4 address
available. Marshalling that into a data structure just to run commands
sequentially would be rather silly. So lets stream that array using `a`:

```
a [0..255].[0..255].[0..255].[0..255] | foreach $ip { ping -c 1 -t 1 $ip }
```

> If you do happen to run this and wonder how to cancel the underlying `a` and
> `foreach` routines, you can press `ctrl`+`\` -- which is a Murex shortcut to
> kill everything in that shell session.

This is obviously an absurd example because nobody in their right mind would
want to ping every valid IPv4 address. But it does demonstrate the advantages
of streaming lists rather than creating arrays.

## Accessing Array Values

There are two main ways to access values inside an array:

* square brackets for immutable copies: `$my_array[index]`

* dot notation, which allows being written to: `$my_array.index`

Why two? Because they support different features.

### Square Brackets (immutable)

With square brackets you can select more than just a single element. For
example, if you wanted the first element you can reference it the same way
you'd reference arrays in any other language:

```
$my_array[0]
```

> Murex arrays begin at `0`.

If you wanted to count from the end of the array, you can use negative values:

```
$my_array[-1]
```

> Watch out here because negative indexes count from `1` because -0 isn't a
> valid number.

#### Multiple Elements

However what if you wanted multiple elements from the array, like 2nd and 4th?
Then just specify multiple elements inside the square brackets:

```
$my_array[1 4]
```

#### Ranges

That's handy, but if I actually want a range of elements? Well then you can use
the range `..` operator like before:

```
$my_array[1..4]
```

And if you don't know the size of your array, you can ignore the index value
entirely. For example:

**Everything from and including the 2nd element:**

```
$my_array[2..]
```

**Everything up to and including the 4th element:**

```
$my_array[..4]
```

> Ranges are indexed from 1. Yes, I know that's stupid and confusing.

#### Elements From A Pipeline

The square brackets can also be used as a function too. Which means any kind of
array or list can be queried from stdin, you don't have to first convert it to
a variable.

```
» %[mon..fri] | [1 4]
[
    "tue",
    "fri"
]
```

This is especially helpful if say something writes JSON, YAML, or any other
structured document, and you only want specific values. For example you want
the last container in your cloud infrastructure, and you know your cloud CLI 
tool (`cloud-api` for our made up purposes here) returns JSON:

```
cloud-api list-containers | :json: [-1]
```

### Making Changes

That's all great, but what if I want to make a change to the host array?

Well this is where dot notation comes in...

## Dot Notation (mutable)

Dot notation is a lot more limited in what you can do because it's designed for
making careful edits of the underlying data structure. So it can only be used
with variables.

### Assignment

You can edit an element, for example renaming **Wednesday** to **Humpday**:

```
» $days = %[Monday..Friday]

» $days.2 = "Humpday"

» $days
[
    "Monday",
    "Tuesday",
    "Humpday",
    "Thursday",
    "Friday"
]
```

> Remember: arrays are zero based

### Printing

You can also use dot notation to return a value, just like you would with the
square braces solution above. But dot notation doesn't support any special
magic and still only works on variables:

```
$my_array.2
```

## Appending Arrays

There are several ways to append an array. You can create a copy of that array
and append via the pipeline:

```
$my_array | append foo bar
```

> You can also prepend using `prepend`)

...or...

```
$my_array ~> %[foo bar]
```

However if you want to update a variable in place, you can use the merge
operator, `<~`:

```
$my_array <~ %[foo bar]
```

## Fin

Now you're an expert in arrays. Hooray `\o/`