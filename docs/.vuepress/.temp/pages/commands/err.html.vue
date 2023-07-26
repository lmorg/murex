<template><div><h1 id="err" tabindex="-1"><a class="header-anchor" href="#err" aria-hidden="true">#</a> <code v-pre>err</code></h1>
<blockquote>
<p>Print a line to the STDERR</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Write parameters to STDERR with a trailing new line character.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>err: string to write -&gt; &lt;stderr&gt;
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» err Hello, World!
Hello, World!
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>err</code> outputs as <code v-pre>string</code> data-type. This can be changed by casting</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>err { "Code": 404, "Message": "Page not found" } ? cast json
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>However passing structured data-types along the STDERR stream is not recommended
as any other function within your code might also pass error messages along the
same stream and thus taint your structured data. This is why Murex does not
supply a <code v-pre>tout</code> function for STDERR. The recommended solution for passing
messages like these which you want separate from your STDOUT stream is to create
a new Murex named pipe.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» pipe: --create messages
» bg { &lt;messages> -> pretty }
» tout: &lt;messages> json { "Code": 404, "Message": "Page not found" }
» pipe: --close messages
{
    "Code": 404,
    "Message": "Page not found"
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="ansi-constants" tabindex="-1"><a class="header-anchor" href="#ansi-constants" aria-hidden="true">#</a> ANSI Constants</h3>
<p><code v-pre>err</code> supports ANSI constants.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/ansi.html">ANSI Constants</RouterLink>:
Infixed constants that return ANSI escape sequences</li>
<li><RouterLink to="/commands/brace-quote.html"><code v-pre>(</code> (brace quote)</RouterLink>:
Write a string to the STDOUT without new line</li>
<li><RouterLink to="/commands/namedpipe.html"><code v-pre>&lt;&gt;</code> / <code v-pre>read-named-pipe</code></RouterLink>:
Reads from a Murex named pipe</li>
<li><RouterLink to="/commands/greater-than-greater-than.html"><code v-pre>&gt;&gt;</code> (append file)</RouterLink>:
Writes STDIN to disk - appending contents if file already exists</li>
<li><RouterLink to="/commands/greater-than.html"><code v-pre>&gt;</code> (truncate file)</RouterLink>:
Writes STDIN to disk - overwriting contents if file already exists</li>
<li><RouterLink to="/commands/bg.html"><code v-pre>bg</code></RouterLink>:
Run processes in the background</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/pipe.html"><code v-pre>pipe</code></RouterLink>:
Manage Murex named pipes</li>
<li><RouterLink to="/commands/pretty.html"><code v-pre>pretty</code></RouterLink>:
Prettifies JSON to make it human readable</li>
<li><RouterLink to="/commands/pt.html"><code v-pre>pt</code></RouterLink>:
Pipe telemetry. Writes data-types and bytes written</li>
<li><RouterLink to="/commands/tout.html"><code v-pre>tout</code></RouterLink>:
Print a string to the STDOUT and set it's data-type</li>
</ul>
</div></template>


