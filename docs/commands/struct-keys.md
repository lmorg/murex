# Print Map / Structure Keys (`struct-keys`)

> Outputs all the keys in a structure as a file path

## Description

`struct-keys` outputs all of the keys in a structured data-type eg JSON, YAML,
TOML, etc.

The output is a JSON array of the keys with each value being a file path
representation of the input structure's node.

## Usage

```
<stdin> -> struct-keys [ depth ] -> <stdout>

<stdin> -> struct-keys [ flags ] -> <stdout>
```

## Examples

### Source object

The source for these examples will be defined in the variable `$example`:

```
» set json example={
      "firstName": "John",
      "lastName": "Smith",
      "isAlive": true,
      "age": 27,
      "address": {
          "streetAddress": "21 2nd Street",
          "city": "New York",
          "state": "NY",
          "postalCode": "10021-3100"
      },
      "phoneNumbers": [
          {
              "type": "home",
              "number": "212 555-1234"
          },
          {
              "type": "office",
              "number": "646 555-4567"
          },
          {
              "type": "mobile",
              "number": "123 456-7890"
          }
      ],
      "children": [],
      "spouse": null
  }
```

### Without flags

```
» $example -> struct-keys
[
    "/lastName",
    "/isAlive",
    "/age",
    "/address",
    "/address/state",
    "/address/postalCode",
    "/address/streetAddress",
    "/address/city",
    "/phoneNumbers",
    "/phoneNumbers/0",
    "/phoneNumbers/0/type",
    "/phoneNumbers/0/number",
    "/phoneNumbers/1",
    "/phoneNumbers/1/number",
    "/phoneNumbers/1/type",
    "/phoneNumbers/2",
    "/phoneNumbers/2/type",
    "/phoneNumbers/2/number",
    "/children",
    "/spouse",
    "/firstName"
]
```

### --depth and --separator

Defining max depth and changing the separator string:

```
» $example -> struct-keys --depth 1 --separator '.'   
[
    ".children",
    ".spouse",
    ".firstName",
    ".lastName",
    ".isAlive",
    ".age",
    ".address",
    ".phoneNumbers"
]
```

### Unicode separator

An example of a unicode character being used as a separator:

```
» $example -> struct-keys --depth 2 --separator ☺                                                                                                                                                           
[
    "☺age",
    "☺address",
    "☺address☺streetAddress",
    "☺address☺city",
    "☺address☺state",
    "☺address☺postalCode",
    "☺phoneNumbers",
    "☺phoneNumbers☺0",
    "☺phoneNumbers☺1",
    "☺phoneNumbers☺2",
    "☺children",
    "☺spouse",
    "☺firstName",
    "☺lastName",
    "☺isAlive"
]
```

### Multi-character separator

Separator can also be multiple characters:

```
» $example -> struct-keys --depth 1 --separator '|||' 
[
    "|||firstName",
    "|||lastName",
    "|||isAlive",
    "|||age",
    "|||address",
    "|||phoneNumbers",
    "|||children",
    "|||spouse"
]
```

## Flags

* `--depth`
    How far to traverse inside the nested structure
* `--separator`
    String to use as a separator between fields (defaults to `/`)
* `-d`
    Alias for `--depth`
* `-s`
    Alias for `--separator`

## Synonyms

* `struct-keys`


## See Also

* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure

<hr/>

This document was generated from [builtins/core/datatools/structkeys_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/datatools/structkeys_doc.yaml).