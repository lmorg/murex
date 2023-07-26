<template><div><h1 id="language-tour" tabindex="-1"><a class="header-anchor" href="#language-tour" aria-hidden="true">#</a> Language Tour</h1>
<h2 id="introduction" tabindex="-1"><a class="header-anchor" href="#introduction" aria-hidden="true">#</a> Introduction</h2>
<p>Murex is a typed shell. By this we mean it still passes byte streams along
POSIX pipes (and thus will work with all your existing command line tools) but
in addition will add annotations to describe the type of data that is being
written and read. This allows Murex to expand upon your command line tools
with some really interesting and advanced features not available in traditional
shells.</p>
<blockquote>
<p>POSIX is a set of underlying standards that Linux, macOS and various other
operating systems support. Most typed shells do not work well with existing
commands whereas Murex does.</p>
</blockquote>
<h3 id="read–eval–print-loop" tabindex="-1"><a class="header-anchor" href="#read–eval–print-loop" aria-hidden="true">#</a> Read–Eval–Print Loop</h3>
<p>If you want to learn more about the interactive shell then there is a dedicated
document detailing <RouterLink to="/user-guide/interactive-shell.html">Murex's REPL features</RouterLink>.</p>
<h3 id="barewords" tabindex="-1"><a class="header-anchor" href="#barewords" aria-hidden="true">#</a> Barewords</h3>
<p>Shells need to <RouterLink to="/blog/split_personalities.html">balance scripting with an efficient interactive terminal</RouterLink>
interface. One of the most common approaches to solving that conflict between
readability and terseness is to make heavy use of barewords. Barewords are
ostensibly just instructions that are not quoted. In our case, command names
and command parameters.</p>
<p>Murex also makes heavy use of barewords and so that places requirements on
the choice of syntax we can use.</p>
<h3 id="expressions-and-statements" tabindex="-1"><a class="header-anchor" href="#expressions-and-statements" aria-hidden="true">#</a> Expressions and Statements</h3>
<p>An <strong>expression</strong> is an evaluation, operation or assignment, for example:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» 6 > 5
» fruit = %[ apples oranges bananas ]
» 5 + 5
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><blockquote>
<p>Expressions are type sensitive</p>
</blockquote>
<p>Whereas a <strong>statement</strong> is a shell command to execute:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» echo "Hello Murex"
» kill 1234
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><blockquote>
<p>All values in a statement are treated as strings</p>
</blockquote>
<p>Due to the expectation of shell commands supporting bareword parameters,
expressions have to be parsed differently to statements. Thus Murex first
parses a command line to see if it is a valid expression, and if it is not, it
then assumes it is an statement and parses it as such.</p>
<p>This allow expressions and statements to be used interchangeably in a pipeline:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» 5 + 5 | grep 10
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h3 id="functions-and-methods" tabindex="-1"><a class="header-anchor" href="#functions-and-methods" aria-hidden="true">#</a> Functions and Methods</h3>
<p>A <strong>function</strong> is command that doesn't take data from STDIN whereas a <strong>method</strong>
is any command that does.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>echo "Hello Murex" | grep "Murex"
^ a function         ^ a method
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>In practical terms, functions and methods are executed in exactly the same way
however some builtins might behave differently depending on whether values are
passed via STDIN or as parameters. Thus you will often find references to
functions and methods, and sometimes for the same command, within these
documents.</p>
<h3 id="the-bang-prefix" tabindex="-1"><a class="header-anchor" href="#the-bang-prefix" aria-hidden="true">#</a> The Bang Prefix</h3>
<p>Some Murex builtins support a bang prefix. This prefix alters the behavior of
those builtins to perform the conceptual opposite of their primary role.</p>
<p>For example, you could grep a file with <code v-pre>regexp 'm/(dogs|cats)/'</code> but then you
might want to exclude any matches by using <code v-pre>!regexp 'm/(dogs|cats)/'</code> instead.</p>
<p>The details for each supported bang prefix will be in the documents for their
respective builtin.</p>
<h2 id="rosetta-stone" tabindex="-1"><a class="header-anchor" href="#rosetta-stone" aria-hidden="true">#</a> Rosetta Stone</h2>
<p>If you already know Bash and looking for the equivalent syntax in Murex, then
our <a href="/rosetta">Rosetta Stone</a> reference will help you to
translate your Bash code into Murex code.</p>
<h2 id="basic-syntax" tabindex="-1"><a class="header-anchor" href="#basic-syntax" aria-hidden="true">#</a> Basic Syntax</h2>
<h3 id="quoting-strings" tabindex="-1"><a class="header-anchor" href="#quoting-strings" aria-hidden="true">#</a> Quoting Strings</h3>
<blockquote>
<p>It is important to note that all strings in expressions are quoted whereas
strings in statements can be barewords.</p>
</blockquote>
<p>There are three ways to quote a string in Murex:</p>
<ul>
<li><code v-pre>'single quote'</code>: use this for string literals (<RouterLink to="/parser/single-quote.html">read more</RouterLink>)</li>
<li><code v-pre>&quot;double quote&quot;</code>: use this for infixing variables (<RouterLink to="/parser/double-quote.html">read more</RouterLink>)</li>
<li><code v-pre>%(brace quote)</code>: use this for nesting quotes (<RouterLink to="/parser/brace-quote.html">read more</RouterLink>)</li>
</ul>
<h3 id="code-comments" tabindex="-1"><a class="header-anchor" href="#code-comments" aria-hidden="true">#</a> Code Comments</h3>
<p>You can comment out a single like, or end of a line with <code v-pre>#</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code># this is a comment

echo Hello Murex # this is also a comment
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Multiple lines or mid-line comments can be achieved with <code v-pre>/#</code> and <code v-pre>#/</code> tokens:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>/#
This is
a multi-line
command
#/
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>...which can also be inlined...</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» echo Hello /# comment #/ Murex
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>(<code v-pre>/#</code> was chosen because it is similar to C-style comments however <code v-pre>/*</code> is a
valid glob so Murex has substituted the asterisks with a hash symbol instead)</p>
<h2 id="variables" tabindex="-1"><a class="header-anchor" href="#variables" aria-hidden="true">#</a> Variables</h2>
<p>All variables can be defined as expressions and their data types are inferred:</p>
<ul>
<li><code v-pre>name = &quot;bob&quot;</code></li>
<li><code v-pre>age = 20 * 2</code></li>
<li><code v-pre>fruit = %[ apples oranges bananas ]</code></li>
</ul>
<p>If any variables are unset then reading from them will produce an error (under
Murex's default behavior):</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» echo $foobar
Error in `echo` (1,1): variable 'foobar' does not exist
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="global-variables" tabindex="-1"><a class="header-anchor" href="#global-variables" aria-hidden="true">#</a> Global variables</h3>
<p>Global variables can be defined using the <code v-pre>$GLOBAL</code> namespace:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» $GLOBAL.foo = "bar"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>You can also force Murex to read the global assignment of <code v-pre>$foo</code> (ignoring
any local assignments, should they exist) using the same syntax. eg:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» $GLOBAL.name = "Tom"
» out $name
Tom

» $name = "Sally"
» out $GLOBAL.name
Tom
» out $name
Sally
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="environmental-variables" tabindex="-1"><a class="header-anchor" href="#environmental-variables" aria-hidden="true">#</a> Environmental Variables</h3>
<p>Environmental Variables are like global variables except they are copied to any
other programs that are launched from your shell session.</p>
<p>Environmental variables can be assigned using the <code v-pre>$ENV</code> namespace:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» $ENV.foo = "bar"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>as well as using the <code v-pre>export</code> statement like with traditional shells. (<RouterLink to="/commands/export.html">read more</RouterLink>)</p>
<p>Like with global variables, you can force Murex to read the environmental
variable, bypassing and local or global variables of the same name, by also
using the <code v-pre>$ENV</code> namespace prefix.</p>
<h3 id="type-inference" tabindex="-1"><a class="header-anchor" href="#type-inference" aria-hidden="true">#</a> Type Inference</h3>
<p>In general, Murex will try to infer the data type of a variable or pipe. It
can do this by checking the <code v-pre>Content-Type</code> HTTP header, the file name
extension or just looking at how that data was constructed (when defined via
expressions). However sometimes you may need to annotate your types. <RouterLink to="/commands/set.html#type-annotations">Read more</RouterLink></p>
<h3 id="scalars" tabindex="-1"><a class="header-anchor" href="#scalars" aria-hidden="true">#</a> Scalars</h3>
<p>In traditional shells, variables are expanded in a way that results in spaces
be parsed as different command parameters. This results in numerous problems
where developers need to remember to enclose variables inside quotes.</p>
<p>Murex parses variables as tokens and expands them into the command line
arguments intuitively. So, there are no more accidental bugs due to spaces in
file names, or other such problems due to developers forgetting to quote
variables. For example:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» file = "file name.txt"
» touch $file # this would normally need to be quoted
» ls
'file name.txt'
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="arrays" tabindex="-1"><a class="header-anchor" href="#arrays" aria-hidden="true">#</a> Arrays</h3>
<p>Due to variables not being expanded into arrays by default, Murex supports an
additional variable construct for arrays. These are <code v-pre>@</code> prefixed:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» files = %[file1.txt, file2.txt, file3.txt]
» touch @files
» ls
file1.txt  file2.txt
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="piping-and-redirection" tabindex="-1"><a class="header-anchor" href="#piping-and-redirection" aria-hidden="true">#</a> Piping and Redirection</h2>
<h3 id="pipes" tabindex="-1"><a class="header-anchor" href="#pipes" aria-hidden="true">#</a> Pipes</h3>
<p>Murex supports multiple different pipe tokens. The main two being <code v-pre>|</code> and
<code v-pre>-&gt;</code>.</p>
<ul>
<li>
<p><code v-pre>|</code> works exactly the same as in any normal shell (<RouterLink to="/parser/pipe-posix.html">read more</RouterLink>)</p>
</li>
<li>
<p><code v-pre>-&gt;</code> displays all of the supported methods (commands that support the output
of the previous command). Think of it a little like object orientated
programming where an object will have functions (methods) attached. (<RouterLink to="/parser/pipe-arrow.html">read more</RouterLink>)</p>
</li>
</ul>
<p>In Murex scripts you can use <code v-pre>|</code> and <code v-pre>-&gt;</code> interchangeably, so there's no need
to remember which commands are methods and which are not. The difference only
applies in the interactive shell where <code v-pre>-&gt;</code> can be used with tab-autocompletion
to display a shortlist of supported functions that can manipulate the data from
the previous command. It's purely a clue to the parser to generate different
autocompletion suggestions to help with your discovery of different commandline
tools.</p>
<h3 id="redirection" tabindex="-1"><a class="header-anchor" href="#redirection" aria-hidden="true">#</a> Redirection</h3>
<p>Redirection of stdout and stderr is very different in Murex. There is no
support for the <code v-pre>2&gt;</code> or <code v-pre>&amp;1</code> tokens, instead you name the pipe inside angle
brackets, in the first parameter(s).</p>
<p><code v-pre>out</code> is that processes stdout (fd1), <code v-pre>err</code> is that processes stderr (fd2), and
<code v-pre>null</code> is the equivalent of piping to <code v-pre>/dev/null</code>.</p>
<p>Any pipes prefixed by a bang means reading from that processes stderr.</p>
<p>So to redirect stderr to stdout you would use <code v-pre>&lt;!out&gt;</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>err &lt;!out> "error message redirected to stdout"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>And to redirect stdout to stderr you would use <code v-pre>&lt;err&gt;</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>out &lt;err> "output redirected to stderr"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Likewise you can redirect either stdout, or stderr to <code v-pre>/dev/null</code> via <code v-pre>&lt;null&gt;</code>
or <code v-pre>&lt;!null&gt;</code> respectively.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>command &lt;!null> # ignore stderr
command &lt;null>  # ignore stdout
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>You can also create your own pipes that are files, network connections, or any
other custom data input or output endpoint. <RouterLink to="/user-guide/namedpipes.html">read more</RouterLink></p>
<h3 id="redirecting-to-files" tabindex="-1"><a class="header-anchor" href="#redirecting-to-files" aria-hidden="true">#</a> Redirecting to files</h3>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>out "message" |> truncate-file.txt
out "message" >> append-file.txt
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="type-conversion" tabindex="-1"><a class="header-anchor" href="#type-conversion" aria-hidden="true">#</a> Type Conversion</h3>
<p>Aside from annotating variables upon definition, you can also transform data
along the pipeline.</p>
<h4 id="cast" tabindex="-1"><a class="header-anchor" href="#cast" aria-hidden="true">#</a> Cast</h4>
<p>Casting doesn't alter the data, it simply changes the meta-information about
how that data should be read.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>out [1,2,3] | cast json | foreach { ... }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>There is also a little syntactic sugar to do the same:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>out [1,2,3] | :json: foreach { ... }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h4 id="format" tabindex="-1"><a class="header-anchor" href="#format" aria-hidden="true">#</a> Format</h4>
<p><code v-pre>format</code> takes the source data and reformats it into another data format:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out [1,2,3] | :json: format yaml
- 1
- 2
- 3
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="sub-shells" tabindex="-1"><a class="header-anchor" href="#sub-shells" aria-hidden="true">#</a> Sub-Shells</h2>
<p>There are two types of emendable sub-shells: strings and arrays.</p>
<ul>
<li>
<p>string sub-shells, <code v-pre>${ command }</code>, take the results from the sub-shell and
return it as a single parameter. This saves the need to encapsulate the shell
inside quotation marks.</p>
</li>
<li>
<p>array sub-shells, <code v-pre>@{ command }</code>, take the results from the sub-shell
and expand it as parameters.</p>
</li>
</ul>
<p><strong>Examples:</strong></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>touch ${ %[1,2,3] } # creates a file named '[1,2,3]'
touch @{ %[1,2,3] } # creates three files, named '1', '2' and '3'
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>The reason Murex breaks from the POSIX tradition of using backticks and
parentheses is because Murex works on the principle that everything inside
a curly bracket is considered a new block of code.</p>
<h2 id="filesystem-wildcards-globbing" tabindex="-1"><a class="header-anchor" href="#filesystem-wildcards-globbing" aria-hidden="true">#</a> Filesystem Wildcards (Globbing)</h2>
<p>While glob expansion is supported in the interactive shell, there isn't
auto-expansion of globbing in shell scripts. This is to protect against
accidental damage. Instead globbing is achieved via sub-shells using either:</p>
<ul>
<li><code v-pre>g</code> - traditional globbing (<RouterLink to="/commands/g.html">read more</RouterLink>)</li>
<li><code v-pre>rx</code> - regexp matching in current directory only (<RouterLink to="/commands/rx.html">read more</RouterLink>)</li>
<li><code v-pre>f</code> - file type matching (<RouterLink to="/commands/f.html">read more</RouterLink>)</li>
</ul>
<p><strong>Examples:</strong></p>
<p>All text files via globbing:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>g *.txt
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>All text and markdown files via regexp:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>rx '\.(txt|md)$'
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>All directories via type matching:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>f +d
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>You can also chain them together, eg all directories named <code v-pre>*.txt</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>g *.txt | f +d
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>To use them in a shell script it could look something a like this:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>rm @{g *.txt | f +s}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>(this deletes any symlinks called <code v-pre>*.txt</code>)</p>
<h2 id="brace-expansion" tabindex="-1"><a class="header-anchor" href="#brace-expansion" aria-hidden="true">#</a> Brace expansion</h2>
<p>In <a href="https://en.wikipedia.org/wiki/Bash_(Unix_shell)#Brace_expansion" target="_blank" rel="noopener noreferrer">bash you can expand lists<ExternalLinkIcon/></a>
using the following syntax: <code v-pre>a{1..5}b</code>. In Murex, like with globbing, brace
expansion is a function: <code v-pre>a: a[1..5]b</code> and supports a much wider range of lists
that can be expanded. (<RouterLink to="/commands/a.html">read more</RouterLink>)</p>
<h2 id="executables" tabindex="-1"><a class="header-anchor" href="#executables" aria-hidden="true">#</a> Executables</h2>
<h3 id="aliases" tabindex="-1"><a class="header-anchor" href="#aliases" aria-hidden="true">#</a> Aliases</h3>
<p>You can create &quot;aliases&quot; to common commands to save you a few keystrokes. For
example:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>alias gc=git commit
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p><code v-pre>alias</code> behaves slightly differently to Bash. (<RouterLink to="/commands/alias.html">read more</RouterLink>)</p>
<h3 id="public-functions" tabindex="-1"><a class="header-anchor" href="#public-functions" aria-hidden="true">#</a> Public Functions</h3>
<p>You can create custom functions in Murex using <code v-pre>function</code>. (<RouterLink to="/commands/function.html">read more</RouterLink>)</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>function gc (message: str) {
    # shorthand for `git commit`

    git commit -m $message
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="private-functions" tabindex="-1"><a class="header-anchor" href="#private-functions" aria-hidden="true">#</a> Private Functions</h3>
<p><code v-pre>private</code> functions are like <a href="#public-functions">public functions</a> except they
are only available within their own modules namespace. (<RouterLink to="/commands/private.html">read more</RouterLink>)</p>
<h3 id="external-executables" tabindex="-1"><a class="header-anchor" href="#external-executables" aria-hidden="true">#</a> External Executables</h3>
<p>External executables (including any programs located in <code v-pre>$PATH</code>) are invoked
via the <code v-pre>exec</code> builtin (<RouterLink to="/commands/exec.html">read more</RouterLink>) however if a command
isn't an expression, alias, function nor builtin, then Murex assumes it is an
external executable and automatically invokes <code v-pre>exec</code>.</p>
<p>For example the two following statements are the same:</p>
<ol>
<li><code v-pre>exec uname</code></li>
<li><code v-pre>uname</code></li>
</ol>
<p>Thus for normal day to day usage, you shouldn't need to include <code v-pre>exec</code>.</p>
<h2 id="control-structures" tabindex="-1"><a class="header-anchor" href="#control-structures" aria-hidden="true">#</a> Control Structures</h2>
<h3 id="using-if-statements" tabindex="-1"><a class="header-anchor" href="#using-if-statements" aria-hidden="true">#</a> Using <code v-pre>if</code> Statements</h3>
<p><code v-pre>if</code> can be used in a number of different ways, the most common being:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>if { true } then {
    # do something
} else {
    # do something else
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><code v-pre>if</code> supports a flexible variety of incarnation to solve different problems. (<RouterLink to="/commands/if.html">read more</RouterLink>)</p>
<h3 id="using-switch-statements" tabindex="-1"><a class="header-anchor" href="#using-switch-statements" aria-hidden="true">#</a> Using <code v-pre>switch</code> Statements</h3>
<p>Because <code v-pre>if ... else if</code> chains are ugly, Murex supports <code v-pre>switch</code> statements:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>switch $USER {
    case "Tom"   { out: "Hello Tom" }
    case "Dick"  { out: "Howdie Richard" }
    case "Sally" { out: "Nice to meet you" }

    default {
        out: "I don't know who you are"
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><code v-pre>switch</code> supports a flexible variety of different usages to solve different
problems. (<RouterLink to="/commands/switch.html">read more</RouterLink>)</p>
<h3 id="using-foreach-loops" tabindex="-1"><a class="header-anchor" href="#using-foreach-loops" aria-hidden="true">#</a> Using <code v-pre>foreach</code> Loops</h3>
<p><code v-pre>foreach</code> allows you to easily iterate through an array or list of any type: (<RouterLink to="/commands/foreach.html">read more</RouterLink>)</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>%[ apples bananas oranges ] | foreach fruit { out "I like $fruit" }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h3 id="using-formap-loops" tabindex="-1"><a class="header-anchor" href="#using-formap-loops" aria-hidden="true">#</a> Using <code v-pre>formap</code> Loops</h3>
<p><code v-pre>formap</code> loops are the equivalent of <code v-pre>foreach</code> but against map objects: (<RouterLink to="/commands/formap.html">read more</RouterLink>)</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>%{
    Bob:     {age: 10},
    Richard: {age: 20},
    Sally:   {age: 30}
} | formap name person {
    out "$name is $person[age] years old"
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="stopping-execution" tabindex="-1"><a class="header-anchor" href="#stopping-execution" aria-hidden="true">#</a> Stopping Execution</h2>
<h3 id="the-continue-statement" tabindex="-1"><a class="header-anchor" href="#the-continue-statement" aria-hidden="true">#</a> The <code v-pre>continue</code> Statement</h3>
<p><code v-pre>continue</code> will terminate execution of an inner block in iteration loops like
<code v-pre>foreach</code> and <code v-pre>formap</code>. Thus <em>continuing</em> the loop from the next iteration:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>%[1..10] | foreach i {
    if { $i == 5 } then {
        continue foreach
        # ^ jump back to the next iteration
    }

    out $i
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><code v-pre>continue</code> requires a parameter to define while block to iterate on. This means
you can use <code v-pre>continue</code> within nested loops and still have readable code. (<RouterLink to="/commands/continue.html">read more</RouterLink>)</p>
<h3 id="the-break-statement" tabindex="-1"><a class="header-anchor" href="#the-break-statement" aria-hidden="true">#</a> The <code v-pre>break</code> Statement</h3>
<p><code v-pre>break</code> will terminate execution of a block (eg <code v-pre>function</code>, <code v-pre>private</code>, <code v-pre>if</code>,
<code v-pre>foreach</code>, etc):</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>%[1..10] | foreach i {
    if { $i == 5 } then {
        break foreach
        # ^ exit foreach
    }

    out $i
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><code v-pre>break</code> requires a parameter to define while block to end. Thus <code v-pre>break</code> can be
considered to exhibit the behavior of <em>return</em> as well as <em>break</em> in other
languages:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>function example {
    if { $USER == "root" } then {
        err "Don't run this as root"
        break example
    }

    # ... do something ...
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><code v-pre>break</code> cannot exit anything above it's callers scope. (<RouterLink to="/commands/break.html">read more</RouterLink>)</p>
<h3 id="the-exit-statement" tabindex="-1"><a class="header-anchor" href="#the-exit-statement" aria-hidden="true">#</a> The <code v-pre>exit</code> Statement</h3>
<p>Terminates Murex. <code v-pre>exit</code> is not scope aware; if it is included in a function
then the whole shell will still exist and not just that function. (<RouterLink to="/commands/exit.html">read more</RouterLink>)</p>
<h3 id="signal-sigint" tabindex="-1"><a class="header-anchor" href="#signal-sigint" aria-hidden="true">#</a> Signal: SIGINT</h3>
<p>This can be invoked by pressing <code v-pre>Ctrl</code> + <code v-pre>c</code>.</p>
<h3 id="signal-sigquit" tabindex="-1"><a class="header-anchor" href="#signal-sigquit" aria-hidden="true">#</a> Signal: SIGQUIT</h3>
<p>This can be invoked by pressing <code v-pre>Ctrl</code> + <code v-pre>\</code></p>
<p>Sending SIGQUIT will terminate all running functions in the current Murex
session. Which is a handy escape hatch if your shell code starts misbehaving.</p>
<h3 id="signal-sigtstp" tabindex="-1"><a class="header-anchor" href="#signal-sigtstp" aria-hidden="true">#</a> Signal: SIGTSTP</h3>
<p>This can be invoked by pressing <code v-pre>Ctrl</code> + <code v-pre>z</code></p>
</div></template>


