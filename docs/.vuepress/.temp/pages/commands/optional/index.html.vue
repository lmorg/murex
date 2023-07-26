<template><div><h1 id="optional-command-reference" tabindex="-1"><a class="header-anchor" href="#optional-command-reference" aria-hidden="true">#</a> Optional Command Reference</h1>
<p>This section is a glossary of Murex optional builtins.</p>
<p>These builtins likely wont be compiled with Murex unless you specifically
request them to be. This might be because they duplicate functionality
already available on POSIX systems or introduce more complex dependencies.
It might also be because that specific builtin is in an alpha stage and thus
not ready to ship with Murex.</p>
<h2 id="other-reference-material" tabindex="-1"><a class="header-anchor" href="#other-reference-material" aria-hidden="true">#</a> Other Reference Material</h2>
<h3 id="language-guides" tabindex="-1"><a class="header-anchor" href="#language-guides" aria-hidden="true">#</a> Language Guides</h3>
<ol>
<li>
<p><a href="./commands">Builtins</a>, for docs
on the core builtins.</p>
</li>
<li>
<p><a href="/tour">Language Tour</a>, which is an introduction into
the Murex language.</p>
</li>
<li>
<p><a href="/rosetta">Rosetta Stone</a>, which is a reference
table comparing Bash syntax to Murex's.</p>
</li>
</ol>
<h3 id="murex-s-source-code" tabindex="-1"><a class="header-anchor" href="#murex-s-source-code" aria-hidden="true">#</a> Murex's Source Code</h3>
<p>The source for each of these builtins can be found on <a href="https://github.com/lmorg/murex/tree/master/builtins/optional" target="_blank" rel="noopener noreferrer">Github<ExternalLinkIcon/></a>.</p>
<h3 id="shell-commands-for-querying-builtins" tabindex="-1"><a class="header-anchor" href="#shell-commands-for-querying-builtins" aria-hidden="true">#</a> Shell Commands For Querying Builtins</h3>
<p>From the shell itself: run <code v-pre>builtins</code> to list the builtin command.</p>
<p>If you require a manual on any of those commands, you can run <code v-pre>murex-docs</code>
to return the same markdown-formatted document as those listed below. eg</p>
<pre><code>murex-docs trypipe
</code></pre>
<h2 id="pages" tabindex="-1"><a class="header-anchor" href="#pages" aria-hidden="true">#</a> Pages</h2>
<ul>
<li><RouterLink to="/commands/optional/bz2.html"><code v-pre>!bz2</code> </RouterLink>:
Decompress a bz2 file</li>
<li><RouterLink to="/commands/optional/base64.html"><code v-pre>base64</code> </RouterLink>:
Encode or decode a base64 string</li>
<li><RouterLink to="/commands/optional/gz.html"><code v-pre>gz</code> </RouterLink>:
Compress or decompress a gzip file</li>
<li><RouterLink to="/commands/optional/qr.html"><code v-pre>qr</code> </RouterLink>:
Creates a QR code from STDIN</li>
<li><RouterLink to="/commands/optional/select.html"><code v-pre>select</code> </RouterLink>:
Inlining SQL into shell pipelines</li>
<li><RouterLink to="/commands/optional/sleep.html"><code v-pre>sleep</code> </RouterLink>:
Suspends the shell for a number of seconds</li>
</ul>
</div></template>


