<template><div><h1 id="base64" tabindex="-1"><a class="header-anchor" href="#base64" aria-hidden="true">#</a> <code v-pre>base64</code></h1>
<blockquote>
<p>Encode or decode a base64 string</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>An optional builtin to encode or decode a base64 string.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>&lt;stdin> -> base64 -> &lt;stdout>

&lt;stdin> -> !base64 -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Encode base64 string</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out: "Hello, World!" -> base64
SGVsbG8sIFdvcmxkIQo=
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>Decode base64 string</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out: "SGVsbG8sIFdvcmxkIQo=" -> !base64
Hello, World!
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>base64</code> is very simplistic - particularly when compared to its GNU coreutil
(for example) counterpart. If you want to use the <code v-pre>base64</code> binary on Linux
or similar platforms then you will need to launch with the <code v-pre>exec</code> builtin:</p>
<pre><code>» out: &quot;Hello, World!&quot; -&gt; exec: base64
SGVsbG8sIFdvcmxkIQo=

» out: &quot;SGVsbG8sIFdvcmxkIQo=&quot; -&gt; exec: base64 -d
Hello, World!
</code></pre>
<p>However for simple tasks this builtin will out perform external tools because
it doesn't require the OS fork processes.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>base64</code></li>
<li><code v-pre>!base64</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/optional/bz2.html"><code v-pre>!bz2</code> </RouterLink>:
Decompress a bz2 file</li>
<li><RouterLink to="/commands/escape.html"><code v-pre>escape</code></RouterLink>:
Escape or unescape input</li>
<li><RouterLink to="/commands/esccli.html"><code v-pre>esccli</code></RouterLink>:
Escapes an array so output is valid shell code</li>
<li><RouterLink to="/commands/eschtml.html"><code v-pre>eschtml</code></RouterLink>:
Encode or decodes text for HTML</li>
<li><RouterLink to="/commands/escurl.html"><code v-pre>escurl</code></RouterLink>:
Encode or decodes text for the URL</li>
<li><RouterLink to="/optional/gz.html"><code v-pre>gz</code> </RouterLink>:
Compress or decompress a gzip file</li>
</ul>
</div></template>


