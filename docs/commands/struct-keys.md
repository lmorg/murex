# _murex_ Shell Docs

## Command Reference: `struct-keys`

> Outputs all the keys in a structure as a file path

## Description

`struct-keys` outputs all of the keys in a structured data-type eg JSON, YAML,
TOML, etc.

The output is a JSON array of the keys with each value being a file path
representation of the input structure's node.

`struct-keys` accepts one optional parameter, an integer value which defines
how deep into a nest `struct-keys` should travel. The default is set at a
modest `20` for runtime performance, however this value might be subject to
change in future releases of _murex_ and thus you are recommended to set this
value in any scripts.

## Usage

    <stdin> -> struct-keys [ depth ] -> <stdout>

## Examples

    » set: json example={
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

## See Also

* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value
* [commands/formap](../commands/formap.md):
  