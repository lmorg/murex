# Calendar Date Ranges

> Create arrays of dates

## Description

Unlike bash, Murex also supports date ranges:

```  
» a [25-dec-2020..05-jan-2021]
» a [..25-dec-2020]
» a [25-dec-2020..]
```

Please refer to [a (mkarray)](../commands/a.md) for more detailed usage of mkarray.

## Usage

```
a: [start..end] -> <stdout>
a: [start..end,start..end] -> <stdout>
a: [start..end][start..end] -> <stdout>
```

All usages also work with `ja` and `ta` as well, eg:

```
ja: [start..end] -> <stdout>
ta: data-type [start..end] -> <stdout>
```

You can also inline arrays with the `%[]` syntax, eg:

```
%[start..end]
```

## Examples

```
» a [25-Dec-2020..01-Jan-2021]
25-Dec-2020
26-Dec-2020
27-Dec-2020
28-Dec-2020
29-Dec-2020
30-Dec-2020
31-Dec-2020
01-Jan-2021
```

```
» a [31-Dec..25-Dec]
31-Dec
30-Dec
29-Dec
28-Dec
27-Dec
26-Dec
25-Dec
```

## Detail

### Current Date

If the start value is missing (eg `[..01-Jan-2020]`) then mkarray (`a` et al)
will start the range from the current date and count up or down to the end.

If the end value is missing (eg `[01-Jan-2020..]`) then mkarray will start at
the start value, as usual, and count up or down to the current date.

For example, if today was 25th December 2020:

```
» a [23-December-2020..]
23-December-2020
24-December-2020
25-December-2020
```

```
» a [..23-December-2020]
25-December-2020
24-December-2020
23-December-2020
```

This can lead so some fun like countdowns:

```
» out "${a: [..01-January-2021] -> len -> =-1} days until the new year!"
7 days until the new year!
```

### Case Sensitivity

Date ranges are case aware. If the ranges are uppercase then the return will be
uppercase. If the ranges are title case (capital first letter) then the return
will be in title case.

#### lower case

```
» a [01-jan..03-jan]
01-jan
02-jan
03-jan
```

#### Title Case

```
» a [01-Jan..03-Jan]
01-Jan
02-Jan
03-Jan
```

#### UPPER CASE

```
» a [01-JAN..03-JAN]
01-JAN
02-JAN
03-JAN
```

### Supported Date Formatting

Below is the source for the supported formatting options for date ranges:

```go
package mkarray

var dateFormat = []string{
	// dd mm yy

	"02-Jan-06",
	"02-January-06",
	"02-Jan-2006",
	"02-January-2006",

	"02 Jan 06",
	"02 January 06",
	"02 Jan 2006",
	"02 January 2006",

	"02/Jan/06",
	"02/January/06",
	"02/Jan/2006",
	"02/January/2006",

	// mm dd yy

	"Jan-02-06",
	"January-02-06",
	"Jan-02-2006",
	"January-02-2006",

	"Jan 02 06",
	"January 02 06",
	"Jan 02 2006",
	"January 02 2006",

	"Jan/02/06",
	"January/02/06",
	"Jan/02/2006",
	"January/02/2006",

	// dd mm

	"02-Jan",
	"02-January",

	"02 Jan",
	"02 January",

	"02/Jan",
	"02/January",
}
```

If you do need any other formatting options not supported there, you can use
`datetime` to convert the output of `a`. eg:

```
» a [01-Jan-2020..03-Jan-2020] -> foreach { -> datetime --in "{go}02-Jan-2006" --out "{py}%A, %d %B"; echo }
Wednesday, 01 January
Thursday, 02 January
Friday, 03 January
```

## See Also

* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create New Array (`ta`)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [Date And Time Conversion (`datetime`)](../commands/datetime.md):
  A date and/or time conversion tool (like `printf` but for date and time values)
* [Filter By Range `[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Reverse Array (`mtac`)](../commands/mtac.md):
  Reverse the order of an array
* [Special Ranges](../mkarray/special.md):
  Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)
* [index](../parser/item-index.md):
  Outputs an element from an array, map or table

<hr/>

This document was generated from [builtins/core/mkarray/ranges_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/mkarray/ranges_doc.yaml).