<template><div><h1 id="exec" tabindex="-1"><a class="header-anchor" href="#exec" aria-hidden="true">#</a> <code v-pre>exec</code></h1>
<blockquote>
<p>Runs an executable</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>With Murex, like most other shells, you launch a process by calling the
name of that executable directly. While this is suitable 99% of the time,
occasionally you might run into an edge case where that wouldn't work. The
primary reason being if you needed to launch a process from a variable, eg</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» set exe=uname
» $exe
uname
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>As you can see here, Murex's behavior here is to output the contents of
the variable rather then executing the contents of the variable. This is
done for safety reasons, however if you wanted to override that behavior
then you could prefix the variable with exec:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» set exe=uname
» exec $exe
Linux
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    &lt;stdin> -> exec
    &lt;stdin> -> exec -> &lt;stdout>
               exec -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» exec printf "Hello, world!"
Hello, world!
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p><strong>Working around aliases:</strong></p>
<p>If you have an alias like <code v-pre>alias ls=ls --color=auto</code> and you wanted to run <code v-pre>ls</code>
but without colour, you might run <code v-pre>exec ls</code>.</p>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>If any command doesn't exist as a builtin, function nor alias, then Murex
will default to forking out to any command with this name (subject to an
absolute path or the order of precedence in <code v-pre>$PATH</code>). Any forked process will
show up in both the operating systems process viewer (eg <code v-pre>ps</code>) but also
Murex's own process viewer, <code v-pre>fid-list</code>. However inside <code v-pre>fid-list</code> you will
notice that all external processes are listed as <code v-pre>exec</code> with the process name
as part of <code v-pre>exec</code>'s parameters. That is because <code v-pre>exec</code> is handler for programs
that aren't native to Murex.</p>
<h3 id="compatibility-with-posix" tabindex="-1"><a class="header-anchor" href="#compatibility-with-posix" aria-hidden="true">#</a> Compatibility with POSIX</h3>
<p>For compatibility with traditional shells like Bash and Zsh, <code v-pre>command</code> is an
alias for <code v-pre>exec</code>.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>exec</code></li>
<li><code v-pre>command</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/bexists.html"><code v-pre>bexists</code></RouterLink>:
Check which builtins exist</li>
<li><RouterLink to="/commands/bg.html"><code v-pre>bg</code></RouterLink>:
Run processes in the background</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>builtins</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/fexec.html"><code v-pre>fexec</code> </RouterLink>:
Execute a command or function, bypassing the usual order of precedence.</li>
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
<li><RouterLink to="/commands/murex-update-exe-list.html"><code v-pre>murex-update-exe-list</code></RouterLink>:
Forces Murex to rescan $PATH looking for executables</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


