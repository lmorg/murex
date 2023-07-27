<template><div><h1 id="named-pipes" tabindex="-1"><a class="header-anchor" href="#named-pipes" aria-hidden="true">#</a> Named Pipes</h1>
<blockquote>
<p>A detailed breakdown of named pipes in Murex</p>
</blockquote>
<h2 id="background" tabindex="-1"><a class="header-anchor" href="#background" aria-hidden="true">#</a> Background</h2>
<p><a href="https://en.wikipedia.org/wiki/Named_pipe" target="_blank" rel="noopener noreferrer">Wikipedia describes<ExternalLinkIcon/></a> a named pipe as the following:</p>
<blockquote>
<p>In computing, a named pipe (also known as a FIFO for its behavior) is an
extension to the traditional pipe concept on Unix and Unix-like systems, and
is one of the methods of inter-process communication (IPC). The concept is
also found in OS/2 and Microsoft Windows, although the semantics differ
substantially. A traditional pipe is &quot;unnamed&quot; and lasts only as long as the
process. A named pipe, however, can last as long as the system is up, beyond
the life of the process. It can be deleted if no longer used. Usually a named
pipe appears as a file, and generally processes attach to it for IPC.</p>
</blockquote>
<p>Where Murex differs from standard Linux/UNIX is that named pipes are not
special files but rather an object or construct within the shell runtime. This
allows for more user friendly tooling and syntactic sugar to implemented around
it while largely still having the same functionality as a more traditional file
based named pipe.</p>
<h2 id="in-murex" tabindex="-1"><a class="header-anchor" href="#in-murex" aria-hidden="true">#</a> In Murex</h2>
<p>In Murex, named pipes are described in code as a value inside angle brackets.
There are four named pipes pre-configured: <code v-pre>&lt;in&gt;</code> (STDIN), <code v-pre>&lt;out&gt;</code> (STDOUT),
<code v-pre>&lt;err&gt;</code> (STDERR), and <code v-pre>&lt;null&gt;</code> (/dev/null equivalent).</p>
<p>You can call a named pipe as either a method, function, or parameter.</p>
<p><strong>As a method:</strong></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>&lt;in> -> command parameter1 parameter2 parameter3
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p><strong>As a function:</strong></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>command parameter1 parameter2 parameter3 -> &lt;out>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p><strong>As a parameter:</strong></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>command &lt;out> &lt;!err> parameter1 parameter2 parameter3
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/namedpipe.html"><code v-pre>&lt;&gt;</code> / <code v-pre>read-named-pipe</code></RouterLink>:
Reads from a Murex named pipe</li>
<li><RouterLink to="/commands/stdin.html"><code v-pre>&lt;stdin&gt;</code> </RouterLink>:
Read the STDIN belonging to the parent code block</li>
<li><RouterLink to="/commands/test.html"><code v-pre>test</code></RouterLink>:
Murex's test framework - define tests, run tests and debug shell scripts</li>
</ul>
</div></template>


