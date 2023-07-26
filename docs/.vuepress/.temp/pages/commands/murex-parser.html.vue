<template><div><h1 id="murex-parser" tabindex="-1"><a class="header-anchor" href="#murex-parser" aria-hidden="true">#</a> <code v-pre>murex-parser</code></h1>
<blockquote>
<p>Runs the Murex parser against a block of code</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>summary</code> define help text for a command. This is effectively like a tooltip
message that appears, by default, in blue in the interactive shell.</p>
<p>Normally this text is populated from the <code v-pre>man</code> pages or <code v-pre>murex-docs</code>, however
if neither exist or if you wish to override their text, then you can use
<code v-pre>summary</code> to define that text.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; murex-parser -&gt; `&lt;stdout&gt;`

murex-parser { code-block } -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>Please note this command is still very much in beta and is likely to change in incompatible ways in the future. If you do happen to like this command and/or have any suggestions on how to improve it, then please leave your feedback on the GitHub repository, https://github.com/lmorg/murex</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/murex-docs.html"><code v-pre>murex-docs</code></RouterLink>:
Displays the man pages for Murex builtins</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


