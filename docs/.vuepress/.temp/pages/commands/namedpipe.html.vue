<template><div><h1 id="read-named-pipe" tabindex="-1"><a class="header-anchor" href="#read-named-pipe" aria-hidden="true">#</a> <code v-pre>&lt;&gt;</code> / <code v-pre>read-named-pipe</code></h1>
<blockquote>
<p>Reads from a Murex named pipe</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Sometimes you will need to start a command line with a Murex named pipe, eg</p>
<pre><code>» &lt;namedpipe&gt; -&gt; match: foobar
</code></pre>
<blockquote>
<p>See the documentation on <code v-pre>pipe</code> for more details about Murex named pipes.</p>
</blockquote>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Read from pipe</p>
<pre><code>&lt;namedpipe&gt; -&gt; `&lt;stdout&gt;`
</code></pre>
<p>Write to pipe</p>
<pre><code>`&lt;stdin&gt;` -&gt; &lt;namedpipe&gt;
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>The follow two examples function the same</p>
<pre><code>» pipe: example
» bg { &lt;example&gt; -&gt; match: 2 }
» a: &lt;example&gt; [1..3]
2
» !pipe: example
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
<li><code v-pre>(murex named pipe)</code></li>
<li><code v-pre>&lt;&gt;</code></li>
<li><code v-pre>read-named-pipe</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/stdin.html"><code v-pre>&lt;stdin&gt;</code> </RouterLink>:
Read the STDIN belonging to the parent code block</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/bg.html"><code v-pre>bg</code></RouterLink>:
Run processes in the background</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/pipe.html"><code v-pre>pipe</code></RouterLink>:
Manage Murex named pipes</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


