<template><div><h1 id="esccli" tabindex="-1"><a class="header-anchor" href="#esccli" aria-hidden="true">#</a> <code v-pre>esccli</code></h1>
<blockquote>
<p>Escapes an array so output is valid shell code</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>esccli</code> takes an array and escapes any characters that might cause problems
when pasted back into the terminal. Typically you'd want to use this against
command parameters.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>&lt;stdin> -> esccli -> &lt;stdout>

esccli @array -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>As a method</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» alias foobar=out 'foo$b@r'
» alias -> [foobar]
[
    "out",
    "foo$b@r"
]
» alias -> [foobar] -> esccli
out foo\$b\@r
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>As a function</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» alias -> [foobar] -> set: fb
» $fb
["out","foo$b@r"]
» esccli: @fb
out foo\$b\@r
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/alias.html"><code v-pre>alias</code></RouterLink>:
Create an alias for a command</li>
<li><RouterLink to="/commands/escape.html"><code v-pre>escape</code></RouterLink>:
Escape or unescape input</li>
<li><RouterLink to="/commands/eschtml.html"><code v-pre>eschtml</code></RouterLink>:
Encode or decodes text for HTML</li>
<li><RouterLink to="/commands/escurl.html"><code v-pre>escurl</code></RouterLink>:
Encode or decodes text for the URL</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
</ul>
</div></template>


