<template><div><h1 id="summary" tabindex="-1"><a class="header-anchor" href="#summary" aria-hidden="true">#</a> <code v-pre>summary</code></h1>
<blockquote>
<p>Defines a summary help text for a command</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>summary</code> define help text for a command. This is effectively like a tooltip
message that appears, by default, in blue in the interactive shell.</p>
<p>Normally this text is populated from the <code v-pre>man</code> pages or <code v-pre>murex-docs</code>, however
if neither exist or if you wish to override their text, then you can use
<code v-pre>summary</code> to define that text.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Define a commands summary</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>summary command description
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Undefine a summary</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>!summary command
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Define a commands summary</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» summary: foobar "Hello, world!"
» runtime: --summaries -> [ foobar ]
Hello, world!
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Undefine a summary</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» !summary: foobar
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>summary</code></li>
<li><code v-pre>!summary</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/bexists.html"><code v-pre>bexists</code></RouterLink>:
Check which builtins exist</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>builtins</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>fid-list</code></RouterLink>:
Lists all running functions within the current Murex session</li>
<li><RouterLink to="/commands/murex-docs.html"><code v-pre>murex-docs</code></RouterLink>:
Displays the man pages for Murex builtins</li>
<li><RouterLink to="/commands/murex-update-exe-list.html"><code v-pre>murex-update-exe-list</code></RouterLink>:
Forces Murex to rescan $PATH looking for executables</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


