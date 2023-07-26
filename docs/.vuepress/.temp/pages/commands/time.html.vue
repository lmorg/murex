<template><div><h1 id="time" tabindex="-1"><a class="header-anchor" href="#time" aria-hidden="true">#</a> <code v-pre>time</code></h1>
<blockquote>
<p>Returns the execution run time of a command or block</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>time</code> is an optional builtin which runs a command or block of code and
returns it's running time.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>time: command parameters -&gt; &lt;stderr&gt;

time: { code-block } -&gt; &lt;stderr&gt;
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» time: sleep 5
5.000151513

» time { out &quot;Going to sleep&quot;; sleep 5; out &quot;Waking up&quot; }
Going to sleep
Waking up
5.000240977
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>time</code>'s output is written to STDERR. However any output and errors written
by the commands executed by time will also be written to <code v-pre>time</code>'s STDOUT
and STDERR as usual.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/optional/sleep.html"><code v-pre>sleep</code> </RouterLink>:
Suspends the shell for a number of seconds</li>
<li><RouterLink to="/commands/source.html"><code v-pre>source</code> </RouterLink>:
Import Murex code from another file of code block</li>
</ul>
</div></template>


