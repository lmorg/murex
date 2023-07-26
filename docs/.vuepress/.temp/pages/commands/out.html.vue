<template><div><h1 id="out" tabindex="-1"><a class="header-anchor" href="#out" aria-hidden="true">#</a> <code v-pre>out</code></h1>
<blockquote>
<p>Print a string to the STDOUT with a trailing new line character</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Write parameters to STDOUT with a trailing new line character.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>out: string to write -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» out Hello, World!
Hello, World!
</code></pre>
<p>For compatibility with other shells, <code v-pre>echo</code> is also supported:</p>
<pre><code>» echo Hello, World!
Hello, World!
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>out</code> / <code v-pre>echo</code> output as <code v-pre>string</code> data-type. This can be changed by casting
(<code v-pre>cast</code>) or using the <code v-pre>tout</code> function.</p>
<h3 id="ansi-constants" tabindex="-1"><a class="header-anchor" href="#ansi-constants" aria-hidden="true">#</a> ANSI Constants</h3>
<p><code v-pre>out</code> supports ANSI constants.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>out</code></li>
<li><code v-pre>echo</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/ansi.html">ANSI Constants</RouterLink>:
Infixed constants that return ANSI escape sequences</li>
<li><RouterLink to="/commands/brace-quote.html"><code v-pre>(</code> (brace quote)</RouterLink>:
Write a string to the STDOUT without new line</li>
<li><RouterLink to="/commands/greater-than-greater-than.html"><code v-pre>&gt;&gt;</code> (append file)</RouterLink>:
Writes STDIN to disk - appending contents if file already exists</li>
<li><RouterLink to="/commands/greater-than.html"><code v-pre>&gt;</code> (truncate file)</RouterLink>:
Writes STDIN to disk - overwriting contents if file already exists</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/err.html"><code v-pre>err</code></RouterLink>:
Print a line to the STDERR</li>
<li><RouterLink to="/commands/pt.html"><code v-pre>pt</code></RouterLink>:
Pipe telemetry. Writes data-types and bytes written</li>
<li><RouterLink to="/commands/read.html"><code v-pre>read</code></RouterLink>:
<code v-pre>read</code> a line of input from the user and store as a variable</li>
<li><RouterLink to="/commands/tout.html"><code v-pre>tout</code></RouterLink>:
Print a string to the STDOUT and set it's data-type</li>
<li><RouterLink to="/commands/tread.html"><code v-pre>tread</code></RouterLink>:
<code v-pre>read</code> a line of input from the user and store as a user defined <em>typed</em> variable (deprecated)</li>
</ul>
</div></template>


