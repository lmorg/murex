<template><div><h1 id="expr" tabindex="-1"><a class="header-anchor" href="#expr" aria-hidden="true">#</a> <code v-pre>expr</code></h1>
<blockquote>
<p>Expressions: mathematical, string comparisons, logical operators</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>expr: expression -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Order of operations:</p>
<pre><code>» 3 * (3 + 1)
12
</code></pre>
<p>JSON array:</p>
<pre><code>» %[apples oranges grapes]
[
    &quot;apples&quot;,
    &quot;oranges&quot;,
    &quot;grapes&quot;
]
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/equ.html"><code v-pre>=</code> (arithmetic evaluation)</RouterLink>:
Evaluate a mathematical function (deprecated)</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


