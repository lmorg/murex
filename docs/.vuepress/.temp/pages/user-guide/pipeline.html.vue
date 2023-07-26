<template><div><h1 id="pipeline" tabindex="-1"><a class="header-anchor" href="#pipeline" aria-hidden="true">#</a> Pipeline</h1>
<blockquote>
<p>Overview of what a &quot;pipeline&quot; is</p>
</blockquote>
<p>In the Murex docs you'll often see the term &quot;pipeline&quot;. This refers to any
commands sequenced together.</p>
<p>A pipeline can be joined via any pipe token (eg <code v-pre>|</code>, <code v-pre>-&gt;</code>, <code v-pre>=&gt;</code>, <code v-pre>?</code>). But,
for the sake of documentation, a pipeline might even be a solitary command.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Typical Murex pipeline:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>open: example.json -> [[ /node/0 ]]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Example of a single command pipeline:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>top
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Pipeline you might see in Bash / Zsh (this is also valid in Murex):</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>cat names.txt | sort | uniq
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Pipeline filtering out a specific error from <code v-pre>example-cmd</code></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>example-cmd ? grep: "File not found"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>A pipeline isn't a Murex specific construct but rather something inherited
from Unix. Where Murex differs is that it can support sending typed
information to compatible functions (unlike standard Unix pipes which are
dumb-byte streams).</p>
<p>Wikipedia has a page on <a href="https://en.wikipedia.org/wiki/Pipeline_(Unix)" target="_blank" rel="noopener noreferrer">Pipeline (Unix)<ExternalLinkIcon/></a>:</p>
<blockquote>
<p>In Unix-like computer operating systems, a pipeline is a mechanism for
inter-process communication using message passing. A pipeline is a set of
processes chained together by their standard streams, so that the output
text of each process (stdout) is passed directly as input (stdin) to the
next one. The second process is started as the first process is still
executing, and they are executed concurrently. The concept of pipelines was
championed by Douglas McIlroy at Unix's ancestral home of Bell Labs, during
the development of Unix, shaping its toolbox philosophy. It is named by
analogy to a physical pipeline. A key feature of these pipelines is their
&quot;hiding of internals&quot; (Ritchie &amp; Thompson, 1974). This in turn allows for
more clarity and simplicity in the system.</p>
</blockquote>
<h2 id="named-pipes" tabindex="-1"><a class="header-anchor" href="#named-pipes" aria-hidden="true">#</a> Named Pipes</h2>
<p>The drawback with pipes is that it assumes each command runs sequentially one
after another and that everything fits neatly into the concept of &quot;output&quot; and
&quot;errors&quot;. The moment you need to use background (<code v-pre>bg</code>) processes, do anything
more specific with data streams (even if just ignore them entirely), or use
more than one data stream, then this concept breaks down. This is where named
pipes come to the rescue. Named pipes are out of scope for this specific
document but you can read more on them in links the links below.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/pipe-arrow.html">Arrow Pipe (<code v-pre>-&gt;</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/user-guide/bang-prefix.html">Bang Prefix</RouterLink>:
Bang prefixing to reverse default actions</li>
<li><RouterLink to="/parser/pipe-generic.html">Generic Pipe (<code v-pre>=&gt;</code>) Token</RouterLink>:
Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/parser/pipe-posix.html">POSIX Pipe (<code v-pre>|</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/parser/pipe-err.html">STDERR Pipe (<code v-pre>?</code>) Token</RouterLink>:
Pipes STDERR from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/user-guide/schedulers.html">Schedulers</RouterLink>:
Overview of the different schedulers (or 'run modes') in Murex</li>
<li><RouterLink to="/commands/bg.html"><code v-pre>bg</code></RouterLink>:
Run processes in the background</li>
</ul>
</div></template>


