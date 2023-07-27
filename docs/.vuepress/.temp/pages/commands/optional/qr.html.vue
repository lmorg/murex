<template><div><h1 id="qr-optional-command-reference" tabindex="-1"><a class="header-anchor" href="#qr-optional-command-reference" aria-hidden="true">#</a> <code v-pre>qr</code> - Optional Command Reference</h1>
<blockquote>
<p>Creates a QR code from STDIN</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>qr</code> is an optional builtin which generates a PNG format image based on the
input from STDIN. <code v-pre>qr</code> must be run as a method.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; qr -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Write the PNG to disk</p>
<pre><code>» out &quot;Hello, World!&quot; -&gt; qr -&gt; &gt; qr.png
</code></pre>
<p>Display PNG in the terminal</p>
<pre><code>» out &quot;Hello, World!&quot; -&gt; qr -&gt; open-image
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>qr</code> sets stdout's data-type to be &quot;image&quot;, which is defined in with the
<code v-pre>open-image</code> optional builtin. So if you have that disabled then you may
have to <code v-pre>cast</code> the output in some circumstances.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/open-image.html"><code v-pre>open-image</code> </RouterLink>:
Renders bitmap image data on your terminal</li>
</ul>
</div></template>


