<template><div><h1 id="null" tabindex="-1"><a class="header-anchor" href="#null" aria-hidden="true">#</a> <code v-pre>null</code></h1>
<blockquote>
<p>null function. Similar to /dev/null</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>null</code> is a function that acts a little like the <code v-pre>null</code> data type and the
UNIX /dev/null device.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; null
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» out: &quot;Hello, world!&quot; -&gt; null
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>While this method does exist, a more idiomatic way to suppress STDOUT is to
use the named pipe property rather than piping to null:</p>
<pre><code>» out: &lt;null&gt; &quot;Hello, world!&quot;
</code></pre>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>null</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/break.html"><code v-pre>break</code></RouterLink>:
Terminate execution of a block within your processes scope</li>
<li><RouterLink to="/commands/die.html"><code v-pre>die</code></RouterLink>:
Terminate murex with an exit number of 1</li>
<li><RouterLink to="/commands/exit.html"><code v-pre>exit</code></RouterLink>:
Exit murex</li>
</ul>
</div></template>


