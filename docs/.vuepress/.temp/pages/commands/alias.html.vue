<template><div><h1 id="alias" tabindex="-1"><a class="header-anchor" href="#alias" aria-hidden="true">#</a> <code v-pre>alias</code></h1>
<blockquote>
<p>Create an alias for a command</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>alias</code> defines an alias for global usage</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>alias: alias=command parameter parameter

!alias: command
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Because aliases are parsed into an array of parameters, you cannot put the
entire alias within quotes. For example:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code># bad :(
» alias hw="out Hello, World!"
» hw
exec: "out\\ Hello,\\ World!": executable file not found in $PATH

# good :)
» alias hw=out "Hello, World!"
» hw
Hello, World!
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Notice how only the command <code v-pre>out &quot;Hello, World!&quot;</code> is quoted in <code v-pre>alias</code> the
same way you would have done if you'd run that command &quot;naked&quot; in the command
line? This is how <code v-pre>alias</code> expects it's parameters and where <code v-pre>alias</code> on Murex
differs from <code v-pre>alias</code> in POSIX shells.</p>
<p>In some ways this makes <code v-pre>alias</code> a little less flexible than it might
otherwise be. However the design of this is to keep <code v-pre>alias</code> focused on it's
core objective. For any more advanced requirements you can use a <code v-pre>function</code>
instead.</p>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="allowed-characters" tabindex="-1"><a class="header-anchor" href="#allowed-characters" aria-hidden="true">#</a> Allowed characters</h3>
<p>Alias names can only include alpha-numeric characters, hyphen and underscore.
The following regex is used to validate the <code v-pre>alias</code>'s parameters:
<code v-pre>^([-_a-zA-Z0-9]+)=(.*?)$</code></p>
<h3 id="undefining-an-alias" tabindex="-1"><a class="header-anchor" href="#undefining-an-alias" aria-hidden="true">#</a> Undefining an alias</h3>
<p>Like all other definable states in Murex, you can delete an alias with the
bang prefix:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» alias hw=out "Hello, World!"
» hw
Hello, World!

» !alias hw
» hw
exec: "hw": executable file not found in $PATH
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="order-of-preference" tabindex="-1"><a class="header-anchor" href="#order-of-preference" aria-hidden="true">#</a> Order of preference</h3>
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
<li><code v-pre>alias</code></li>
<li><code v-pre>!alias</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/commands/export.html"><code v-pre>export</code></RouterLink>:
Define an environmental variable and set it's value</li>
<li><RouterLink to="/commands/fexec.html"><code v-pre>fexec</code> </RouterLink>:
Execute a command or function, bypassing the usual order of precedence.</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
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
</ul>
</div></template>


