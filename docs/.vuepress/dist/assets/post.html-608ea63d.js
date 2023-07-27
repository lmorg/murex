import{_ as s}from"./plugin-vue_export-helper-c27b6911.js";import{r as d,o as r,c as i,d as e,b as a,w as n,e as t,f as c}from"./app-45f7c304.js";const l={},h=c(`<h1 id="post" tabindex="-1"><a class="header-anchor" href="#post" aria-hidden="true">#</a> <code>post</code></h1><blockquote><p>HTTP POST request with a JSON-parsable return</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Fetches a page from a URL via HTTP/S POST request.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>post url -&gt; \`&lt;stdout&gt;\`

\`&lt;stdin&gt;\` -&gt; post url content-type -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» post google.com -&gt; [ Status ]
{
    &quot;Code&quot;: 405,
    &quot;Message&quot;: &quot;Method Not Allowed&quot;
}
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="json-return" tabindex="-1"><a class="header-anchor" href="#json-return" aria-hidden="true">#</a> JSON return</h3><p><code>POST</code> returns a JSON object with the following fields:</p><pre><code>{
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
</code></pre><p>The concept behind this is it provides and easier path for scripting eg pulling specific fields via the index, <code>[</code>, function.</p><h3 id="post-as-a-method" tabindex="-1"><a class="header-anchor" href="#post-as-a-method" aria-hidden="true">#</a> <code>post</code> as a method</h3><p>Running <code>post</code> as a method will transmit the contents of STDIN as part of the body of the HTTP POST request. When run as a method you have to include a second parameter specifying the Content-Type MIME.</p><h3 id="configurable-options" tabindex="-1"><a class="header-anchor" href="#configurable-options" aria-hidden="true">#</a> Configurable options</h3><p><code>post</code> has a number of behavioral options which can be configured via Murex&#39;s standard <code>config</code> tool:</p><pre><code>config: -&gt; [ http ]
</code></pre><p>To change a default, for example the user agent string:</p><pre><code>config: set http user-agent &quot;bob&quot;
post: google.com
</code></pre><p>This enables sane, repeatable and readable defaults. Read the documents on <code>config</code> for more details about it&#39;s usage and the rational behind the command.</p><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,22),u=e("code",null,"[[",-1),p=e("code",null,"[",-1),f=e("code",null,"config",-1),m=e("code",null,"get",-1),g=e("code",null,"getfile",-1);function b(_,q){const o=d("RouterLink");return r(),i("div",null,[h,e("ul",null,[e("li",null,[a(o,{to:"/commands/element.html"},{default:n(()=>[u,t(" (element)")]),_:1}),t(": Outputs an element from a nested structure")]),e("li",null,[a(o,{to:"/commands/index2.html"},{default:n(()=>[p,t(" (index)")]),_:1}),t(": Outputs an element from an array, map or table")]),e("li",null,[a(o,{to:"/commands/config.html"},{default:n(()=>[f]),_:1}),t(": Query or define Murex runtime settings")]),e("li",null,[a(o,{to:"/commands/get.html"},{default:n(()=>[m]),_:1}),t(": Makes a standard HTTP request and returns the result as a JSON object")]),e("li",null,[a(o,{to:"/commands/getfile.html"},{default:n(()=>[g]),_:1}),t(": Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.")])])])}const S=s(l,[["render",b],["__file","post.html.vue"]]);export{S as default};
