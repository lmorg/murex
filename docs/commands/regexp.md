# `regexp` - Command Reference

> Regexp tools for arrays / lists of strings

## Description

`regexp` provides a few tools for text matching and manipulation against an
array or list of strings - thus `regexp` is _murex_ data-type aware.

## Usage

    <stdin> -> regexp expression -> <stdout>

## Examples

### Find elements

    » ja: [monday..sunday] -> regexp 'f/^([a-z]{3})day/'
    [
        "mon",
        "fri",
        "sun"
    ]
    
This returns only 3 days because only 3 days match the expression (where
the days have to be 6 characters long) and then it only returns the first 3
characters because those are inside the parenthesis.

### Match elements

Elements containing

    » ja: [monday..sunday] -> regexp 'm/(mon|fri|sun)day/'
    [
        "monday",
        "friday",
        "sunday"
    ]
    
Elements excluding

    » ja: [monday..sunday] -> !regexp 'm/(mon|fri|sun)day/'
    [
        "tuesday",
        "wednesday",
        "thursday",
        "saturday"
    ]
    
### Substitute expression

    » ja: [monday..sunday] -> regexp 's/day/night/'
    [
        "monnight",
        "tuesnight",
        "wednesnight",
        "thursnight",
        "frinight",
        "saturnight",
        "sunnight"
    ]

## Flags

* `f`
    output found expressions (doesn't support bang prefix)
* `m`
    output elements that match expression (supports bang prefix)
* `s`
    output all elements - substituting elements that match expression (doesn't support bang prefix)

## Detail

`regexp` is data-type aware so will work against lists or arrays of whichever
_murex_ data-type is passed to it via STDIN and return the output in the
same data-type.

## Synonyms

* `regexp`
* `!regexp`
* `list.regex`


## See Also

* [`2darray` ](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`count`](../commands/count.md):
  Count items in a map, list or array
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [`map` ](../commands/map.md):
  Creates a map from two data sources
* [`match`](../commands/match.md):
  Match an exact value in an array
* [`msort` ](../commands/msort.md):
  Sorts an array - data type agnostic
* [`prefix`](../commands/prefix.md):
  Prefix a string to every item in a list
* [`prepend` ](../commands/prepend.md):
  Add data to the start of an array
* [`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [`suffix`](../commands/suffix.md):
  Prefix a string to every item in a list
* [`ta` (mkarray)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type