<template><div><h1 id="array-token" tabindex="-1"><a class="header-anchor" href="#array-token" aria-hidden="true">#</a> Array (<code v-pre>@</code>) Token</h1>
<blockquote>
<p>Expand values as an array</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>The array token is used to tell Murex to expand the string as multiple
parameters (an array) rather than as a single parameter string.</p>
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
<p>Since arrays are expanded over multiple parameters, you cannot expand an array
inside quoted strings like you can with a string variable:</p>
<pre><code>» out: &quot;foo ${ ja: [1..5] } bar&quot;
foo [&quot;1&quot;,&quot;2&quot;,&quot;3&quot;,&quot;4&quot;,&quot;5&quot;] bar

» out: &quot;foo @{ ja: [1..5] } bar&quot;
foo  1 2 3 4 5  bar

» %(${ ja: [1..5] })
[&quot;1&quot;,&quot;2&quot;,&quot;3&quot;,&quot;4&quot;,&quot;5&quot;]

» %(@{ ja: [1..5] })
@{ ja: [1..5] }
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/brace-quote.html">Brace Quote (<code v-pre>%(</code>, <code v-pre>)</code>) Tokens</RouterLink>:
Initiates or terminates a string (variables expanded)</li>
<li><RouterLink to="/parser/double-quote.html">Double Quote (<code v-pre>&quot;</code>) Token</RouterLink>:
Initiates or terminates a string (variables expanded)</li>
<li><RouterLink to="/parser/single-quote.html">Single Quote (<code v-pre>'</code>) Token</RouterLink>:
Initiates or terminates a string (variables not expanded)</li>
<li><RouterLink to="/parser/string.html">String (<code v-pre>$</code>) Token</RouterLink>:
Expand values as a string</li>
<li><RouterLink to="/parser/tilde.html">Tilde (<code v-pre>~</code>) Token</RouterLink>:
Home directory path variable</li>
<li><RouterLink to="/commands/brace-quote.html"><code v-pre>(</code> (brace quote)</RouterLink>:
Write a string to the STDOUT without new line</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


