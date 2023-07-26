<template><div><h1 id="what-s-new-in-murex-v3-1-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v3-1-change-log" aria-hidden="true">#</a> What's new in murex v3.1 - Change Log</h1>
<p>This release includes mostly bug fixes and new experimental features which are opt into. To enable all experimental features, set the environmental variable <code v-pre>MUREX_EXPERIMENTAL</code> to any value. Or you can enable specific features individually via `config</p>
<p>Breaking Changes:</p>
<ul>
<li>None</li>
</ul>
<p>Features:</p>
<ul>
<li>
<p><code v-pre>default</code> statement added to <code v-pre>switch</code> (this will deprecate <code v-pre>catch</code>, albeit <code v-pre>catch</code> will remain supported for compatibility)</p>
</li>
<li>
<p>Option to change start directory: <code v-pre>config: set shell start-directory $path</code> (where $path is the working directory you wish to start the shell in)</p>
</li>
<li>
<p><code v-pre>onCommandCompletion</code> event no spawns a pty. This allows capturing of stdout while still presenting a tty to the executing process</p>
</li>
<li>
<p>EXPERIMENTAL: preview files. This can be enabled via <code v-pre>config: set shell preview-enabled true</code></p>
</li>
<li>
<p>EXPERIMENTAL: preview images. This can be enabled via <code v-pre>config: set shell preview-images true</code></p>
</li>
<li>
<p>EXPERIMENTAL: tty buffer. This can be enabled via <code v-pre>config: set shell tty-buffer-enabled true</code></p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>r</code> has been tweaked to make it more user friendly</p>
</li>
<li>
<p>Long overdue SIGWINCH signal handler to reflow after terminal resized</p>
</li>
<li>
<p>Lots of improvements to documentation</p>
</li>
</ul>
<p>Non-User Facing / Maintenance Changes:</p>
<ul>
<li>none</li>
</ul>
<p>Bug Fixes:</p>
<ul>
<li>
<p>re-added <code v-pre>columns</code> data-type to fix <code v-pre>jobs</code> builtin</p>
</li>
<li>
<p>added bang alternatives to the glob exception list</p>
</li>
<li>
<p>added missing support for <code v-pre>&amp;&amp;</code> tokens in expressions</p>
</li>
<li>
<p><code v-pre>![]</code> can now produce empty maps</p>
</li>
<li>
<p>fixed regression bug in file pipe mutex</p>
</li>
<li>
<p>fixed missing bounds check in statement parser</p>
</li>
</ul>
<hr>
<p>Published: 12.02.2023 at 21:22</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><a href="/rosetta">Rosetta Stone</a>:
A tabulated list of Bashism's and their equivalent Murex syntax</li>
<li><RouterLink to="/commands/switch.html"><code v-pre>switch</code></RouterLink>:
Blocks of cascading conditionals</li>
</ul>
</div></template>


