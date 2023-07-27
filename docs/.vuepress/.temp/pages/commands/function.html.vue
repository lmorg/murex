<template><div><h1 id="function" tabindex="-1"><a class="header-anchor" href="#function" aria-hidden="true">#</a> <code v-pre>function</code></h1>
<blockquote>
<p>Define a function block</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>function</code> defines a block of code as a function</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Define a function:</p>
<pre><code>function: name { code-block }
</code></pre>
<p>Define a function with variable names defined (<strong>default value</strong> and
<strong>description</strong> are optional parameters):</p>
<pre><code>function: name (
    variable1: data-type [default-value] &quot;description&quot;,
    variable2: data-type [default-value] &quot;description&quot;
) {
    code-block
}
</code></pre>
<p>Undefine a function:</p>
<pre><code>!function: command
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» function hw { out "Hello, World!" }
» hw
Hello, World!

» !function hw
» hw
exec: "hw": executable file not found in $PATH
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="allowed-characters" tabindex="-1"><a class="header-anchor" href="#allowed-characters" aria-hidden="true">#</a> Allowed characters</h3>
<p>Function names can only include any characters apart from dollar (<code v-pre>$</code>).
This is to prevent functions from overwriting variables (see the order of
preference below).</p>
<h3 id="undefining-a-function" tabindex="-1"><a class="header-anchor" href="#undefining-a-function" aria-hidden="true">#</a> Undefining a function</h3>
<p>Like all other definable states in Murex, you can delete a function with
the bang prefix (see the example above).</p>
<h3 id="using-parameterized-variable-names" tabindex="-1"><a class="header-anchor" href="#using-parameterized-variable-names" aria-hidden="true">#</a> Using parameterized variable names</h3>
<p>By default, if you wanted to query the parameters passed to a Murex function
you would have to use either:</p>
<ul>
<li>
<p>the Bash syntax where of <code v-pre>$2</code> style numbered reserved variables,</p>
</li>
<li>
<p>and/or the Murex convention of <code v-pre>$PARAM</code> / <code v-pre>$ARGS</code> arrays (see <strong>reserved-vars</strong>
document below),</p>
</li>
<li>
<p>and/or the older Murex convention of the <code v-pre>args</code> builtin for any flags.</p>
</li>
</ul>
<p>Starting from Murex <code v-pre>2.7.x</code> it's been possible to declare parameters from
within the function declaration:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>function: name (
    variable1: data-type [default-value] "description",
    variable2: data-type [default-value] "description"
) {
    code-block
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h4 id="syntax" tabindex="-1"><a class="header-anchor" href="#syntax" aria-hidden="true">#</a> Syntax</h4>
<p>First off, the syntax doesn't have to follow exactly as above:</p>
<ul>
<li>
<p><strong>Variables</strong> shouldn't be prefixed with a dollar (<code v-pre>$</code>). This is a little like
declaring variables via <code v-pre>set</code>, etc. However it should be followed by a colon
(<code v-pre>:</code>) or comma (<code v-pre>,</code>). Normal rules apply with regards to allowed characters
in variable names: limited to ASCII letters (upper and lower case), numbers,
underscore (<code v-pre>_</code>), and hyphen (<code v-pre>-</code>). Unicode characters as variable names are
not currently supported.</p>
</li>
<li>
<p><strong>data-type</strong> is the Murex data type. This is an optional field in version
<code v-pre>2.8.x</code> (defaults to <code v-pre>str</code>) but is required in <code v-pre>2.7.x</code>.</p>
</li>
<li>
<p>The <strong>default value</strong> must be inside square brackets (<code v-pre>[...]</code>). Any value is
allowed (including Unicode) <em>except</em> for carriage returns / new lines (<code v-pre>\r</code>,
<code v-pre>\n</code>) and a closing square bracket (<code v-pre>]</code>) as the latter would indicate the end
of this field. You cannot escape these characters either.</p>
<p>This field is optional.</p>
</li>
<li>
<p>The <strong>description</strong> must sit inside double quotes (<code v-pre>&quot;...&quot;</code>). Any value is allowed
(including Unicode) <em>except</em> for carriage returns / new lines (<code v-pre>\r</code>, <code v-pre>\n</code>)
and double quotes (<code v-pre>&quot;</code>) as the latter would indicate the end of this field.
You cannot escape these characters either.</p>
<p>This field is optional.</p>
</li>
<li>
<p>You do not need a new line between each parameter, however you do need to
separate them with a comma (like with JSON, there should not be a trailing
comma at the end of the parameters). Thus the following is valid:
<code v-pre>variable1: data-type, variable2: data-type</code>.</p>
</li>
</ul>
<h4 id="variables" tabindex="-1"><a class="header-anchor" href="#variables" aria-hidden="true">#</a> Variables</h4>
<p>Any variable name you declare in your function declaration will be exposed in
your function body as a local variable. For example:</p>
<pre><code>function: hello (name: str) {
    out: &quot;Hello $name, pleased to meet you.&quot;
}
</code></pre>
<p>If the function isn't called with the complete list of parameters and it is
running in the foreground (ie not part of <code v-pre>autocomplete</code>, <code v-pre>event</code>, <code v-pre>bg</code>, etc)
then you will be prompted for it's value. That could look something like this:</p>
<pre><code>» function: hello (name: str) {
»     out: &quot;Hello $name, pleased to meet you.&quot;
» }

» hello
Please enter a value for 'name': Bob
Hello Bob, pleased to meet you.
</code></pre>
<p>(in this example you typed <code v-pre>Bob</code> when prompted)</p>
<h4 id="data-types" tabindex="-1"><a class="header-anchor" href="#data-types" aria-hidden="true">#</a> Data-Types</h4>
<p>This is the Murex data type of the variable. From version <code v-pre>2.8.x</code> this field
is optional and will default to <code v-pre>str</code> when omitted.</p>
<p>The advantage of setting this field is that values are type checked and the
function will fail early if an incorrect value is presented. For example:</p>
<pre><code>» function: age (age: int) { out: &quot;$age is a great age.&quot; }

» age
Please enter a value for 'age': ten
Error in `age` ( 2,1): cannot convert parameter 1 'ten' to data type 'int'

» age ten
Error in `age` ( 2,1): cannot convert parameter 1 'ten' to data type 'int'
</code></pre>
<p>However it will try to automatically convert values if it can:</p>
<pre><code>» age 1.2
1 is a great age.
</code></pre>
<h4 id="default-values" tabindex="-1"><a class="header-anchor" href="#default-values" aria-hidden="true">#</a> Default values</h4>
<p>Default values are only relevant when functions are run interactively. It
allows the user to press enter without inputting a value:</p>
<pre><code>» function: hello (name: str [John]) { out: &quot;Hello $name, pleased to meet you.&quot; }

» hello
Please enter a value for 'name' [John]:
Hello John, pleased to meet you.
</code></pre>
<p>Here no value was entered so <code v-pre>$name</code> defaulted to <code v-pre>John</code>.</p>
<p>Default values will not auto-populate when the function is run in the
background. For example:</p>
<pre><code>» bg {hello}
Error in `hello` ( 2,2): cannot prompt for parameters when a function is run in the background: too few parameters
</code></pre>
<h4 id="description-1" tabindex="-1"><a class="header-anchor" href="#description-1" aria-hidden="true">#</a> Description</h4>
<p>Descriptions are only relevant when functions are run interactively. It allows
you to define a more useful prompt should that function be called without
sufficient parameters. For example:</p>
<pre><code>» function hello (name: str &quot;What is your name?&quot;) { out &quot;Hello $name&quot; }

» hello
What is your name?: Sally
Hello Sally
</code></pre>
<h3 id="order-of-precedence" tabindex="-1"><a class="header-anchor" href="#order-of-precedence" aria-hidden="true">#</a> Order of precedence</h3>
<p>There is an order of precedence for which commands are looked up:</p>
<ol>
<li>
<p><code v-pre>runmode</code>: this is executed before the rest of the script. It is invoked by
the pre-compiler forking process and is required to sit at the top of any
scripts.</p>
</li>
<li>
<p><code v-pre>test</code> and <code v-pre>pipe</code> functions also alter the behavior of the compiler and thus
are executed ahead of any scripts.</p>
</li>
<li>
<p>private functions - defined via <code v-pre>private</code>. Private's cannot be global and
are scoped only to the module or source that defined them. For example, You
cannot call a private function directly from the interactive command line
(however you can force an indirect call via <code v-pre>fexec</code>).</p>
</li>
<li>
<p>Aliases - defined via <code v-pre>alias</code>. All aliases are global.</p>
</li>
<li>
<p>Murex functions - defined via <code v-pre>function</code>. All functions are global.</p>
</li>
<li>
<p>Variables (dollar prefixed) which are declared via <code v-pre>global</code>, <code v-pre>set</code> or <code v-pre>let</code>.
Also environmental variables too, declared via <code v-pre>export</code>.</p>
</li>
<li>
<p>globbing: however this only applies for commands executed in the interactive
shell.</p>
</li>
<li>
<p>Murex builtins.</p>
</li>
<li>
<p>External executable files</p>
</li>
</ol>
<p>You can override this order of precedence via the <code v-pre>fexec</code> and <code v-pre>exec</code> builtins.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>function</code></li>
<li><code v-pre>!function</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/reserved-vars.html">Reserved Variables</RouterLink>:
Special variables reserved by Murex</li>
<li><RouterLink to="/commands/alias.html"><code v-pre>alias</code></RouterLink>:
Create an alias for a command</li>
<li><RouterLink to="/commands/args.html"><code v-pre>args</code> </RouterLink>:
Command line flag parser for Murex shell scripting</li>
<li><RouterLink to="/commands/break.html"><code v-pre>break</code></RouterLink>:
Terminate execution of a block within your processes scope</li>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/commands/export.html"><code v-pre>export</code></RouterLink>:
Define an environmental variable and set it's value</li>
<li><RouterLink to="/commands/fexec.html"><code v-pre>fexec</code> </RouterLink>:
Execute a command or function, bypassing the usual order of precedence.</li>
<li><RouterLink to="/commands/g.html"><code v-pre>g</code></RouterLink>:
Glob pattern matching for file system objects (eg <code v-pre>*.txt</code>)</li>
<li><RouterLink to="/commands/global.html"><code v-pre>global</code></RouterLink>:
Define a global variable and set it's value</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/method.html"><code v-pre>method</code></RouterLink>:
Define a methods supported data-types</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
<li><RouterLink to="/commands/source.html"><code v-pre>source</code> </RouterLink>:
Import Murex code from another file of code block</li>
<li><RouterLink to="/commands/version.html"><code v-pre>version</code> </RouterLink>:
Get Murex version</li>
</ul>
</div></template>


