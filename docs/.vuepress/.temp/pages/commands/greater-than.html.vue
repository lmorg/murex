<template><div><h1 id="truncate-file" tabindex="-1"><a class="header-anchor" href="#truncate-file" aria-hidden="true">#</a> <code v-pre>&gt;</code> (truncate file)</h1>
<blockquote>
<p>Writes STDIN to disk - overwriting contents if file already exists</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Redirects output to file.</p>
<p>If a file already exists, the contents will be truncated (overwritten).
Otherwise a new file is created.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` |&gt; filename
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>g * |&gt; files.txt
</code></pre>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>&gt;</code></li>
<li><code v-pre>fwrite</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/pipe-arrow.html">Arrow Pipe (<code v-pre>-&gt;</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/parser/pipe-posix.html">POSIX Pipe (<code v-pre>|</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/parser/pipe-err.html">STDERR Pipe (<code v-pre>?</code>) Token</RouterLink>:
Pipes STDERR from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/commands/namedpipe.html"><code v-pre>&lt;&gt;</code> / <code v-pre>read-named-pipe</code></RouterLink>:
Reads from a Murex named pipe</li>
<li><RouterLink to="/commands/greater-than-greater-than.html"><code v-pre>&gt;&gt;</code> (append file)</RouterLink>:
Writes STDIN to disk - appending contents if file already exists</li>
<li><RouterLink to="/commands/g.html"><code v-pre>g</code></RouterLink>:
Glob pattern matching for file system objects (eg <code v-pre>*.txt</code>)</li>
<li><RouterLink to="/commands/pipe.html"><code v-pre>pipe</code></RouterLink>:
Manage Murex named pipes</li>
<li><RouterLink to="/commands/tmp.html"><code v-pre>tmp</code></RouterLink>:
Create a temporary file and write to it</li>
</ul>
</div></template>


