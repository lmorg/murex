<template><div><h1 id="generic-pipe-token" tabindex="-1"><a class="header-anchor" href="#generic-pipe-token" aria-hidden="true">#</a> Generic Pipe (<code v-pre>=&gt;</code>) Token</h1>
<blockquote>
<p>Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This token behaves much like the <code v-pre>|</code> pipe would except it injects <code v-pre>format generic</code> into the pipeline. The purpose of a formatted pipe is to support
piping out to external commands which don't support Murex data types. For
example they might expect arrays as lists rather than JSON objects).</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» ja: [Mon..Wed] =&gt; cat
Mon
Tue
Wed
</code></pre>
<p>The above is literally the same as typing:</p>
<pre><code>» ja: [Mon..Wed] -&gt; format generic -&gt; cat
Mon
Tue
Wed
</code></pre>
<p>To demonstrate how the previous pipeline might look without a formatted pipe:</p>
<pre><code>» ja: [Mon..Wed] -&gt; cat
[&quot;Mon&quot;,&quot;Tue&quot;,&quot;Wed&quot;]

» ja: [Mon..Wed] | cat
[&quot;Mon&quot;,&quot;Tue&quot;,&quot;Wed&quot;]

» ja: [Mon..Wed]
[
    &quot;Mon&quot;,
    &quot;Tue&quot;,
    &quot;Wed&quot;
]
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/pipe-arrow.html">Arrow Pipe (<code v-pre>-&gt;</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/parser/pipe-posix.html">POSIX Pipe (<code v-pre>|</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/user-guide/pipeline.html">Pipeline</RouterLink>:
Overview of what a &quot;pipeline&quot; is</li>
<li><RouterLink to="/parser/pipe-err.html">STDERR Pipe (<code v-pre>?</code>) Token</RouterLink>:
Pipes STDERR from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/commands/namedpipe.html"><code v-pre>&lt;&gt;</code> / <code v-pre>read-named-pipe</code></RouterLink>:
Reads from a Murex named pipe</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
</ul>
</div></template>


