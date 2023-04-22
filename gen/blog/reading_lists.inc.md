## Reading lines

Bash shell is a popular command-line interface for Unix and Linux operating systems. One of its many useful features is the ability to read files line by line. This can be helpful for processing large files or performing repetitive tasks on a file's contents. The most basic way to read a file line by line is to use a while loop with the `read` command:

```
#!/bin/bash

while read line; do
    echo $line
done < file.txt
```

In this example, the `while` loop reads each line of the `file.txt` file and stores it in the `$line` variable. The `echo` command then prints the contents of the `$line` variable to the console. The `<` symbol tells Bash to redirect the contents of the file into the loop.

The `read` command is what actually reads each line of the file. By default, it reads one line at a time and stores it in the variable specified. You can also use the `-r` option with the `read` command to disable backslash interpretation, which can be useful when dealing with files that contain backslashes.

Another useful feature of Bash is the ability to perform operations on each line of a file before processing it. For example, you can use `sed` to replace text within each line of a file:

```
#!/bin/bash

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

So how do you read lists from objects in, for example, JSON? In Bash, this isn't so easy. You need to rely on third party tools like `jq` and often end up throwing out all of the older core utilities, like `sed`, that have become muscle memory. In Powershell it is a lot easier but you then need .NET wrappers around your existing command line tools. In Murex you can have your proverbial cake and eat it.

## Structured iteration in Murex

The following examples will use JSON as the input format, however Murex natively supports other structured data formats too.

### The `foreach` builtin

Lets say you have an array that looks something like the following:

```
» %[January..December]
[
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December"
]
```

And lets say you wanted to only return months that ended in **ber** (excuse the following video but I find this track be B.E.R. to be a particularly effective earworm)

<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/MKtW-k8za7I?controls=0" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

Sure you could `grep` that in Bash but what you're left with isn't JSON. And what if that JSON is minified?

```
["January","February","March","April","May","June","July","August","September","October","November","December"]
```

Now reading each line isn't going to work.

Murex doesn't just treat files as byte streams, it passes type annotations too. And it uses those annotations to dynamically alter how to read files. So you could grep only **ber** from minified JSON with a simple `foreach { | grep "ber" }`:

```
$months = %[January..December]
$months | foreach {
    | grep ber
}
```

The best thing is `foreach` will understand arrays from JSON, YAML, S-Expressions, CSV, and others... as well as regular terminal output. So it is one syntax to learn regardless of the input file format.

### The `formap` builtin

Arrays are easy though. What about maps? Or dictionaries, objects, hashes... as other languages might refer to them. Key value pairs are always going to be harder to parse because even with file formats like YAML, arrays are line delimited where as splitting key value pairs requires extra grokking.

Murex also has an iterator for such constructs: `formap`:

```
» echo '{"a":1,"b":2,"c":3}' | :json: formap key value { echo "key=$key, value=$value" }
key=a, value=1
key=b, value=2
key=c, value=3
```

Here we are using `:json:` to cast / annotate the STDIN stream for `formap`.

### Lambdas

Lambdas were introduced in version 4.0 of Murex. Lambdas offer shortcuts around common matching problems with structured data. For example:

```
» $example = %{a:1, b:2, c:3}
» @example[{$.val == 2}]
{
    "b": 2
}
```

The way this particular lambda works is that for each element in `$example` the code block `{$.val == 2}` runs with the `$.` variable is assigned with a structure that looks like this:

```
# first iteration
{
    "key": "a",
    "val": 1
}
# second iteration
{
    "key": "b",
    "val": 2
}
# third iteration
{
    "key": "c",
    "val": 3
}
```

Thus `$.val` (eg **2**) would be compared to the number **2**. If the result is true, that element is output. If the result is false then that element is ignored.

## Conclusion

There are a multitude of ways to iterate through structured data on Linux and UNIX systems.