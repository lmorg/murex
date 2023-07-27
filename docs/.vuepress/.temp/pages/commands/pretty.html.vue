<template><div><h1 id="pretty" tabindex="-1"><a class="header-anchor" href="#pretty" aria-hidden="true">#</a> <code v-pre>pretty</code></h1>
<blockquote>
<p>Prettifies JSON to make it human readable</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Takes JSON from the STDIN and reformats it to make it human readable, then
outputs that to STDOUT.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; pretty -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» tout: json {&quot;Array&quot;:[1,2,3],&quot;Map&quot;:{&quot;String&quot;: &quot;Foobar&quot;,&quot;Number&quot;:123.456}} -&gt; pretty
{
    &quot;Array&quot;: [
        1,
        2,
        3
    ],
    &quot;Map&quot;: {
        &quot;String&quot;: &quot;Foobar&quot;,
        &quot;Number&quot;: 123.456
    }
}
</code></pre>
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


