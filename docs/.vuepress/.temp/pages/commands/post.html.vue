<template><div><h1 id="post" tabindex="-1"><a class="header-anchor" href="#post" aria-hidden="true">#</a> <code v-pre>post</code></h1>
<blockquote>
<p>HTTP POST request with a JSON-parsable return</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Fetches a page from a URL via HTTP/S POST request.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>post url -&gt; `&lt;stdout&gt;`

`&lt;stdin&gt;` -&gt; post url content-type -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» post google.com -&gt; [ Status ]
{
    &quot;Code&quot;: 405,
    &quot;Message&quot;: &quot;Method Not Allowed&quot;
}
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="json-return" tabindex="-1"><a class="header-anchor" href="#json-return" aria-hidden="true">#</a> JSON return</h3>
<p><code v-pre>POST</code> returns a JSON object with the following fields:</p>
<pre><code>{
    &quot;Status&quot;: {
        &quot;Code&quot;: integer,
        &quot;Message&quot;: string,
    },
    &quot;Headers&quot;: {
        string [
            string...
        ]
    },
    &quot;Body&quot;: string
}
</code></pre>
<p>The concept behind this is it provides and easier path for scripting eg pulling
specific fields via the index, <code v-pre>[</code>, function.</p>
<h3 id="post-as-a-method" tabindex="-1"><a class="header-anchor" href="#post-as-a-method" aria-hidden="true">#</a> <code v-pre>post</code> as a method</h3>
<p>Running <code v-pre>post</code> as a method will transmit the contents of STDIN as part of the
body of the HTTP POST request. When run as a method you have to include a second
parameter specifying the Content-Type MIME.</p>
<h3 id="configurable-options" tabindex="-1"><a class="header-anchor" href="#configurable-options" aria-hidden="true">#</a> Configurable options</h3>
<p><code v-pre>post</code> has a number of behavioral options which can be configured via Murex's
standard <code v-pre>config</code> tool:</p>
<pre><code>config: -&gt; [ http ]
</code></pre>
<p>To change a default, for example the user agent string:</p>
<pre><code>config: set http user-agent &quot;bob&quot;
post: google.com
</code></pre>
<p>This enables sane, repeatable and readable defaults. Read the documents on
<code v-pre>config</code> for more details about it's usage and the rational behind the command.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/get.html"><code v-pre>get</code></RouterLink>:
Makes a standard HTTP request and returns the result as a JSON object</li>
<li><RouterLink to="/commands/getfile.html"><code v-pre>getfile</code></RouterLink>:
Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.</li>
</ul>
</div></template>


