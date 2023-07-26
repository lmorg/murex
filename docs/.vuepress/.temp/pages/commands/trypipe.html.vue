<template><div><h1 id="trypipe" tabindex="-1"><a class="header-anchor" href="#trypipe" aria-hidden="true">#</a> <code v-pre>trypipe</code></h1>
<blockquote>
<p>Checks state of each function in a pipeline and exits block on error</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>trypipe</code> checks the state of each function and exits the block if any of them
fail. Where <code v-pre>trypipe</code> differs from regular <code v-pre>try</code> blocks is <code v-pre>trypipe</code> will check
every process along the pipeline as well as the terminating function (which
<code v-pre>try</code> only validates against). The downside to this is that piped functions can
no longer run in parallel.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>trypipe { code-block } -> &lt;stdout>

&lt;stdin> -> trypipe { -> code-block } -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>trypipe {
    out: "Hello, World!" -> grep: "non-existent string" -> cat
    out: "This command will be ignored"
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Formated pager (<code v-pre>less</code>) where the pager isn't called if the formatter (<code v-pre>pretty</code>) fails (eg input isn't valid JSON):</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>func pless {
    -> trypipe { -> pretty -> less }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>A failure is determined by:</p>
<ul>
<li>Any process that returns a non-zero exit number</li>
<li>Any process that returns more output via STDERR than it does via STDOUT</li>
</ul>
<p>You can see which run mode your functions are executing under via the <code v-pre>fid-list</code>
command.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/schedulers.html">Schedulers</RouterLink>:
Overview of the different schedulers (or 'run modes') in Murex</li>
<li><RouterLink to="/commands/catch.html"><code v-pre>catch</code></RouterLink>:
Handles the exception code raised by <code v-pre>try</code> or <code v-pre>trypipe</code></li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>fid-list</code></RouterLink>:
Lists all running functions within the current Murex session</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/commands/runmode.html"><code v-pre>runmode</code></RouterLink>:
Alter the scheduler's behaviour at higher scoping level</li>
<li><RouterLink to="/commands/switch.html"><code v-pre>switch</code></RouterLink>:
Blocks of cascading conditionals</li>
<li><RouterLink to="/commands/try.html"><code v-pre>try</code></RouterLink>:
Handles errors inside a block of code</li>
</ul>
</div></template>


