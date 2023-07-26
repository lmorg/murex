<template><div><h1 id="version" tabindex="-1"><a class="header-anchor" href="#version" aria-hidden="true">#</a> <code v-pre>version</code></h1>
<blockquote>
<p>Get Murex version</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Returns Murex version number</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>version [ flags ] -&gt; &lt;stdout&gt;
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Ran without any parameters</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» version
murex: 0.51.1200 BETA
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>Ran with the <code v-pre>--no-app-name</code> parameter</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» version --no-app-name
0.51.1200 BETA
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>Ran with the <code v-pre>--short</code> parameter</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» version --short
0.51
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--no-app-name</code>
Returns full version string minus app name</li>
<li><code v-pre>--short</code>
Returns only the major and minor version as a <code v-pre>num</code> data-type</li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/args.html"><code v-pre>args</code> </RouterLink>:
Command line flag parser for Murex shell scripting</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/murex-parser.html"><code v-pre>murex-parser</code> </RouterLink>:
Runs the Murex parser against a block of code</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/source.html"><code v-pre>source</code> </RouterLink>:
Import Murex code from another file of code block</li>
</ul>
</div></template>


