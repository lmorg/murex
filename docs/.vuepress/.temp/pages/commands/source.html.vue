<template><div><h1 id="source" tabindex="-1"><a class="header-anchor" href="#source" aria-hidden="true">#</a> <code v-pre>source</code></h1>
<blockquote>
<p>Import Murex code from another file of code block</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>source</code> imports code from another file or code block. It can be used as either
an &quot;import&quot; / &quot;include&quot; directive (eg Python, Go, C, etc) or an &quot;eval&quot; (eg
Python, Perl, etc).</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Execute source from STDIN</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>&lt;stdin> -> source
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Execute source from a file</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>source: filename.mx
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Execute a code block from parameter</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>source: { code-block }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Execute source from stdin:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» tout: block { out: "Hello, world!" } -> source
Hello, world!
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>Execute source from file:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» tout: block { out: "Hello, world!" } |> example.mx
» source: example.mx
Hello, world!
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Execute a code block from parameter</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» source { out: "Hello, world!" }
Hello, world!
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>source</code></li>
<li><code v-pre>.</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/args.html"><code v-pre>args</code> </RouterLink>:
Command line flag parser for Murex shell scripting</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/commands/fexec.html"><code v-pre>fexec</code> </RouterLink>:
Execute a command or function, bypassing the usual order of precedence.</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/murex-parser.html"><code v-pre>murex-parser</code> </RouterLink>:
Runs the Murex parser against a block of code</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/version.html"><code v-pre>version</code> </RouterLink>:
Get Murex version</li>
</ul>
</div></template>


