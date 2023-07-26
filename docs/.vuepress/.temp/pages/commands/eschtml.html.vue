<template><div><h1 id="eschtml" tabindex="-1"><a class="header-anchor" href="#eschtml" aria-hidden="true">#</a> <code v-pre>eschtml</code></h1>
<blockquote>
<p>Encode or decodes text for HTML</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>eschtml</code> takes input from either STDIN or the parameters and returns the same
data, HTML escaped.</p>
<p><code v-pre>!eschtml</code> does the same process in reverse, where it takes HTML escaped data
and returns its unescaped counterpart.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Escape</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>&lt;stdin> -> eschtml -> &lt;stdout>

eschtml string to escape -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Unescape</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>&lt;stdin> -> !eschtml -> &lt;stdout>

!eschtml string to unescape -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Escape</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out: "&lt;h1>foo &amp; bar&lt;/h1>" -> eschtml
&amp;lt;h1&amp;gt;foo &amp;amp; bar&amp;lt;/h1&amp;gt;
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>Unescape</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out: '&amp;lt;h1&amp;gt;foo &amp;amp; bar&amp;lt;/h1&amp;gt;' -> !eschtml
&lt;h1>foo &amp; bar&lt;/h1>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>eschtml</code></li>
<li><code v-pre>!eschtml</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/escape.html"><code v-pre>escape</code></RouterLink>:
Escape or unescape input</li>
<li><RouterLink to="/commands/esccli.html"><code v-pre>esccli</code></RouterLink>:
Escapes an array so output is valid shell code</li>
<li><RouterLink to="/commands/escurl.html"><code v-pre>escurl</code></RouterLink>:
Encode or decodes text for the URL</li>
<li><RouterLink to="/commands/get.html"><code v-pre>get</code></RouterLink>:
Makes a standard HTTP request and returns the result as a JSON object</li>
<li><RouterLink to="/commands/getfile.html"><code v-pre>getfile</code></RouterLink>:
Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.</li>
<li><RouterLink to="/commands/post.html"><code v-pre>post</code></RouterLink>:
HTTP POST request with a JSON-parsable return</li>
</ul>
</div></template>


