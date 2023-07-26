<template><div><h1 id="fexec" tabindex="-1"><a class="header-anchor" href="#fexec" aria-hidden="true">#</a> <code v-pre>fexec</code></h1>
<blockquote>
<p>Execute a command or function, bypassing the usual order of precedence.</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>fexec</code> allows you to execute a command or function, bypassing the usual order
of precedence.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>fexec: flag command [ parameters... ] -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>fexec: private /source/builtin/autocomplete.alias
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>builtin</code>
Execute a Murex builtin</li>
<li><code v-pre>function</code>
Execute a Murex public function</li>
<li><code v-pre>help</code>
Display help message</li>
<li><code v-pre>private</code>
Execute a Murex private function</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
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
<h3 id="compatibility-with-posix" tabindex="-1"><a class="header-anchor" href="#compatibility-with-posix" aria-hidden="true">#</a> Compatibility with POSIX</h3>
<p>For compatibility with traditional shells like Bash and Zsh, <code v-pre>builtin</code> is an
alias to `fexec builtin</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>fexec</code></li>
<li><code v-pre>builtin</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/alias.html"><code v-pre>alias</code></RouterLink>:
Create an alias for a command</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/bg.html"><code v-pre>bg</code></RouterLink>:
Run processes in the background</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>builtins</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/event.html"><code v-pre>event</code></RouterLink>:
Event driven programming for shell scripts</li>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/commands/fg.html"><code v-pre>fg</code></RouterLink>:
Sends a background process into the foreground</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>jobs</code></RouterLink>:
Lists all running functions within the current Murex session</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/source.html"><code v-pre>source</code> </RouterLink>:
Import Murex code from another file of code block</li>
</ul>
</div></template>


