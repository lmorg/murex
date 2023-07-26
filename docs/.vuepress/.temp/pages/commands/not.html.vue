<template><div><h1 id="not" tabindex="-1"><a class="header-anchor" href="#not" aria-hidden="true">#</a> <code v-pre>!</code> (not)</h1>
<blockquote>
<p>Reads the STDIN and exit number from previous process and not's it's condition</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Reads the STDIN and exit number from previous process and not's it's condition.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; ! -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» echo &quot;Hello, world!&quot; -&gt; !
false

» false -&gt; !
true
</code></pre>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>!</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/and.html"><code v-pre>and</code></RouterLink>:
Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether multiple conditions are met</li>
<li><RouterLink to="/commands/false.html"><code v-pre>false</code></RouterLink>:
Returns a <code v-pre>false</code> value</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/commands/or.html"><code v-pre>or</code></RouterLink>:
Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.</li>
<li><RouterLink to="/commands/true.html"><code v-pre>true</code></RouterLink>:
Returns a <code v-pre>true</code> value</li>
</ul>
</div></template>


