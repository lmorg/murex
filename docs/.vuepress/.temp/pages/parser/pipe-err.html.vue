<template><div><h1 id="stderr-pipe-token" tabindex="-1"><a class="header-anchor" href="#stderr-pipe-token" aria-hidden="true">#</a> STDERR Pipe (<code v-pre>?</code>) Token</h1>
<blockquote>
<p>Pipes STDERR from the left hand command to STDIN of the right hand command</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This token swaps the STDOUT and STDERR streams of the left hand command.</p>
<p>Please note that this token is only effective when it is prefixed by white
space.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» err Hello, world! ? regexp s/world/Earth/
Hello, Earth!
</code></pre>
<p>In following example the first command is writing to STDOUT rather than STDERR
so <code v-pre>Hello, world!</code> doesn't get pipelined and thus isn't affected by <code v-pre>regexp</code>:</p>
<pre><code>» out Hello, world! ? regexp s/world/Earth/
Hello, world!
</code></pre>
<p>In following example the STDERR token isn't whitespace padded so is treated
like any ordinary printable character:</p>
<pre><code>» err Hello, world!? regexp s/world/Earth/
Hello, world!? regexp s/world/Earth/
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/pipe-arrow.html">Arrow Pipe (<code v-pre>-&gt;</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/parser/pipe-generic.html">Generic Pipe (<code v-pre>=&gt;</code>) Token</RouterLink>:
Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/parser/pipe-posix.html">POSIX Pipe (<code v-pre>|</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/user-guide/pipeline.html">Pipeline</RouterLink>:
Overview of what a &quot;pipeline&quot; is</li>
<li><RouterLink to="/commands/namedpipe.html"><code v-pre>&lt;&gt;</code> / <code v-pre>read-named-pipe</code></RouterLink>:
Reads from a Murex named pipe</li>
<li><RouterLink to="/commands/err.html"><code v-pre>err</code></RouterLink>:
Print a line to the STDERR</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/regexp.html"><code v-pre>regexp</code></RouterLink>:
Regexp tools for arrays / lists of strings</li>
</ul>
</div></template>


