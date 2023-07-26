<template><div><h1 id="openagent" tabindex="-1"><a class="header-anchor" href="#openagent" aria-hidden="true">#</a> <code v-pre>openagent</code></h1>
<blockquote>
<p>Creates a handler function for `open</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>openagent</code> creates and destroys handler functions for writing data to the
terminal when accessed via <code v-pre>open</code> and STDOUT is a TTY.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Display code block for an associated data-type:</p>
<pre><code>openagent get data-type
</code></pre>
<p>Define an <code v-pre>open</code> handler function:</p>
<pre><code>openagent set data-type { code-block }
</code></pre>
<p>Undefine an <code v-pre>open</code> handler:</p>
<pre><code>!openagent data-type
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="fileref" tabindex="-1"><a class="header-anchor" href="#fileref" aria-hidden="true">#</a> FileRef</h3>
<p>It is possible to track which shell script or module installed what <code v-pre>open</code>
handler by checking <code v-pre>runtime --open-agents</code> and checking it's <strong>FileRef</strong>.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>openagent</code></li>
<li><code v-pre>!openagent</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/fileref.html">FileRef</RouterLink>:
How to track what code was loaded and from where</li>
<li><RouterLink to="/user-guide/modules.html">Modules and Packages</RouterLink>:
An introduction to Murex modules and packages</li>
<li><RouterLink to="/commands/fexec.html"><code v-pre>fexec</code> </RouterLink>:
Execute a command or function, bypassing the usual order of precedence.</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


