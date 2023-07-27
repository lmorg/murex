<template><div><h1 id="escape" tabindex="-1"><a class="header-anchor" href="#escape" aria-hidden="true">#</a> <code v-pre>escape</code></h1>
<blockquote>
<p>Escape or unescape input</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>escape</code> takes input from either STDIN or the parameters and returns the same
data, escaped.</p>
<p><code v-pre>!escape</code> does the same process in reverse, where it takes escaped data and
returns its unescaped counterpart.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Escape</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>&lt;stdin> -> escape -> &lt;stdout>

escape string to escape -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Unescape</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>&lt;stdin> -> !escape -> &lt;stdout>

!escape string to unescape -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Escape</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out (multi
» line
» string) -> escape
"multi\nline\nstring\n"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>escape</code></li>
<li><code v-pre>!escape</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/esccli.html"><code v-pre>esccli</code></RouterLink>:
Escapes an array so output is valid shell code</li>
<li><RouterLink to="/commands/eschtml.html"><code v-pre>eschtml</code></RouterLink>:
Encode or decodes text for HTML</li>
<li><RouterLink to="/commands/escurl.html"><code v-pre>escurl</code></RouterLink>:
Encode or decodes text for the URL</li>
</ul>
</div></template>


