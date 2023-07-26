<template><div><h1 id="open-image" tabindex="-1"><a class="header-anchor" href="#open-image" aria-hidden="true">#</a> <code v-pre>open-image</code></h1>
<blockquote>
<p>Renders bitmap image data on your terminal</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>open-image</code> is an optional builtin which will render images (JPEG, GIF,
PNG, BMP, TIFF and WebP) to the terminal using block characters and ANSI
colour sequences.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; open-image -&gt; `&lt;stdout&gt;`

open-image file-path -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>As a method</p>
<pre><code>» cat example.png -&gt; open-image
</code></pre>
<p>As a function</p>
<pre><code>» open-image example.png
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>open-image</code> will fail if STDOUT is not a TTY.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/optional/qr.html"><code v-pre>qr</code> </RouterLink>:
Creates a QR code from STDIN</li>
</ul>
</div></template>


