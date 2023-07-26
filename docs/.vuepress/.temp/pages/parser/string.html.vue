<template><div><h1 id="string-token" tabindex="-1"><a class="header-anchor" href="#string-token" aria-hidden="true">#</a> String (<code v-pre>$</code>) Token</h1>
<blockquote>
<p>Expand values as a string</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>The string token is used to tell Murex to expand variables and subshells as a
string (ie one single parameter) irrespective of the data that is stored in the
string. One handy common use case is file names where traditional POSIX shells
would treat spaces as a new file, whereas Murex treats spaces as a printable
character unless explicitly told to do otherwise.</p>
<p>The string token must be followed with one of the following characters:
alpha, numeric, underscore (<code v-pre>_</code>) or a full stop / period (<code v-pre>.</code>).</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p><strong>ASCII variable names:</strong></p>
<pre><code>» $example = &quot;foobar&quot;
» out $example
foobar
</code></pre>
<p><strong>Unicode variable names:</strong></p>
<p>Variable names can be non-ASCII however they have to be surrounded by
parenthesis. eg</p>
<pre><code>» $(比如) = &quot;举手之劳就可以使办公室更加环保，比如，使用再生纸。&quot;
» out $(比如)
举手之劳就可以使办公室更加环保，比如，使用再生纸。
</code></pre>
<p><strong>Infixing inside text:</strong></p>
<p>Sometimes you need to denote the end of a variable and have text follow on.</p>
<pre><code>» $partial_word = &quot;orl&quot;
» out &quot;Hello w$(partial_word)d!&quot;
Hello world!
</code></pre>
<p><strong>Variables are tokens:</strong></p>
<p>Please note the new line (<code v-pre>\n</code>) character. This is not split using <code v-pre>$</code>:</p>
<pre><code>» $example = &quot;foo\nbar&quot;
</code></pre>
<p>Output as a string:</p>
<pre><code>» out $example
foo
bar
</code></pre>
<p>Output as an array:</p>
<pre><code>» out @example
foo bar
</code></pre>
<p>The string and array tokens also works for subshells:</p>
<pre><code>» out ${ %[Mon..Fri] }
[&quot;Mon&quot;,&quot;Tue&quot;,&quot;Wed&quot;,&quot;Thu&quot;,&quot;Fri&quot;]

» out @{ %[Mon..Fri] }
Mon Tue Wed Thu Fri
</code></pre>
<blockquote>
<p><code v-pre>out</code> will take an array and output each element, space delimited. Exactly
the same how <code v-pre>echo</code> would in Bash.</p>
</blockquote>
<p><strong>Variable as a command:</strong></p>
<p>If a variable is used as a commend then Murex will just print the content of
that variable.</p>
<pre><code>» $example = &quot;Hello World!&quot;

» $example
Hello World!
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>Strings and subshells can be expanded inside double quotes, brace quotes as
well as used as barewords. But they cannot be expanded inside single quotes.</p>
<pre><code>» set: example=&quot;World!&quot;

» out: Hello $example
Hello World!

» out: 'Hello $example'
Hello $example

» out: &quot;Hello $example&quot;
Hello World!

» out: %(Hello $example)
Hello World!
</code></pre>
<p>However you cannot expand arrays (<code v-pre>@</code>) inside any form of quotation since
it wouldn't be clear how that value should be expanded relative to the
other values inside the quote. This is why array and object builders (<code v-pre>%[]</code>
and <code v-pre>%{}</code> respectively) support array variables but string builders (<code v-pre>%()</code>)
do not.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/array.html">Array (<code v-pre>@</code>) Token</RouterLink>:
Expand values as an array</li>
<li><RouterLink to="/parser/brace-quote.html">Brace Quote (<code v-pre>%(</code>, <code v-pre>)</code>) Tokens</RouterLink>:
Initiates or terminates a string (variables expanded)</li>
<li><RouterLink to="/parser/double-quote.html">Double Quote (<code v-pre>&quot;</code>) Token</RouterLink>:
Initiates or terminates a string (variables expanded)</li>
<li><RouterLink to="/user-guide/reserved-vars.html">Reserved Variables</RouterLink>:
Special variables reserved by Murex</li>
<li><RouterLink to="/parser/single-quote.html">Single Quote (<code v-pre>'</code>) Token</RouterLink>:
Initiates or terminates a string (variables not expanded)</li>
<li><RouterLink to="/parser/tilde.html">Tilde (<code v-pre>~</code>) Token</RouterLink>:
Home directory path variable</li>
<li><RouterLink to="/commands/brace-quote.html"><code v-pre>(</code> (brace quote)</RouterLink>:
Write a string to the STDOUT without new line</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


