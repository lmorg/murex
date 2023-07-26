<template><div><h1 id="append-pipe-token" tabindex="-1"><a class="header-anchor" href="#append-pipe-token" aria-hidden="true">#</a> Append Pipe (<code v-pre>&gt;&gt;</code>) Token</h1>
<blockquote>
<p>Redirects STDOUT to a file and append its contents</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This is used to redirect the STDOUT of a command and append it to a file. If
that file does not exist, then the file is created.</p>
<p>This behaves similarly to the <a href="https://www.gnu.org/software/bash/manual/bash.html#Appending-Redirected-Output" target="_blank" rel="noopener noreferrer">Bash (et al) token<ExternalLinkIcon/></a>
except it doesn't support adding alternative file descriptor numbers. Instead
you will need to use named pipes to achieve the same effect in Murex.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» out: &quot;Hello&quot; &gt;&gt; example.txt
» out: &quot;World!&quot; &gt;&gt; example.txt
» open: example.txt
Hello
World!
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>This is just syntactic sugar for <code v-pre>-&gt; &gt;&gt;</code>. Thus when the parser reads code like
the following:</p>
<pre><code>echo &quot;foobar&quot; &gt;&gt; example.txt
</code></pre>
<p>it will compile an abstract syntax tree which would reflect the following code
instead:</p>
<pre><code>echo &quot;foobar&quot; | &gt;&gt; example.txt
</code></pre>
<h3 id="truncating-a-file" tabindex="-1"><a class="header-anchor" href="#truncating-a-file" aria-hidden="true">#</a> Truncating a file</h3>
<p>To truncate a file (ie overwrite its contents) use <code v-pre>|&gt;</code> instead.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/pipe-arrow.html">Arrow Pipe (<code v-pre>-&gt;</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/parser/pipe-posix.html">POSIX Pipe (<code v-pre>|</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/user-guide/pipeline.html">Pipeline</RouterLink>:
Overview of what a &quot;pipeline&quot; is</li>
<li><RouterLink to="/parser/pipe-err.html">STDERR Pipe (<code v-pre>?</code>) Token</RouterLink>:
Pipes STDERR from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/commands/namedpipe.html"><code v-pre>&lt;&gt;</code> / <code v-pre>read-named-pipe</code></RouterLink>:
Reads from a Murex named pipe</li>
<li><RouterLink to="/commands/greater-than-greater-than.html"><code v-pre>&gt;&gt;</code> (append file)</RouterLink>:
Writes STDIN to disk - appending contents if file already exists</li>
<li><RouterLink to="/commands/greater-than.html"><code v-pre>&gt;</code> (truncate file)</RouterLink>:
Writes STDIN to disk - overwriting contents if file already exists</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
</ul>
</div></template>


