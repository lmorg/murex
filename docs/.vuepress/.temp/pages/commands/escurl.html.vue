<template><div><h1 id="escurl" tabindex="-1"><a class="header-anchor" href="#escurl" aria-hidden="true">#</a> <code v-pre>escurl</code></h1>
<blockquote>
<p>Encode or decodes text for the URL</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>escurl</code> takes input from either STDIN or the parameters and returns the same
data, escaped for the URL.</p>
<p><code v-pre>!eschtml</code> does the same process in reverse, where it takes URL escaped data
and returns its unescaped counterpart.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Escape</p>
<pre><code>`&lt;stdin&gt;` -&gt; escurl -&gt; `&lt;stdout&gt;`

escurl string to escape -&gt; `&lt;stdout&gt;`
</code></pre>
<p>Unescape</p>
<pre><code>`&lt;stdin&gt;` -&gt; !escurl -&gt; `&lt;stdout&gt;`

!escurl string to unescape -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Escape</p>
<pre><code>» out: &quot;!? &lt;&gt;&quot; -&gt; escurl
%21%3F%20%3C%3E%0A
</code></pre>
<p>Unescape</p>
<pre><code>out: '%21%3F%20%3C%3E%0A' -&gt; !escurl
!? &lt;&gt;
</code></pre>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>escurl</code></li>
<li><code v-pre>!escurl</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/escape.html"><code v-pre>escape</code></RouterLink>:
Escape or unescape input</li>
<li><RouterLink to="/commands/esccli.html"><code v-pre>esccli</code></RouterLink>:
Escapes an array so output is valid shell code</li>
<li><RouterLink to="/commands/eschtml.html"><code v-pre>eschtml</code></RouterLink>:
Encode or decodes text for HTML</li>
<li><RouterLink to="/commands/get.html"><code v-pre>get</code></RouterLink>:
Makes a standard HTTP request and returns the result as a JSON object</li>
<li><RouterLink to="/commands/getfile.html"><code v-pre>getfile</code></RouterLink>:
Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.</li>
<li><RouterLink to="/commands/post.html"><code v-pre>post</code></RouterLink>:
HTTP POST request with a JSON-parsable return</li>
</ul>
</div></template>


