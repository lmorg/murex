<template><div><h1 id="cast" tabindex="-1"><a class="header-anchor" href="#cast" aria-hidden="true">#</a> <code v-pre>cast</code></h1>
<blockquote>
<p>Alters the data type of the previous function without altering it's output</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>cast</code> works a little like when you case variables in lower level languages
where the value of the variable is unchanged. In Murex the contents in
the pipeline are preserved however the reported data type is altered.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>&lt;stdin&gt; -&gt; cast data-type -&gt; &lt;stdout&gt;
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out: {"Array":[1,2,3],"Map":{"String": "Foobar","Number":123.456}} -> cast json
{"Array":[1,2,3],"Map":{"String": "Foobar","Number":123.456}}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>If you want to reformat the STDIN into the new data type then use <code v-pre>format</code>
instead.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/tout.html"><code v-pre>tout</code></RouterLink>:
Print a string to the STDOUT and set it's data-type</li>
</ul>
</div></template>


