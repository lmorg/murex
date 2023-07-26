<template><div><h1 id="g-autoglob" tabindex="-1"><a class="header-anchor" href="#g-autoglob" aria-hidden="true">#</a> <code v-pre>@g</code> (autoglob)</h1>
<blockquote>
<p>Command prefix to expand globbing (deprecated)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><strong>This feature is now deprecated and only applies to murex version 2:</strong></p>
<p>By default Murex does not expand globbing (<code v-pre>*</code> and <code v-pre>?</code> wildcards) instead
encouraging the use of <code v-pre>g</code> (and similar) inside a subshell. While the aim of
this is to promote correctness, it can be a little annoying while working in
the interactive shell. For this reason you can prefix any command with <code v-pre>@g</code> to
enable Bash-like globbing.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>@g command ...
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>@g echo *
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>As of Murex <code v-pre>2.9</code> and above it is possible to enable automatic globbing in
the interactive shell without having to prefix the command with <code v-pre>@g</code> by
enabling the following <code v-pre>config</code> option:</p>
<pre><code>config: set shell auto-glob true
</code></pre>
<p>It is enabled by default on from version 3.x onwards (and renamed to
<code v-pre>expand-glob</code>)</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/profile.html">Murex Profile Files</RouterLink>:
A breakdown of the different files loaded on start up</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/f.html"><code v-pre>f</code></RouterLink>:
Lists or filters file system objects (eg files)</li>
<li><RouterLink to="/commands/g.html"><code v-pre>g</code></RouterLink>:
Glob pattern matching for file system objects (eg <code v-pre>*.txt</code>)</li>
<li><RouterLink to="/commands/rx.html"><code v-pre>rx</code></RouterLink>:
Regexp pattern matching for file system objects (eg <code v-pre>.*\\.txt</code>)</li>
</ul>
</div></template>


