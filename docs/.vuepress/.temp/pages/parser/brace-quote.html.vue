<template><div><h1 id="brace-quote-tokens" tabindex="-1"><a class="header-anchor" href="#brace-quote-tokens" aria-hidden="true">#</a> Brace Quote (<code v-pre>%(</code>, <code v-pre>)</code>) Tokens</h1>
<blockquote>
<p>Initiates or terminates a string (variables expanded)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Brace quote is used to initiate and terminate strict strings where variables
can be expanded.</p>
<p>While brace quotes are untraditional compared to your typical string quotations
in POSIX shells, brace quotes have one advantage in that the open and close
grapheme differ (ie <code v-pre>(</code> is a different character to <code v-pre>)</code>). This brings benefits
when nesting quotes as it saves the developer from having to carefully escape
the nested quotation marks just the right number of times.</p>
<p>Commands cannot be quoted using brace quotes because <code v-pre>%(</code> is recognized as its
own function.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<h4 id="as-a-parameter" tabindex="-1"><a class="header-anchor" href="#as-a-parameter" aria-hidden="true">#</a> As a parameter:</h4>
<pre><code>name = %(Bob)
</code></pre>
<h4 id="as-a-function" tabindex="-1"><a class="header-anchor" href="#as-a-function" aria-hidden="true">#</a> As a function:</h4>
<pre><code>» %(hello world)
hello world
</code></pre>
<h4 id="nested-quotes" tabindex="-1"><a class="header-anchor" href="#nested-quotes" aria-hidden="true">#</a> Nested quotes:</h4>
<pre><code>» murex -c %(out: %(Hello &quot;${murex -c %(out: %(Bob))}&quot;))
Hello &quot;Bob&quot;
</code></pre>
<p>In this example we are calling Murex to execute code as a command line
parameter (the <code v-pre>-c</code> flag). That code outputs <code v-pre>Hello &quot;...&quot;</code> but inside the
double quotes is a name that is generated from a sub-shell. That sub-shell
itself runs another murex instance which also executes another command line
parameter, this time outputting the name <strong>Bob</strong>.</p>
<p>The example is contrived but it does demonstrate how you can heavily nest
quotes and even mix and match that with other quotation marks if desired.</p>
<p>This is something that is extremely difficult to write in traditional shells
because it would require lots of escaping, and even escaping the escape
characters (and so on) the further deep you get in your nest.</p>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="multi-line-quotes" tabindex="-1"><a class="header-anchor" href="#multi-line-quotes" aria-hidden="true">#</a> Multi-Line Quotes</h3>
<p>Quotes can also work over multiple lines</p>
<pre><code>» out: %(foo
» bar)
foo
bar
</code></pre>
<h3 id="legacy-support" tabindex="-1"><a class="header-anchor" href="#legacy-support" aria-hidden="true">#</a> Legacy Support</h3>
<p>Version 3.x of Murex introduced support for the <code v-pre>%</code> token, before that brace
quotes worked without it. However to retain backwards compatibility, the older
syntax is still supported...albeit officially classed as &quot;deprecated&quot; and may
be removed from a future release.</p>
<p>Below is a little detail about how the legacy syntax worked:</p>
<h4 id="deprecated-syntax" tabindex="-1"><a class="header-anchor" href="#deprecated-syntax" aria-hidden="true">#</a> Deprecated Syntax</h4>
<p>The open brace character is only recognized as a brace quote token if it is the
start of a parameter.</p>
<pre><code>» set: example=(World!)
» out: (Hello $example)
Hello (World!)
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/array.html">Array (<code v-pre>@</code>) Token</RouterLink>:
Expand values as an array</li>
<li><RouterLink to="/user-guide/code-block.html">Code Block Parsing</RouterLink>:
Overview of how code blocks are parsed</li>
<li><RouterLink to="/parser/curly-brace.html">Curly Brace (<code v-pre>{</code>, <code v-pre>}</code>) Tokens</RouterLink>:
Initiates or terminates a code block</li>
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
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


