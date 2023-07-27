<template><div><h1 id="what-s-new-in-murex-v2-0-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-0-change-log" aria-hidden="true">#</a> What's new in murex v2.0 - Change Log</h1>
<p>This release comes with spellchecking, inlined images, smarter syntax completion and more</p>
<p>This release seems some major upgrade to the default experience:</p>
<ul>
<li>
<p>Inlined spellchecking enabled by default (requires <code v-pre>aspell</code> installed)</p>
<pre><code>[![asciicast](https://asciinema.org/a/408024.svg)](https://asciinema.org/a/408024)

(see spellcheck link below)
</code></pre>
</li>
<li>
<p>Support for inlining images enabled by default (was previously an <a href="https://github.com/lmorg/murex-module-open-image" target="_blank" rel="noopener noreferrer">external
module<ExternalLinkIcon/></a>)</p>
<pre><code>[![asciicast](https://asciinema.org/a/408028.svg)](https://asciinema.org/a/408028)
</code></pre>
</li>
<li>
<p>Massively overhauled syntax completion</p>
<p><a href="https://asciinema.org/a/408029" target="_blank" rel="noopener noreferrer"><img src="https://asciinema.org/a/408029.svg" alt="asciicast" loading="lazy"><ExternalLinkIcon/></a></p>
</li>
<li>
<p>Additional tab autocompletions included in base install</p>
</li>
<li>
<p>Minor <code v-pre>readline</code> bug fixes: <a href="https://github.com/lmorg/murex/pull/312/commits/5064cf418f768d2ba4a6bbc7c74e46629ef3b5f3" target="_blank" rel="noopener noreferrer">git commits<ExternalLinkIcon/></a></p>
</li>
<li>
<p>Rewritten config backend which copies less data around</p>
</li>
<li>
<p>Significantly more features documented: <a href="https://murex.rocks" target="_blank" rel="noopener noreferrer">https://murex.rocks<ExternalLinkIcon/></a></p>
</li>
<li>
<p>Plus this new blog!</p>
</li>
</ul>
<hr>
<p>Published: 17.04.2021 at 12:49</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/modules.html">Modules and Packages</RouterLink>:
An introduction to Murex modules and packages</li>
<li><RouterLink to="/user-guide/interactive-shell.html">Murex's Interactive Shell</RouterLink>:
What's different about Murex's interactive shell?</li>
<li><RouterLink to="/user-guide/spellcheck.html">Spellcheck</RouterLink>:
How to enable inline spellchecking</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/murex-package.html"><code v-pre>murex-package</code></RouterLink>:
Murex's package manager</li>
<li><RouterLink to="/commands/open-image.html"><code v-pre>open-image</code> </RouterLink>:
Renders bitmap image data on your terminal</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
</ul>
</div></template>


