<template><div><h1 id="os" tabindex="-1"><a class="header-anchor" href="#os" aria-hidden="true">#</a> <code v-pre>os</code></h1>
<blockquote>
<p>Output the auto-detected OS name</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Output the auto-detected OS name.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>os -&gt; `&lt;stdout&gt;`

os string -&gt; `&lt;stdout&gt;`
```
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» os
linux
</code></pre>
<p>Or if you want to check if the host is one of a number of platforms:</p>
<pre><code># When run on Linux or FreeBSD
» os linux freebsd
true

# When run on another platform, eg Windows or Darwin (OSX)
# (exit number would also be `1`)
» os linux freebsd
false
</code></pre>
<p><code v-pre>posix</code> is also supported:</p>
<pre><code># When run on Linux, FreeBSD or Darwin (for example)
» os posix
true

# When run on Windows or Plan 9
# (exit number would also be `1`)
» os posix
false
</code></pre>
<p>Please note that although Plan 9 shares similarities with POSIX, it is not
POSIX-compliant. For that reason, <code v-pre>os</code> returns false with the <code v-pre>posix</code>
parameter when run on Plan 9. If you want to include Plan 9 in the check
then please write it as <code v-pre>os posix plan9</code>.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/cpuarch.html"><code v-pre>cpuarch</code></RouterLink>:
Output the hosts CPU architecture</li>
<li><RouterLink to="/commands/cpucount.html"><code v-pre>cpucount</code></RouterLink>:
Output the number of CPU cores available on your host</li>
</ul>
</div></template>


