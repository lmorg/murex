<template><div><h1 id="man-summary" tabindex="-1"><a class="header-anchor" href="#man-summary" aria-hidden="true">#</a> <code v-pre>man-summary</code></h1>
<blockquote>
<p>Outputs a man page summary of a command</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>man-summary</code> reads the man pages for a given command and outputs it's
summary (if one exists).</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>man-summary command -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» man-summary: man
man - an interface to the on-line reference manuals
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/man-get-flags.html"><code v-pre>man-get-flags</code> </RouterLink>:
Parses man page files for command line flags</li>
<li><RouterLink to="/commands/murex-docs.html"><code v-pre>murex-docs</code></RouterLink>:
Displays the man pages for Murex builtins</li>
<li><RouterLink to="/commands/summary.html"><code v-pre>summary</code> </RouterLink>:
Defines a summary help text for a command</li>
</ul>
</div></template>


