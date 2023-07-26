<template><div><h1 id="private" tabindex="-1"><a class="header-anchor" href="#private" aria-hidden="true">#</a> <code v-pre>private</code></h1>
<blockquote>
<p>Define a private function block</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>private</code> defines a function who's scope is limited to that module or source
file.</p>
<p>Privates cannot be called from one module to another (unless they're wrapped
around a global <code v-pre>function</code>) and nor can they be called from the interactive
command line. The purpose of a <code v-pre>private</code> is to reduce repeated code inside
a module or source file without cluttering up the global namespace.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>private: name { code-block }
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code># The following cannot be entered via the command line. You need to write
# it to a file and execute it from there.

private hw {
    out &quot;Hello, World!&quot;
}

function tom {
    hw
    out &quot;My name is Tom.&quot;
}

function dick {
    hw
    out &quot;My name is Dick.&quot;
}

function harry {
    hw
    out &quot;My name is Harry.&quot;
}
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="allowed-characters" tabindex="-1"><a class="header-anchor" href="#allowed-characters" aria-hidden="true">#</a> Allowed characters</h3>
<p>Private names can only include any characters apart from dollar (<code v-pre>$</code>).
This is to prevent functions from overwriting variables (see the order of
preference below).</p>
<h3 id="undefining-a-private" tabindex="-1"><a class="header-anchor" href="#undefining-a-private" aria-hidden="true">#</a> Undefining a private</h3>
<p>Because private functions are fixed to the source file that declares them,
there isn't much point in undefining them. Thus at this point in time, it
is not possible to do so.</p>
<h3 id="order-of-preference" tabindex="-1"><a class="header-anchor" href="#order-of-preference" aria-hidden="true">#</a> Order of preference</h3>
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
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/alias.html"><code v-pre>alias</code></RouterLink>:
Create an alias for a command</li>
<li><RouterLink to="/commands/break.html"><code v-pre>break</code></RouterLink>:
Terminate execution of a block within your processes scope</li>
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
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
<li><RouterLink to="/commands/source.html"><code v-pre>source</code> </RouterLink>:
Import Murex code from another file of code block</li>
</ul>
</div></template>


