<template><div><h1 id="fid-kill" tabindex="-1"><a class="header-anchor" href="#fid-kill" aria-hidden="true">#</a> <code v-pre>fid-kill</code></h1>
<blockquote>
<p>Terminate a running Murex function</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>fid-kill</code> will terminate a running Murex function in a similar way
that the POSIX <code v-pre>kill</code> (superficially speaking).</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>fid-kill fid
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>fid-kill</code> doesn't send a kernel signal to the process since Murex is
a multi-threaded shell with a single signal, <code v-pre>fid-kill</code> will send a
cancellation context to any builtins executing (which covers builtins,
aliases, public and private functions and any external executables running
which were launched within the current Murex shell).</p>
<p>The FID (function ID) sent is not the same as a POSIX (eg Linux, macOS, BSD)
PID (process ID). You can obtain a FID from <code v-pre>fid-list</code>.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/bexists.html"><code v-pre>bexists</code></RouterLink>:
Check which builtins exist</li>
<li><RouterLink to="/commands/bg.html"><code v-pre>bg</code></RouterLink>:
Run processes in the background</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>builtins</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/commands/fexec.html"><code v-pre>fexec</code> </RouterLink>:
Execute a command or function, bypassing the usual order of precedence.</li>
<li><RouterLink to="/commands/fg.html"><code v-pre>fg</code></RouterLink>:
Sends a background process into the foreground</li>
<li><RouterLink to="/commands/fid-killall.html"><code v-pre>fid-killall</code></RouterLink>:
Terminate <em>all</em> running Murex functions</li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>fid-list</code></RouterLink>:
Lists all running functions within the current Murex session</li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>jobs</code></RouterLink>:
Lists all running functions within the current Murex session</li>
<li><RouterLink to="/commands/murex-update-exe-list.html"><code v-pre>murex-update-exe-list</code></RouterLink>:
Forces Murex to rescan $PATH looking for executables</li>
</ul>
</div></template>


