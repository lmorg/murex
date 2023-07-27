<template><div><h1 id="exit" tabindex="-1"><a class="header-anchor" href="#exit" aria-hidden="true">#</a> <code v-pre>exit</code></h1>
<blockquote>
<p>Exit murex</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Exit's Murex with either a exit number of 0 (by default if no parameters
supplied) or a custom value specified by the first parameter.</p>
<p><code v-pre>exit</code> is not scope aware; if it is included in a function then the whole
shell will still exist and not just that function.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>exit
exit number
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» exit

» exit 42
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/break.html"><code v-pre>break</code></RouterLink>:
Terminate execution of a block within your processes scope</li>
<li><RouterLink to="/commands/die.html"><code v-pre>die</code></RouterLink>:
Terminate murex with an exit number of 1</li>
<li><RouterLink to="/commands/devnull.html"><code v-pre>null</code></RouterLink>:
null function. Similar to /dev/null</li>
</ul>
</div></template>


