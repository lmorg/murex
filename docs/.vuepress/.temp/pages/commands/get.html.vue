<template><div><h1 id="get" tabindex="-1"><a class="header-anchor" href="#get" aria-hidden="true">#</a> <code v-pre>get</code></h1>
<blockquote>
<p>Makes a standard HTTP request and returns the result as a JSON object</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Fetches a page from a URL via HTTP/S GET request</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>get url -&gt; `&lt;stdout&gt;`

`&lt;stdin&gt;` -&gt; get url -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» get google.com -&gt; [ Status ]
{
    &quot;Code&quot;: 200,
    &quot;Message&quot;: &quot;OK&quot;
}
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="json-return" tabindex="-1"><a class="header-anchor" href="#json-return" aria-hidden="true">#</a> JSON return</h3>
<p><code v-pre>get</code> returns a JSON object with the following fields:</p>
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

The concept behind this is it provides and easier path for scripting eg pulling
specific fields via the index, `[`, function.

### `get` as a method

Running `get` as a method will transmit the contents of STDIN as part of the
body of the HTTP GET request. When run as a method you have to include a second
parameter specifying the Content-Type MIME.

### Configurable options

`get` has a number of behavioral options which can be configured via Murex's
standard `config` tool:

config: -&gt; [ http ]

To change a default, for example the user agent string:

config: set http user-agent &quot;bob&quot;
</code></pre>
<p>get: google.com</p>
<pre><code>This enables sane, repeatable and readable defaults. Read the documents on
`config` for more details about it's usage and the rational behind the command.
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/getfile.html"><code v-pre>getfile</code></RouterLink>:
Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.</li>
<li><RouterLink to="/commands/post.html"><code v-pre>post</code></RouterLink>:
HTTP POST request with a JSON-parsable return</li>
</ul>
</div></template>


