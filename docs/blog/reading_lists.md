# Reading Lists From The Command Line

> How hard can it be to read a list of data from the command line? If your list is line delimited then it should be easy. However what if your list is a JSON array? This post will explore how to work with lists in a different command line environments.

<h2>Table of Contents</h2>

<div id="toc">

- [Preface](#preface)
- [Reading lines in Bash and similar shells](#reading-lines-in-bash-and-similar-shells)
- [But what if my files aren't line delimited?](#but-what-if-my-files-arent-line-delimited)
- [Iteration in Bash via `jq`](#iteration-in-bash-via-jq)
- [Iteration in Murex via `foreach`](#iteration-in-murex-via-foreach)
- [Reading JSON arrays in PowerShell](#reading-json-arrays-in-powershell)
- [Conclusion](#conclusion)

</div>

## Preface

A common problem we resort to shell scripting for is iterating through lists. This was easy in the days of old when most data was `\n` (new line) delimited but these days structured data is common place with formats like JSON, YAML, TOML, XML and even S-Expressions appearing commonly throughout developer and DevOps tooling.

So lets explore a few techniques for iterating through lists.

## Reading lines in Bash and similar shells

Bash shell is a popular command-line interface for Unix and Linux operating systems. One of its many useful features is the ability to read files line by line. This can be helpful for processing large files or performing repetitive tasks on a file's contents. The most basic way to read a file line by line is to use a while loop with the `read` command:

```
while read line; do
    echo $line
done < file.txt
```

In this example, the `while` loop reads each line of the `file.txt` file and stores it in the `$line` variable. The `echo` command then prints the contents of the `$line` variable to the console. The `<` symbol tells Bash to redirect the contents of the file into the loop.

The `read` command is what actually reads each line of the file. By default, it reads one line at a time and stores it in the variable specified. You can also use the `-r` option with the `read` command to disable backslash interpretation, which can be useful when dealing with files that contain backslashes.

Another useful feature of Bash is the ability to perform operations on each line of a file before processing it. For example, you can use `sed` to replace text within each line of a file:

```
while read line; do
    new_line=$(echo $line | sed 's/foo/bar/g')
    echo $new_line
done < file.txt
```

In this example, `sed` replaces all instances of "foo" with "bar" in each line of the file. The modified line is then stored in the `$new_line` variable and printed to the console.

Of course you could just run

```
sed 's/foo/bar/g' file.txt
```

...but the reasons for the for this contrived example will follow.

## But what if my files aren't line delimited?

The problem with Bash, and all traditional Linux or UNIX shells, is that they operate on byte streams. To be fair, this isn't so much a fault of Bash _per se_ but more a result of the design of UNIX where (almost) everything is a file, including pipes. This means everything is treated as bytes. Unlike, for example, Powershell which passes .NET objects around. Byte streams make complete sense when you're working on '70s or '80s mainframes but it is a little less productive in the modern world of structured formats like JSON.

So how do you read lists from objects in, for example, JSON? In Bash, this isn't so easy. You need to rely on third party tools like `jq`. However you do have the benefit of compatibility with all of the older core utilities, like `sed`, that have become muscle memory by now. This does also come with its own drawbacks as well, which I'll explore in the following section.

## Iteration in Bash via `jq`

`jq` is a fantastic tool that has become a staple of many a CI/CD pipeline however it is not part of most operating systems base platform, so it would need to be installed separately. This also creates additional complications whereby you end up having a language within a language -- like running `awk` or `sed` inside Bash, you're now introducing `jq` too. Thus its syntax isn't always the easiest to grok when delving deep into nested JSON with conditionals and such like compared with shells that offer first party tools for working with objects. We can delve deeper into the power of `jq` in another article but for now we are going to keep things intentionally simple:

Lets create a JSON array:

```
json='["Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"]'
```

Straight away you should be able to see that Bash, with its reliance on whitespace delimitations, couldn't natively parse this. So now lets run it through `jq`:

```
$ echo $json | jq -r '.[]' | while read -r day do; echo "Happy $day"; done
Happy Monday
Happy Tuesday
Happy Wednesday
Happy Thursday
Happy Friday
Happy Saturday
Happy Sunday
```

What's happening here is the `jq` tool is converting our JSON array into a `\n` delimited list. And from there, we can use `while` and `read` just like we did in our first example at the start of this article.

The `-r` flag tells `jq` to strip quotation marks around the values. Without `-r` you'd see `Happy "Monday"'` and so on and so forth.

`.[]` is `jq` syntax for "all elements (`[]`) in the root object space (`.`).

## Iteration in Murex via `foreach`

Murex doesn't just treat files as byte streams, it passes type annotations too. And it uses those annotations to dynamically alter how to read files. The following examples will also use JSON as the input format, however Murex natively supports other structured data formats too, like YAML, CSV and S-Expressions.

Lets use the same JSON array as we did earlier, except use one of Murex's features to generate arrays programmatically:

```
» %[Monday..Sunday]
[
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
    "Sunday"
]
```

The `jq` example rewritten in Murex would look like the following:

```
%[Monday..Sunday] | foreach day { out "Happy $day" }
```

What's happening here is `%[...]` creates the JSON array (as described above) and then the `foreach` builtin iterates through the array and assigns that element to a variable named `day`.

> `out` in Murex is the equivalent of `echo` in Bash. In fact you can still use `echo` in Murex albeit that is just aliased to `out`.

## Reading JSON arrays in PowerShell

Microsoft PowerShell is a typed shell, like Murex, which was originally built for Windows but has since been ported to macOS and Linux too. Where PowerShell differs is that rather than using byte streams with type annotations, PowerShell passes .NET objects. Thus you'll see a little more boilerplate code in PowerShell where you need to explicitly convert types -- whereas Murex can get away with implicit definitions.

```
$jsonString = '["Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"]'
$jsonObject = ConvertFrom-Json $jsonString
foreach ($day in $jsonObject) { Write-Host "Hello $day" } 
```

The first line is just creating a JSON string and allocating that to `$jsonString`. We can largely ignore that as it is the same as we've seen in Bash already. The second line is more interesting as that is where the type conversion happens. `ConvertFrom-Json` does exactly as it describes -- PowerShell is generally pretty good at having descriptive names for commands,

From there on it looks relatively similar to Murex syntax: `foreach` being the statement name, followed by the variable name.

## Conclusion

Iterating over a JSON array from the command line can be done in various ways using different shells. PowerShell, `jq`, and Murex offer their unique approaches and syntaxes, making it easy to work with JSON data in different environments. Whether you're a Windows user who prefers PowerShell or a Linux user who prefers Bash or Murex, there are many options available to suit your needs. Regardless of the shell you choose, mastering the art of iterating over JSON arrays can greatly enhance your command-line skills and help you work more efficiently with JSON data.

<hr>

Published: 22.04.2023 at 11:43

## See Also

* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)
* [`%[]` Array Builder](../parser/create-array.md):
  Quickly generate arrays

<hr/>

This document was generated from [gen/blog/reading_lists_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/blog/reading_lists_doc.yaml).