<template><div><h1 id="getfile" tabindex="-1"><a class="header-anchor" href="#getfile" aria-hidden="true">#</a> <code v-pre>getfile</code></h1>
<blockquote>
<p>Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Fetches a resource from a URL - setting STDOUT data-type</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>getfile url -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>getfile google.com
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>This simply fetches a resource (via HTTP GET request) from a URL and returns the
byte stream to STDOUT. It will set STDOUT's data-type based on MIME defined in
the <code v-pre>Content-Type</code> HTTP header.</p>
<p>It is recommended that you only use this command if you're pipelining the output
(eg writing to file or passing on to another function). If you just want to
render the output to the terminal then use <code v-pre>open</code> which has hooks for smart
terminal rendering.</p>
<h3 id="configurable-options" tabindex="-1"><a class="header-anchor" href="#configurable-options" aria-hidden="true">#</a> Configurable options</h3>
<p><code v-pre>getfile</code> has a number of behavioral options which can be configured via
Murex's standard <code v-pre>config</code> tool:</p>
<pre><code>config: -&gt; [ http ]
</code></pre>
<p>To change a default, for example the user agent string:</p>
<pre><code>config: set http user-agent &quot;bob&quot;
getfile: google.com
</code></pre>
<p>This enables sane, repeatable and readable defaults. Read the documents on
<code v-pre>config</code> for more details about it's usage and the rational behind the command.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/get.html"><code v-pre>get</code></RouterLink>:
Makes a standard HTTP request and returns the result as a JSON object</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/post.html"><code v-pre>post</code></RouterLink>:
HTTP POST request with a JSON-parsable return</li>
</ul>
</div></template>


