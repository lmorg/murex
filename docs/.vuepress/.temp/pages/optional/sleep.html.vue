<template><div><h1 id="sleep" tabindex="-1"><a class="header-anchor" href="#sleep" aria-hidden="true">#</a> <code v-pre>sleep</code></h1>
<blockquote>
<p>Suspends the shell for a number of seconds</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>sleep</code> is an optional builtin which suspends the shell for a defined number
of seconds.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>sleep: integer
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» sleep 5
# murex sleeps for 5 seconds
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>sleep</code> is very simplistic - particularly when compared to its GNU coreutil
(for example) counterpart. If you want to use the <code v-pre>sleep</code> binary on Linux
or similar platforms then you will need to launch with the <code v-pre>exec</code> builtin:</p>
<pre><code>» exec: sleep 5
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/commands/source.html"><code v-pre>source</code> </RouterLink>:
Import Murex code from another file of code block</li>
<li><RouterLink to="/commands/time.html"><code v-pre>time</code> </RouterLink>:
Returns the execution run time of a command or block</li>
</ul>
</div></template>


