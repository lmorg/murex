<template><div><h1 id="curly-brace-tokens" tabindex="-1"><a class="header-anchor" href="#curly-brace-tokens" aria-hidden="true">#</a> Curly Brace (<code v-pre>{</code>, <code v-pre>}</code>) Tokens</h1>
<blockquote>
<p>Initiates or terminates a code block</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Curly braces are used to denote the start and end of a code block. Like with
the single quotation marks (<code v-pre>'</code>), any code inside a curly brace is not parsed.
Also unlike any other quotation tokens, the curly brace is included as part
of the parsed string.</p>
<pre><code>» out: {example}
{example}
</code></pre>
<p>Also like the brace quote (<code v-pre>(</code>, <code v-pre>)</code>), the curly brace character is only
recognized as a curly brace token if it is the start of a parameter.</p>
<p>Curly braces are also used for other fields besides code blocks. For example
inlining JSON.</p>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="multiline-blocks" tabindex="-1"><a class="header-anchor" href="#multiline-blocks" aria-hidden="true">#</a> Multiline Blocks</h3>
<p>Curly braces can work over multiple lines</p>
<pre><code>» out: {foo
» bar}
{foo
bar}
</code></pre>
<h3 id="code-golfing" tabindex="-1"><a class="header-anchor" href="#code-golfing" aria-hidden="true">#</a> Code Golfing</h3>
<p>Curly braces can be used to terminate the parsing of the command name / start
the parsing of the first parameter however each new parameter would still need
to be separated by whitespace:</p>
<pre><code># Valid
if{true} {out: &quot;Yipee&quot;}

# Invalid
if{true}{out: &quot;Yipee&quot;}
</code></pre>
<h3 id="ansi-constants" tabindex="-1"><a class="header-anchor" href="#ansi-constants" aria-hidden="true">#</a> ANSI Constants</h3>
<p>Some builtins (like <code v-pre>out</code>) also support infixing using the curly brace. eg</p>
<pre><code>out: &quot;{GREEN}PASSED{RESET}&quot;
</code></pre>
<p>This is a separate layer of parsing and happens at the parameter level for
specific builtins which opt to support ANSI constants. See the ANSI Constant
user guide (link below) for more information on supporting builtins and which
constants are available.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/ansi.html">ANSI Constants</RouterLink>:
Infixed constants that return ANSI escape sequences</li>
<li><RouterLink to="/parser/array.html">Array (<code v-pre>@</code>) Token</RouterLink>:
Expand values as an array</li>
<li><RouterLink to="/parser/brace-quote.html">Brace Quote (<code v-pre>%(</code>, <code v-pre>)</code>) Tokens</RouterLink>:
Initiates or terminates a string (variables expanded)</li>
<li><RouterLink to="/user-guide/code-block.html">Code Block Parsing</RouterLink>:
Overview of how code blocks are parsed</li>
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
<li><RouterLink to="/commands/err.html"><code v-pre>err</code></RouterLink>:
Print a line to the STDERR</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
<li><RouterLink to="/commands/tout.html"><code v-pre>tout</code></RouterLink>:
Print a string to the STDOUT and set it's data-type</li>
</ul>
</div></template>


