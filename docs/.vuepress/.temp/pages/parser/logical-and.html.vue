<template><div><h1 id="and-logical-operator" tabindex="-1"><a class="header-anchor" href="#and-logical-operator" aria-hidden="true">#</a> And (<code v-pre>&amp;&amp;</code>) Logical Operator</h1>
<blockquote>
<p>Continues next operation if previous operation passes</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>When in the <strong>normal</strong> run mode (see &quot;schedulers&quot; link below) this will only
run the command on the right hand side if the command on the left hand side
does not error. Neither STDOUT nor STDERR are piped.</p>
<p>This has no effect in <code v-pre>try</code> nor <code v-pre>trypipe</code> run modes because they automatically
apply stricter error handling.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Second command runs because the first command doesn't error:</p>
<pre><code>» out: one &amp;&amp; out: two
one
two
</code></pre>
<p>Second command does not run because the first command produces an error:</p>
<pre><code>» err: one &amp;&amp; out: two
one
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>This is equivelent to a <code v-pre>try</code> block:</p>
<pre><code>try {
    err: one
    out: two
}
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/pipeline.html">Pipeline</RouterLink>:
Overview of what a &quot;pipeline&quot; is</li>
<li><RouterLink to="/parser/pipe-err.html">STDERR Pipe (<code v-pre>?</code>) Token</RouterLink>:
Pipes STDERR from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/user-guide/schedulers.html">Schedulers</RouterLink>:
Overview of the different schedulers (or 'run modes') in Murex</li>
<li><RouterLink to="/commands/err.html"><code v-pre>err</code></RouterLink>:
Print a line to the STDERR</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/try.html"><code v-pre>try</code></RouterLink>:
Handles errors inside a block of code</li>
<li><RouterLink to="/commands/trypipe.html"><code v-pre>trypipe</code></RouterLink>:
Checks state of each function in a pipeline and exits block on error</li>
</ul>
</div></template>


