<template><div><h1 id="create-object-constructor" tabindex="-1"><a class="header-anchor" href="#create-object-constructor" aria-hidden="true">#</a> Create object (<code v-pre>%{}</code>) constructor</h1>
<blockquote>
<p>Quickly generate objects and maps</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>%{}</code> is a way of defining objects in expressions and statements. Whenever an
<code v-pre>%{}</code> object is outputted as a string, it will be converted to minified JSON.</p>
<p>Object elements inside <code v-pre>%{}</code> can be new line and/or comma delimited. This
allows for compatibility with JSON but without the pain of accidentally invalid
comma management breaking JSON parsers. However a colon is still required to
separate keys from values.</p>
<p>Like with YAML, strings in <code v-pre>%[]</code> do not need to be quoted unless you need to
force numeric or boolean looking values to be stored as strings.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p><strong>Object passed as a JSON string:</strong></p>
<pre><code>» echo %{foo: bar}
{&quot;foo&quot;:&quot;bar&quot;}
</code></pre>
<p><strong>The <code v-pre>%</code> prefix for the nested object is optional:</strong></p>
<pre><code>» %{foo: bar, baz: [1 2 3]}
{
    &quot;baz&quot;: [
        1,
        2,
        3
    ],
    &quot;foo&quot;: &quot;bar&quot;
}
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/brace-quote.html">Brace Quote (<code v-pre>%(</code>, <code v-pre>)</code>) Tokens</RouterLink>:
Initiates or terminates a string (variables expanded)</li>
<li><RouterLink to="/parser/create-array.html">Create array (<code v-pre>%[]</code>) constructor</RouterLink>:
Quickly generate arrays</li>
<li><RouterLink to="/parser/double-quote.html">Double Quote (<code v-pre>&quot;</code>) Token</RouterLink>:
Initiates or terminates a string (variables expanded)</li>
<li><RouterLink to="/parser/single-quote.html">Single Quote (<code v-pre>'</code>) Token</RouterLink>:
Initiates or terminates a string (variables not expanded)</li>
<li><RouterLink to="/mkarray/special.html">Special Ranges</RouterLink>:
Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)</li>
<li><RouterLink to="/commands/expr.html"><code v-pre>expr</code></RouterLink>:
Expressions: mathematical, string comparisons, logical operators</li>
</ul>
</div></template>


