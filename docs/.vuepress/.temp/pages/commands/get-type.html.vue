<template><div><h1 id="get-type" tabindex="-1"><a class="header-anchor" href="#get-type" aria-hidden="true">#</a> <code v-pre>get-type</code></h1>
<blockquote>
<p>Returns the data-type of a variable or pipe</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>get-type</code> returns the Murex data-type of a variable or pipe without
reading the data from it.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>get-type: \$variable -&gt; `&lt;stdout&gt;`

get-type: stdin -&gt; `&lt;stdout&gt;`

get-type: pipe -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Get the data-type of a variable</p>
<pre><code>» set: json example={[1,2,3]}
» get-type: \$example
json
</code></pre>
<blockquote>
<p>Please note that you will need to escape the dollar sign. If you don't
the value of the variable will be passed to <code v-pre>get-type</code> rather than the
name.</p>
</blockquote>
<p>Get the data-type of a functions STDIN</p>
<pre><code>» function: example { get-type stdin }
» tout: json {[1,2,3]} -&gt; example
json
</code></pre>
<p>Get the data-type of a Murex named pipe</p>
<pre><code>» pipe: example
» tout: &lt;example&gt; json {[1,2,3]}
» get-type: example
» !pipe: example
json
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/reserved-vars.html">Reserved Variables</RouterLink>:
Special variables reserved by Murex</li>
<li><RouterLink to="/user-guide/scoping.html">Variable and Config Scoping</RouterLink>:
How scoping works within Murex</li>
<li><RouterLink to="/commands/debug.html"><code v-pre>debug</code></RouterLink>:
Debugging information</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/pipe.html"><code v-pre>pipe</code></RouterLink>:
Manage Murex named pipes</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
<li><RouterLink to="/commands/tout.html"><code v-pre>tout</code></RouterLink>:
Print a string to the STDOUT and set it's data-type</li>
</ul>
</div></template>


