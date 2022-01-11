# _murex_ Shell Docs

## mkarray: Special Ranges

> Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)

## Description

Unlike bash, _murex_ also supports some special ranges:

```  
» a: [mon..sun]
» a: [monday..sunday]
» a: [jan..dec]
» a: [january..december]
» a: [spring..winter]
```

Please refer to [a (mkarray)](../commands/a.md) for more detailed usage of mkarray.

## Usage

    a: [start..end] -> <stdout>
    a: [start..end,start..end] -> <stdout>
    a: [start..end][start..end] -> <stdout>
    
All usages also work with `ja` and `ta` as well:

    ja: [start..end] -> <stdout>
    ta: data-type [start..end] -> <stdout>

## Examples

    » a: [summer..winter]
    summer
    autumn
    winter

## Detail

### Case Sensitivity

Special ranges are case aware. If the ranges are uppercase then the return will
be uppercase. If the ranges are title case (capital first letter) then the
return will be in title case.

#### lower case

    » a: [monday..wednesday]
    monday
    tuesday
    wednesday
    
#### Title Case

    » a: [Monday..Wednesday]
    Monday
    Tuesday
    Wednesday
    
#### UPPER CASE

    » a: [MONDAY..WEDNESDAY]
    MONDAY
    TUESDAY
    WEDNESDAY
    
### Looping vs Negative Ranges

Where the special ranges differ from a regular range is they cannot
cannot down. eg `a: [3..1]` would output

    » a: [3..1]
    3
    2
    1
    
however a negative range in special ranges will cycle through to the end
of the range and then loop back from the start:

    » a: [Thursday..Wednesday]
    Thursday
    Friday
    Saturday
    Sunday
    Monday
    Tuesday
    Wednesday
    
This decision was made because generally with ranges of this type, you
would more often prefer to cycle through values rather than iterate
backwards through the list.

If you did want to reverse then pipe the output into another tool:

    » a: [Monday..Friday] -> mtac
    Friday
    Thursday
    Wednesday
    Tuesday
    Monday
    
There are other UNIX tools which aren't data type aware but would work in
this specific scenario:

* `tac` (Linux),

* `tail -r` (BSD / OS X)

* `perl -e "print reverse <>"` (Multi-platform but requires Perl installed)

### Supported Dictionary Terms

Below is the source for the supported dictionary terms:

```go
package mkarray

var mapRanges = []map[string]int{
	rangeWeekdayLong,
	rangeWeekdayShort,
	rangeMonthLong,
	rangeMonthShort,
	rangeSeason,
	rangeMoon,
}

var rangeWeekdayLong = map[string]int{
	"monday":    1,
	"tuesday":   2,
	"wednesday": 3,
	"thursday":  4,
	"friday":    5,
	"saturday":  6,
	"sunday":    7,
}

var rangeWeekdayShort = map[string]int{
	"mon": 1,
	"tue": 2,
	"wed": 3,
	"thu": 4,
	"fri": 5,
	"sat": 6,
	"sun": 7,
}

var rangeMonthLong = map[string]int{
	"january":   1,
	"february":  2,
	"march":     3,
	"april":     4,
	"may":       5,
	"june":      6,
	"july":      7,
	"august":    8,
	"september": 9,
	"october":   10,
	"november":  11,
	"december":  12,
}

var rangeMonthShort = map[string]int{
	"jan": 1,
	"feb": 2,
	"mar": 3,
	"apr": 4,
	"may": 5,
	"jun": 6,
	"jul": 7,
	"aug": 8,
	"sep": 9,
	"oct": 10,
	"nov": 11,
	"dec": 12,
}

var rangeSeason = map[string]int{
	"spring": 1,
	"summer": 2,
	"autumn": 3,
	"winter": 4,
}

var rangeMoon = map[string]int{
	"new moon":        1,
	"waxing crescent": 2,
	"first quarter":   3,
	"waxing gibbous":  4,
	"full moon":       5,
	"waning gibbous":  6,
	"third quarter":   7,
	"waning crescent": 8,
}
```

## See Also

* [mkarray/Calendar Date Ranges](../mkarray/date.md):
  Create arrays of dates
* [commands/`@[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`datetime` ](../commands/datetime.md):
  A date and/or time conversion tool (like `printf` but for date and time values)
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`len` ](../commands/len.md):
  Outputs the length of an array
* [commands/`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [commands/`ta` (mkarray)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type