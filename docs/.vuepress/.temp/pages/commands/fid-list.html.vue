<template><div><h1 id="fid-list" tabindex="-1"><a class="header-anchor" href="#fid-list" aria-hidden="true">#</a> <code v-pre>fid-list</code></h1>
<blockquote>
<p>Lists all running functions within the current Murex session</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>fid-list</code> is a tool for outputting all the functions currently managed by that
Murex session. Those functions could be Murex functions, builtins or any
external executables launched from Murex.</p>
<p>Conceptually <code v-pre>fid-list</code> is a little like <code v-pre>ps</code> (on POSIX systems) however
<code v-pre>fid-list</code> was not written to be POSIX compliant.</p>
<p>Multiple flags cannot be used with each other.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>fid-list [ flag ] -&gt; `&lt;stdout&gt;`
</code></pre>
<p><code v-pre>jobs</code> is an alias for <code v-pre>fid-list: --jobs</code>:
jobs -&gt; <code v-pre>&lt;stdout&gt;</code></p>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--background</code>
Returns a <code v-pre>json</code> map of background jobs</li>
<li><code v-pre>--csv</code>
Output table in a <code v-pre>csv</code> format</li>
<li><code v-pre>--help</code>
Outputs a list of parameters and a descriptions</li>
<li><code v-pre>--jobs</code>
Show background and stopped jobs</li>
<li><code v-pre>--jsonl</code>
Output table in a jsonlines (<code v-pre>jsonl</code>) format (defaulted to when piped)</li>
<li><code v-pre>--stopped</code>
Returns a <code v-pre>json</code> map of stopped jobs</li>
<li><code v-pre>--tty</code>
Force default TTY output even when piped</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>Because Murex is a multi-threaded shell, builtins are not forked processes
like in a traditional / POSIX shell. This means that you cannot use the
operating systems default process viewer (eg <code v-pre>ps</code>) to list Murex functions.
This is where <code v-pre>fid-list</code> comes into play. It is used to view all the functions
and processes that are managed by the current Murex session. That would
include:</p>
<ul>
<li>any aliases within Murex</li>
<li>public and private Murex functions</li>
<li>builtins (eg <code v-pre>fid-list</code> is a builtin command)</li>
<li>any external processes that were launched from within this shell session</li>
<li>any background functions or processes of any of the above</li>
</ul>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>fid-list</code></li>
<li><code v-pre>jobs</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/types/generic.html"><code v-pre>*</code> (generic) </RouterLink>:
generic (primitive)</li>
<li><RouterLink to="/commands/bexists.html"><code v-pre>bexists</code></RouterLink>:
Check which builtins exist</li>
<li><RouterLink to="/commands/bg.html"><code v-pre>bg</code></RouterLink>:
Run processes in the background</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>builtins</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/types/csv.html"><code v-pre>csv</code> </RouterLink>:
CSV files (and other character delimited tables)</li>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/commands/fexec.html"><code v-pre>fexec</code> </RouterLink>:
Execute a command or function, bypassing the usual order of precedence.</li>
<li><RouterLink to="/commands/fg.html"><code v-pre>fg</code></RouterLink>:
Sends a background process into the foreground</li>
<li><RouterLink to="/commands/fid-kill.html"><code v-pre>fid-kill</code></RouterLink>:
Terminate a running Murex function</li>
<li><RouterLink to="/commands/fid-killall.html"><code v-pre>fid-killall</code></RouterLink>:
Terminate <em>all</em> running Murex functions</li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>jobs</code></RouterLink>:
Lists all running functions within the current Murex session</li>
<li><RouterLink to="/types/jsonl.html"><code v-pre>jsonl</code> </RouterLink>:
JSON Lines</li>
<li><RouterLink to="/commands/murex-update-exe-list.html"><code v-pre>murex-update-exe-list</code></RouterLink>:
Forces Murex to rescan $PATH looking for executables</li>
</ul>
</div></template>


