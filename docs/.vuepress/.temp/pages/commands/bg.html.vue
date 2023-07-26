<template><div><h1 id="bg" tabindex="-1"><a class="header-anchor" href="#bg" aria-hidden="true">#</a> <code v-pre>bg</code></h1>
<blockquote>
<p>Run processes in the background</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>bg</code> supports two modes: it can either be run as a function block which will
execute in the background, or it can take stopped processes and daemonize
them.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Any operating system:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    bg { code block }

    &lt;stdin> -> bg
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>POSIX only:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    bg { code block }

    &lt;stdin> -> bg

    bg fid
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>As a function:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>bg { sleep 5; out "Morning" }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>As a method:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ({ sleep 5; out "Morning" }) -> bg
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>The examples above will work on any system (Windows included). However the
<code v-pre>ctrl+z</code> usage of backgrounding a stopped process (like Bash) is only
supported on POSIX systems due to the limitation of required signals on
non-platforms. This means the usage described in the examples is cross
cross platform while <code v-pre>bg int</code> currently does not work on Windows nor Plan 9.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/commands/fg.html"><code v-pre>fg</code></RouterLink>:
Sends a background process into the foreground</li>
<li><RouterLink to="/commands/fid-kill.html"><code v-pre>fid-kill</code></RouterLink>:
Terminate a running Murex function</li>
<li><RouterLink to="/commands/fid-killall.html"><code v-pre>fid-killall</code></RouterLink>:
Terminate <em>all</em> running Murex functions</li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>fid-list</code></RouterLink>:
Lists all running functions within the current Murex session</li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>jobs</code></RouterLink>:
Lists all running functions within the current Murex session</li>
</ul>
</div></template>


