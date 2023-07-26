<template><div><h1 id="pipe" tabindex="-1"><a class="header-anchor" href="#pipe" aria-hidden="true">#</a> <code v-pre>pipe</code></h1>
<blockquote>
<p>Manage Murex named pipes</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>pipe</code> creates and destroys Murex named pipes.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Create pipe</p>
<pre><code>pipe: name [ pipe-type ]
</code></pre>
<p>Destroy pipe</p>
<pre><code>!pipe: name
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Create a standard pipe:</p>
<pre><code>pipe: example
</code></pre>
<p>Delete a pipe:</p>
<pre><code>!pipe: example
</code></pre>
<p>Create a TCP pipe (deleting a pipe is the same regardless of the type of pipe):</p>
<pre><code>pipe example --tcp-dial google.com:80
bg { &lt;example&gt; }
out: &quot;GET /&quot; -&gt; &lt;example&gt;
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="what-are-murex-named-pipes" tabindex="-1"><a class="header-anchor" href="#what-are-murex-named-pipes" aria-hidden="true">#</a> What are Murex named pipes?</h3>
<p>In POSIX, there is a concept of STDIN, STDOUT and STDERR, these are FIFO files
while are &quot;piped&quot; from one executable to another. ie STDOUT for application 'A'
would be the same file as STDIN for application 'B' when A is piped to B:
<code v-pre>A | B</code>. Murex adds a another layer around this to enable support for passing
data types and builtins which are agnostic to the data serialization format
traversing the pipeline. While this does add overhead the advantage is this new
wrapper can be used as a primitive for channelling any data from one point to
another.</p>
<p>Murex named pipes are where these pipes are created in a global store,
decoupled from any executing functions, named and can then be used to pass
data along asynchronously.</p>
<p>For example</p>
<pre><code>pipe: example

bg {
    &lt;example&gt; -&gt; match: Hello
}

out: &quot;foobar&quot;        -&gt; &lt;example&gt;
out: &quot;Hello, world!&quot; -&gt; &lt;example&gt;
out: &quot;foobar&quot;        -&gt; &lt;example&gt;

!pipe: example
</code></pre>
<p>This returns <code v-pre>Hello, world!</code> because <code v-pre>out</code> is writing to the <strong>example</strong> named
pipe and <code v-pre>match</code> is also reading from it in the background (<code v-pre>bg</code>).</p>
<p>Named pipes can also be inlined into the command parameters with <code v-pre>&lt;&gt;</code> tags</p>
<pre><code>pipe: example

bg {
    &lt;example&gt; -&gt; match: Hello
}

out: &lt;example&gt; &quot;foobar&quot;
out: &lt;example&gt; &quot;Hello, world!&quot;
out: &lt;example&gt; &quot;foobar&quot;

!pipe: example
</code></pre>
<blockquote>
<p>Please note this is also how <code v-pre>test</code> works.</p>
</blockquote>
<p>Murex named pipes can also represent network sockets, files on a disk or any
other read and/or write endpoint. Custom builtins can also be written in Golang
to support different abstractions so your Murex code can work with those read
or write endpoints transparently.</p>
<p>To see the different supported types run</p>
<pre><code>runtime --pipes
</code></pre>
<h3 id="namespaces-and-usage-in-modules-and-packages" tabindex="-1"><a class="header-anchor" href="#namespaces-and-usage-in-modules-and-packages" aria-hidden="true">#</a> Namespaces and usage in modules and packages</h3>
<p>Pipes created via <code v-pre>pipe</code> are created in the global namespace. This allows pipes
to be used across different functions easily however it does pose a risk with
name clashes where Murex named pipes are used heavily. Thus is it recommended
that pipes created in modules should be prefixed with the name of its package.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>pipe</code></li>
<li><code v-pre>!pipe</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/pipeline.html">Pipeline</RouterLink>:
Overview of what a &quot;pipeline&quot; is</li>
<li><RouterLink to="/commands/namedpipe.html"><code v-pre>&lt;&gt;</code> / <code v-pre>read-named-pipe</code></RouterLink>:
Reads from a Murex named pipe</li>
<li><RouterLink to="/commands/namedpipe.html"><code v-pre>&lt;&gt;</code> / <code v-pre>read-named-pipe</code></RouterLink>:
Reads from a Murex named pipe</li>
<li><RouterLink to="/commands/stdin.html"><code v-pre>&lt;stdin&gt;</code> </RouterLink>:
Read the STDIN belonging to the parent code block</li>
<li><RouterLink to="/commands/bg.html"><code v-pre>bg</code></RouterLink>:
Run processes in the background</li>
<li><RouterLink to="/commands/match.html"><code v-pre>match</code></RouterLink>:
Match an exact value in an array</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/test.html"><code v-pre>test</code></RouterLink>:
Murex's test framework - define tests, run tests and debug shell scripts</li>
</ul>
</div></template>


