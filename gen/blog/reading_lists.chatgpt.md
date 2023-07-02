## Preface

One common challenge for which we often resort to shell scripting is iterating through lists. In the past, this was straightforward when most data was delimited by new lines (`\n`). However, in today's world, structured data formats like JSON, YAML, TOML, XML, and S-Expressions have become commonplace in developer and DevOps tooling. In this blog post, we will explore various techniques for iterating through lists.

## Reading lines in Bash and similar shells

The Bash shell is a widely used command-line interface for Unix and Linux operating systems. It offers a useful feature: the ability to read files line by line. This capability comes in handy when processing large files or performing repetitive tasks on a file's contents. The simplest way to read a file line by line is by employing a while loop with the `read` command:

```bash
while read line; do
    echo $line
done < file.txt
```

In this example, the `while` loop reads each line from the `file.txt` file and stores it in the `$line` variable. The `echo` command then prints the contents of the `$line` variable to the console. The `<` symbol instructs Bash to redirect the file's contents into the loop.

The `read` command is responsible for reading each line of the file. By default, it reads one line at a time and stores it in the specified variable. Additionally, you can use the `-r` option with the `read` command to disable backslash interpretation, which proves useful when dealing with files containing backslashes.

Another useful feature of Bash is the ability to perform operations on each line of a file before processing it. For instance, you can utilize `sed` to replace text within each line of a file:

```bash
while read line; do
    new_line=$(echo $line | sed 's/foo/bar/g')
    echo $new_line
done < file.txt
```

In this example, `sed` replaces all occurrences of "foo" with "bar" in each line of the file. The modified line is then stored in the `$new_line` variable and printed to the console.

Of course, you could achieve the same result with a simpler command:

```bash
sed 's/foo/bar/g' file.txt
```

However, I present this contrived example to highlight the reasons behind the approach.

## But what if my files aren't line delimited?

The issue with Bash, as well as traditional Linux or UNIX shells, is that they primarily operate on byte streams. This characteristic is not solely a flaw of Bash itself but rather a consequence of the UNIX design, where almost everything, including pipes, is treated as files. As a result, everything is handled as bytes. In contrast, modern structured formats like JSON are prevalent in today's world, and working with byte streams proves less productive in these cases.

So, how can you read lists from objects like JSON in Bash? Unfortunately, Bash does not offer native support for this functionality, and you need to rely on third-party tools like `jq`. However, Bash still provides compatibility with older core utilities such as `sed`, which have become ingrained in our workflows. Nonetheless, this approach has its drawbacks, which we will explore in the following section.

## Iteration in Bash via `jq`

`jq` is a powerful tool that has become a staple in many CI/CD pipelines. However, it is not included by default in most operating systems and requires a separate installation. Introducing `jq` means incorporating a language within a language, similar to running `awk` or `sed` inside Bash. Consequently, its syntax may not always be the easiest to grasp, particularly when working with nested JSON structures and conditionals

, compared to shells that provide first-party tools for working with objects. While the true potential of `jq` warrants another article, for now, we'll keep things intentionally simple.

Let's create a JSON array:

```bash
json='["Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"]'
```

At first glance, it's evident that Bash, with its reliance on whitespace delimitation, cannot natively parse this array. Therefore, let's process it using `jq`:

```bash
$ echo $json | jq -r '.[]' | while read -r day; do echo "Happy $day"; done
Happy Monday
Happy Tuesday
Happy Wednesday
Happy Thursday
Happy Friday
Happy Saturday
Happy Sunday
```

What's happening here is that `jq` converts our JSON array into a newline-delimited list. From there, we can use the `while` loop and `read` just like we did in the initial example of this article.

The `-r` flag instructs `jq` to remove quotation marks surrounding the values. Without `-r`, the output would show `Happy "Monday"` and so on.

In `jq` syntax, `.[]` refers to "all elements (`[]`) in the root object space (`.`)."

## Iteration in Murex via `foreach`

Murex takes a different approach than treating files as mere byte streams; it also considers type annotations. It leverages these annotations to dynamically adjust file reading methods. The following examples utilize JSON as the input format, but Murex natively supports other structured data formats such as YAML, CSV, and S-Expressions.

Let's use the same JSON array as before but leverage one of Murex's features to generate arrays programmatically:

```murex
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

The `jq` example can be rewritten in Murex as follows:

```murex
%[Monday..Sunday] | foreach day { out "Happy $day" }
```

Here, `%[...]` creates the JSON array, as described earlier, and the `foreach` built-in iterates through the array, assigning each element to the variable `day`.

> In Murex, the `out` command is equivalent to `echo` in Bash. In fact, you can still use `echo` in Murex, as it is aliased to `out`.

It's worth mentioning that since Murex version 3.1, lambdas have been available, allowing you to write code that looks like this:

```murex
$json[{out "Hello $."}]
```

But we'll delve into that in a different article.

## Reading JSON arrays in PowerShell

Microsoft PowerShell is a typed shell, similar to Murex, originally built for Windows but since ported to macOS and Linux as well. Unlike Bash, PowerShell passes .NET objects instead of byte streams with type annotations. Consequently, you may encounter slightly more boilerplate code in PowerShell, as you need to explicitly convert types, whereas Murex can rely on implicit definitions.

```powershell
$jsonString = '["Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"]'
$jsonObject = ConvertFrom-Json $jsonString
foreach ($day in $jsonObject) { Write-Host "Hello $day" }
```

In the above code, the first line creates a JSON string and assigns it to the variable `$jsonString`. We can largely overlook this step, as it is similar to what we saw in Bash. The second line is more interesting as

 it involves type conversion. The `ConvertFrom-Json` command, as its name suggests, converts the JSON string into a PowerShell object. PowerShell generally has descriptive command names.

After that, the code resembles Murex syntax: using `foreach` as the statement name, followed by the variable name.

## Conclusion

Iterating over a JSON array from the command line can be accomplished in various ways using different shells. PowerShell, `jq`, and Murex offer their unique approaches and syntaxes, making it easy to work with JSON data in different environments. Whether you're a Windows user who prefers PowerShell or a Linux user who favors Bash or Murex, there are multiple options available to suit your needs. Mastering the art of iterating over JSON arrays, regardless of the shell you choose, can significantly enhance your command-line skills and improve your efficiency when working with JSON data.