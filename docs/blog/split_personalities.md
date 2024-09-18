# The Split Personalities of Shell Usage

> Shell usage is split between the need to write something quickly and frequently verses the need to write something more complex but only the once. In this article is explore those opposing use cases and how different $SHELLs have chosen to address them.

## A Very Brief History

![Thompson (sitting) and Ritchie working together at a PDP-11](https://nojs.murex.rocks/images/blog/split_personalities/thompson.jpg?v=undef)

In the very early days of UNIX you had the Thompson shell which supported
pipes, some basic control structures and wildcards. Thompson shell was based
after the Multics shell, which in turn was inspired from `RUNCOM`. In fact the
'rc' extension often seen in shell profiles is directly taken from `RUNCOM`.

It wasn't until a little later that variables were a feature in shells. That
came with the PWB shell, which was designed to be upwardly-compatible with the
Thompson shell, supporting Thompson syntax while bringing advancements intended
to make shell scripting much more practical.

While the inspiration behind modern shells, `RUNCOM`, is a program that
literally just ran commands from a file; it is this authors opinion that early
UNIX shells were originally designed to be interactive terminals for launching
applications first and foremost, with scripting as a feature that took a few
years to mature. Furthermore, the ALGOL-inspired scripting commands were
originally external executables and only later rewritten as shell builtins for
performance reasons. For example running `if` in the shell would originally
`fork()` the executable `/bin/if` but that quickly became call a builtin
function that was part of the shell itself.

I believe it is these reasons why $SHELLs based on that lineage, be it the
Bourne shell, Bash or Zsh, all share a scripting syntax which very much feels
like it is extended from REPL usage.

## Opposing Requirements

![Opposing Requirements](https://nojs.murex.rocks/images/blog/split_personalities/conflict.png?v=undef)

The problem with shell usage is it falls into two contradictory categories
equally:

1. You need an interactive terminal that is optimized for the operators
   productivity. Since it is a REPL environment, any instructions you do pass
   are going to be write-many read-once. In other words, the syntax needs to be
   quick to type because it's going to be typed often. However it doesn't have
   to be particularly readable because you're not going to save and read back
   whatever instructions you've keyed into the REPL.

2. You need the ability to write short scripts. The language here needs to be
   familiar because it is aimed at non-developers (otherwise they might just as
   well use C, FORTRAN, ALGOL or others) and succinct (again, otherwise a
   developer might as well use a compiled language). However it also should be
   readable because scripts are saved, recalled, reused and often extended over
   time. So they fall into the write-once read-many category.

In an interactive program manager it makes sense to forgo quotation marks
around strings, commas to separate parameters and semi-colons to terminate the
line. Even the C shell, `csh` then later `tcsh`, doesn't follow C's syntax that
strictly -- instead understanding that brevity is required for interactive use.

When I first started writing my own shell, Murex, I originally started out
with syntax that was inspired by the C. A pipeline would look something like
the following:

```
cat ("./example.csv") | grep ("-n", "foobar")
```

While this came with some readability improvements, it was a _massive_ pain to
write over and over. So I added some syntax completion to the terminal,
inspired by IDE's and how they attempt to minimize the repetition of entering
syntax tokens. However this didn't remove the pain entirely, it just masked it
a little. So I removed the redundant braces. But the enforced quotation marks
were still annoying, so I decided to make the quotation marks optional. Then
the commas were removed...and before I knew it, I'd basically just reinvented
the same syntax for writing commands as everyone had already been using for a
multitude of decades prior. What started out as the example above ended up
looking more like the example below:

```
cat ./example.csv | grep -n foobar
```

(please excuse the useless use of `cat` in these examples -- it's purely there
for illustrative reasons)

## The Traditional

![The Traditional](https://nojs.murex.rocks/images/blog/split_personalities/old.jpg?v=undef)

As I've already hinted in the section before, Bourne, Bash, Zsh all fall nicely
into the first camp. The write-many read-once camp. And that makes sense to me
when I consider the evolution of those shells. Their heritage does stem from
interactive terminals firstly and scripting secondly.

The problem with traditional shells is that their grammar is lousy for anyone
who needs a write-once read-many language. Worse still, while a significant
amount of their grammar has now been included as builtins, for practical use
operators often find themselves inlining other languages anyway, such as awk,
sed, Perl and others. So it is understandable that a great many chose to do
away with traditional shells for scripting and instead use more other, more
powerful and readable languages like Python.

Unfortunately the same problems transfer the other way too, in that I have
already demonstrated why Python (and other programming languages) don't always
make good shells. While I will conceded that there is a loyal fanbase who will
swear by their Python REPL for terminal usage, and if they're happy with that
then I salute them, their usage is as niche as those who enjoy using Bash for
complex scripts. Perhaps the only language I've used which translates well both
for terse REPLs and lengthier scripts is LISP.

## The Modern

![The Modern](https://nojs.murex.rocks/images/blog/split_personalities/new.jpg?v=undef)

So how are modern shells addressing these split concerns?

### Powershell

Microsoft had the benefit of being able to start from a clean room. They didn't
need to inherit 50+ years of UNIX legacy when they wrote Powershell. So their
approach was naturally to base their shell on .NET. Passing .NET objects around
has a number of advantages over the POSIX specification of passing files, byte
streams, to applications. This allows developers to write richer command line
applications in their preferred .NET language rather than being tied to the
shell's syntax. However one could argue the same is true with POSIX shells and
how you can write a program in any language you like. But in Powershell those
other .NET programs feel more tightly integrated into Powershell than a forked
process does in Bash. Again, I put this down to Powershell passing .NET objects
along the pipeline.

Where Powershell falls down for me is in two key areas:

1. Many of the flags passed are verbose. Calling .NET objects can be verbose.
   Take this example of base64 encoding a string:
   ```
   [Convert]::ToBase64String([System.Text.Encoding]::Unicode.GetBytes("TextToEncode"))
   ```

2. Powershell doesn't play nicely with POSIX. Okay, I'm arguably contradicting
   myself now because earlier I raised this as a benefit. And in many ways it
   is. However if you wish to run Powershell on Linux, which you can do, you
   may find that you'll want to work with CLI tools that do "think" in terms of
   byte streams. Many of these tools have equivalent aliases written in .NET so
   you can appear to use them without escaping the rich programming environment
   provided by Powershell. However you may, and I often did, run into a great
   many scenarios where my expectations didn't match the practicalities of
   Powershell.

(I will talk more about the second point in another article where I'll discuss
pipelines, data types and the need for modern shells to understand rich data
rather than treating everything as a flat stream of bytes)

There is no question that Powershell is a more powerful REPL than Bash but it
definitely slides more towards the "write-once read-many" end of the spectrum.

### Oil

[Oil](https://www.oilshell.org/) describes itself as the following:

> Oil is a new Unix shell. It's our upgrade path from bash to a better language
> and runtime. It's also for Python and JavaScript users who avoid shell!

The way Oil achieves this is a lot of how PWB improved upon the Thompson shell
in the 1970s. Oil aims to be upwards-compatible with Bash. Any command line or
shell script you can run in Bash should, eventually, be supported in Oil as
well. Oil can extend on that and support a syntax and grammar that is more
readable and sane to write longer lived scripts in. Thus bridging the conflict
between "write-many" and "read-many" languages.

This make Oil one of the most interesting alternative shells I have come
across.

### Murex

![Murex](https://nojs.murex.rocks/images/blog/split_personalities/murex.png?v=undef)

The approach Murex takes sits somewhere in between the previous two shells.
It attempts to retain familiarity with POSIX syntax but isn't afraid to break
compatibility where it makes sense. The emphasis is on creating grammar that
is both succinct but also readable. This mission was driven from originally
attempting to create something more familiar to Javascript developers then
falling back to some old Bash-ism's when I realized that for all of it's warts,
Bash and its kin aren't actually bad for quick REPL usage of C-style braces
over ALGOL style named scopes:

**POSIX:**

```
if [ 0 -eq 1 ]; then 
    echo '0 == 1'
else
    echo '0 != 1'
fi
```

**Murex:**

```
if { 0 == 1 } then {
    echo '0 == 1'
} else {
    echo '0 != 1'
}
```

But since the curly braces are tokens, grammar like `then` / `else` become
superfluous words that only exist for readability. So then we can make them
optional. And you end up with a syntax that allows for a certain amount of
golfing in the REPL should the operator want to save a few key strokes

```
if { 0 == 1 } { echo '0 == 1' } { echo '0 != 1' }
```

## Conclusion

The write-many read-once tendencies of the interactive terminal and the
write-once read-many demands of scripting might be difficult to consolidate
but I do think it is achievable and I'm not convinced the current heavy weights
do a good job at addressing those conflicting concerns. Whereas alternative
shells like [Oil](https://www.oilshell.org/), [Elfish](https://elv.sh/) and
[Murex](https://github.com/lmorg/murex) seem to be putting a lot more thought
into this problem and it is really exciting seeing the different ideas that are
being produced.

<hr>

Published: 02.10.2021 at 22:42

## See Also

* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Reading Lists From The Command Line](../blog/reading_lists.md):
  How hard can it be to read a list of data from the command line? If your list is line delimited then it should be easy. However what if your list is a JSON array? This post will explore how to work with lists in a different command line environments.
* [Rosetta Stone](../user-guide/rosetta-stone.md):
  A tabulated list of Bashism's and their equivalent Murex syntax

<hr/>

This document was generated from [gen/blog/split_personalities_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/blog/split_personalities_doc.yaml).