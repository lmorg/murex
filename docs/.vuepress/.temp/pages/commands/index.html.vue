<template><div><h1 id="command-reference" tabindex="-1"><a class="header-anchor" href="#command-reference" aria-hidden="true">#</a> Command Reference</h1>
<p>This section is a glossary of Murex builtin commands.</p>
<p>Because Murex is loosely modelled on the functional paradigm, it means
all language constructs are exposed via functions and those are typically
builtins because they can share the Murex runtime virtual machine.
However any executable command can also be called from within Murex;
be that either via the <code v-pre>exec</code> builtin or natively like you would from any
Linux, UNIX, or even Windows command prompt.</p>
<h2 id="other-reference-material" tabindex="-1"><a class="header-anchor" href="#other-reference-material" aria-hidden="true">#</a> Other Reference Material</h2>
<h3 id="language-guides" tabindex="-1"><a class="header-anchor" href="#language-guides" aria-hidden="true">#</a> Language Guides</h3>
<ol>
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
<p>The source for each of these builtins can be found on <a href="https://github.com/lmorg/murex/tree/master/builtins/core" target="_blank" rel="noopener noreferrer">Github<ExternalLinkIcon/></a>.</p>
<h3 id="shell-commands-for-querying-builtins" tabindex="-1"><a class="header-anchor" href="#shell-commands-for-querying-builtins" aria-hidden="true">#</a> Shell Commands For Querying Builtins</h3>
<p>From the shell itself: run <code v-pre>builtins</code> to list the builtin command.</p>
<p>If you require a manual on any of those commands, you can run <code v-pre>murex-docs</code>
to return the same markdown-formatted document as those listed below. eg</p>
<pre><code>murex-docs trypipe
</code></pre>
</div></template>


