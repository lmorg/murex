<template><div><h1 id="lockfile" tabindex="-1"><a class="header-anchor" href="#lockfile" aria-hidden="true">#</a> <code v-pre>lockfile</code></h1>
<blockquote>
<p>Create and manage lock files</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>lockfile</code> is used to create and manage lock files</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Create a lock file with the name <code v-pre>identifier</code></p>
<pre><code>lockfile: lock identifier
</code></pre>
<p>Delete a lock file with the name <code v-pre>identifier</code></p>
<pre><code>lockfile: unlock identifier
</code></pre>
<p>Wait until lock file with the name <code v-pre>identifier</code> has been deleted</p>
<pre><code>lockfile: wait identifier
</code></pre>
<p>Output the the file name and path of a lock file with the name <code v-pre>identifier</code></p>
<pre><code>lockfile: path identifier -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>lockfile: lock example
out: &quot;lock file created: ${lockfile path example}&quot;

bg {
    sleep: 10
    lockfile: unlock example
}

out: &quot;waiting for lock file to be deleted (sleep 10 seconds)....&quot;
lockfile: wait example
out: &quot;lock file gone!&quot;
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/bg.html"><code v-pre>bg</code></RouterLink>:
Run processes in the background</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
</ul>
</div></template>


