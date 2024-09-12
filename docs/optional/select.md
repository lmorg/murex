# Inline SQL (`select`)

> Inlining SQL into shell pipelines

## Description

`select` imports tabulated data into an in memory sqlite3 database and
executes SQL queries against the data. It returns a table of the same
data type as the input type

## Usage

```
<stdin> -> select * | ... WHERE ... -> <stdout>

select * | ... FROM file[.gz] WHERE ... -> <stdout>
```

## Examples

### Count rows from ps

List a count of all the processes running against each user ID:

```
» ps aux -> select count(*), user GROUP BY user ORDER BY 1
count(*) USER
1       _analyticsd
1       _applepay
1       _atsserver
1       _captiveagent
1       _cmiodalassistants
1       _ctkd
1       _datadetectors
1       _displaypolicyd
1       _distnote
1       _gamecontrollerd
1       _hidd
1       _iconservices
1       _installcoordinationd
1       _mdnsresponder
1       _netbios
1       _networkd
1       _reportmemoryexception
1       _timed
1       _usbmuxd
2       _appleevents
3       _assetcache
3       _fpsd
3       _nsurlsessiond
3       _softwareupdate
4       _windowserver
5       _coreaudiod
6       _spotlight
7       _locationd
144     root
308     foobar
```

## Detail

### Default Table Name

The table created is called `main`, however you do not need to include a `FROM`
condition in your SQL as Murex will inject `FROM main` into your SQL if it is
missing. In fact, it is recommended that you exclude `FROM` from your SQL
queries for the sake of brevity.

### `config` Options

`select`'s behavior is configurable:

```
» config -> [ select ]
{
    "fail-irregular-columns": {
        "Data-Type": "bool",
        "Default": false,
        "Description": "When importing a table into sqlite3, fail if there is an irregular number of columns",
        "Dynamic": false,
        "Global": false,
        "Value": false
    },
    "merge-trailing-columns": {
        "Data-Type": "bool",
        "Default": true,
        "Description": "When importing a table into sqlite3, if `fail-irregular-columns` is set to `false` and there are more columns than headings, then any additional columns are concatenated into the last column (space delimitated). If `merge-trailing-columns` is set to `false` then any trailing columns are ignored",
        "Dynamic": false,
        "Global": false,
        "Value": true
    },
    "print-headings": {
        "Data-Type": "bool",
        "Default": true,
        "Description": "Print headings when writing results",
        "Dynamic": false,
        "Global": false,
        "Value": true
    },
    "table-includes-headings": {
        "Data-Type": "bool",
        "Default": true,
        "Description": "When importing a table into sqlite3, treat the first row as headings (if `false`, headings are Excel style column references starting at `A`)",
        "Dynamic": false,
        "Global": false,
        "Value": true
    }
}
```

(See below for how to use `config`)

### Read All vs Sequential Reads

At present, `select` only supports reading the entire table from stdin before
importing that data into sqlite3. There is some prototype code being written to
support sequential imports but this is hugely experimental and not yet enabled.

This might make `select` unsuitable for large datasets.

### Early Release

This is a very early release so there almost certainly will be bugs hiding.
Which is another reason why this is currently only an optional builtin.

If you do run into any issues then please raise them on [Github](https://github.com/lmorg/murex/issues).

## Synonyms

* `select`
* `table.select`


## See Also

* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [`*` (generic)](../types/generic.md):
  generic (primitive)
* [`csv`](../types/csv.md):
  CSV files (and other character delimited tables)
* [`jsonl`](../types/jsonl.md):
  JSON Lines
* [v2.1](../changelog/v2.1.md):
  This release comes with support for inlining SQL and some major bug fixes plus a breaking change for `config`. Please read for details.

<hr/>

This document was generated from [builtins/optional/select/select_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/optional/select/select_doc.yaml).