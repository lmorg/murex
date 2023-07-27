import{_ as s}from"./plugin-vue_export-helper-c27b6911.js";import{r,o as d,c as i,d as e,b as a,w as o,e as t,f as l}from"./app-45f7c304.js";const u={},c=l(`<h1 id="get" tabindex="-1"><a class="header-anchor" href="#get" aria-hidden="true">#</a> <code>get</code></h1><blockquote><p>Makes a standard HTTP request and returns the result as a JSON object</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Fetches a page from a URL via HTTP/S GET request</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>get url -&gt; \`&lt;stdout&gt;\`

\`&lt;stdin&gt;\` -&gt; get url -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» get google.com -&gt; [ Status ]
{
    &quot;Code&quot;: 200,
    &quot;Message&quot;: &quot;OK&quot;
}
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="json-return" tabindex="-1"><a class="header-anchor" href="#json-return" aria-hidden="true">#</a> JSON return</h3><p><code>get</code> returns a JSON object with the following fields:</p><pre><code>{
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
specific fields via the index, \`[\`, function.

### \`get\` as a method

Running \`get\` as a method will transmit the contents of STDIN as part of the
body of the HTTP GET request. When run as a method you have to include a second
parameter specifying the Content-Type MIME.

### Configurable options

\`get\` has a number of behavioral options which can be configured via Murex&#39;s
standard \`config\` tool:

config: -&gt; [ http ]

To change a default, for example the user agent string:

config: set http user-agent &quot;bob&quot;
</code></pre><p>get: google.com</p><pre><code>This enables sane, repeatable and readable defaults. Read the documents on
\`config\` for more details about it&#39;s usage and the rational behind the command.
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,15),h=e("code",null,"[[",-1),g=e("code",null,"[",-1),p=e("code",null,"config",-1),f=e("code",null,"getfile",-1),m=e("code",null,"post",-1);function _(b,q){const n=r("RouterLink");return d(),i("div",null,[c,e("ul",null,[e("li",null,[a(n,{to:"/commands/element.html"},{default:o(()=>[h,t(" (element)")]),_:1}),t(": Outputs an element from a nested structure")]),e("li",null,[a(n,{to:"/commands/index2.html"},{default:o(()=>[g,t(" (index)")]),_:1}),t(": Outputs an element from an array, map or table")]),e("li",null,[a(n,{to:"/commands/config.html"},{default:o(()=>[p]),_:1}),t(": Query or define Murex runtime settings")]),e("li",null,[a(n,{to:"/commands/getfile.html"},{default:o(()=>[f]),_:1}),t(": Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.")]),e("li",null,[a(n,{to:"/commands/post.html"},{default:o(()=>[m]),_:1}),t(": HTTP POST request with a JSON-parsable return")])])])}const S=s(u,[["render",_],["__file","get.html.vue"]]);export{S as default};
