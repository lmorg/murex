<template><div><h1 id="pt" tabindex="-1"><a class="header-anchor" href="#pt" aria-hidden="true">#</a> <code v-pre>pt</code></h1>
<blockquote>
<p>Pipe telemetry. Writes data-types and bytes written</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Pipe telemetry, <code v-pre>pt</code>, writes statistics about the pipeline. The telemetry is written
directly to the OS's STDERR so to preserved the pipeline.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; pt -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>curl -s https://example.com/bigfile.bin -&gt; pt -&gt; &gt; bigfile.bin
</code></pre>
<p>(though Murex does also have it's own HTTP clients, <code v-pre>get</code>, <code v-pre>post</code> and
<code v-pre>getfile</code>)</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/greater-than-greater-than.html"><code v-pre>&gt;&gt;</code> (append file)</RouterLink>:
Writes STDIN to disk - appending contents if file already exists</li>
<li><RouterLink to="/commands/greater-than.html"><code v-pre>&gt;</code> (truncate file)</RouterLink>:
Writes STDIN to disk - overwriting contents if file already exists</li>
<li><RouterLink to="/commands/get.html"><code v-pre>get</code></RouterLink>:
Makes a standard HTTP request and returns the result as a JSON object</li>
<li><RouterLink to="/commands/getfile.html"><code v-pre>getfile</code></RouterLink>:
Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.</li>
<li><RouterLink to="/commands/post.html"><code v-pre>post</code></RouterLink>:
HTTP POST request with a JSON-parsable return</li>
</ul>
</div></template>


