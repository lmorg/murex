<template><div><h1 id="history" tabindex="-1"><a class="header-anchor" href="#history" aria-hidden="true">#</a> <code v-pre>history</code></h1>
<blockquote>
<p>Outputs murex's command history</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Outputs <em>mutex</em>'s command history.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>history -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» history
...
{
    &quot;Index&quot;: 16782,
    &quot;DateTime&quot;: &quot;2019-01-19T22:43:21.124273664Z&quot;,
    &quot;Block&quot;: &quot;tout: json ([\&quot;a\&quot;, \&quot;b\&quot;, \&quot;c\&quot;]) -\u003e len&quot;
},
{
    &quot;Index&quot;: 16783,
    &quot;DateTime&quot;: &quot;2019-01-19T22:50:42.114986768Z&quot;,
    &quot;Block&quot;: &quot;clear&quot;
},
{
    &quot;Index&quot;: 16784,
    &quot;DateTime&quot;: &quot;2019-01-19T22:51:39.82077789Z&quot;,
    &quot;Block&quot;: &quot;map { tout: json ([\&quot;key 1\&quot;, \&quot;key 2\&quot;, \&quot;key 3\&quot;]) }&quot;
},
...
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>The history file is typically located on disk in a file called <code v-pre>~/.murex.history</code>.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


