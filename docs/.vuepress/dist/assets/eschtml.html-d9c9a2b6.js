import{_ as c}from"./plugin-vue_export-helper-c27b6911.js";import{r as l,o as d,c as n,d as e,b as a,w as s,e as o,f as r}from"./app-45f7c304.js";const h={},i=r('<h1 id="eschtml" tabindex="-1"><a class="header-anchor" href="#eschtml" aria-hidden="true">#</a> <code>eschtml</code></h1><blockquote><p>Encode or decodes text for HTML</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>eschtml</code> takes input from either STDIN or the parameters and returns the same data, HTML escaped.</p><p><code>!eschtml</code> does the same process in reverse, where it takes HTML escaped data and returns its unescaped counterpart.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><p>Escape</p><pre><code>`&lt;stdin&gt;` -&gt; eschtml -&gt; `&lt;stdout&gt;`\n\neschtml string to escape -&gt; `&lt;stdout&gt;`\n</code></pre><p>Unescape</p><pre><code>`&lt;stdin&gt;` -&gt; !eschtml -&gt; `&lt;stdout&gt;`\n\n!eschtml string to unescape -&gt; `&lt;stdout&gt;`\n</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Escape</p><pre><code>» out: &quot;&lt;h1&gt;foo &amp; bar&lt;/h1&gt;&quot; -&gt; eschtml\n&amp;lt;h1&amp;gt;foo &amp;amp; bar&amp;lt;/h1&amp;gt;\n</code></pre><p>Unescape</p><pre><code>» out: &#39;&amp;lt;h1&amp;gt;foo &amp;amp; bar&amp;lt;/h1&amp;gt;&#39; -&gt; !eschtml\n&lt;h1&gt;foo &amp; bar&lt;/h1&gt;\n</code></pre><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>eschtml</code></li><li><code>!eschtml</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>',18),p=e("code",null,"escape",-1),u=e("code",null,"esccli",-1),m=e("code",null,"escurl",-1),_=e("code",null,"get",-1),g=e("code",null,"getfile",-1),f=e("code",null,"post",-1);function x(b,T){const t=l("RouterLink");return d(),n("div",null,[i,e("ul",null,[e("li",null,[a(t,{to:"/commands/escape.html"},{default:s(()=>[p]),_:1}),o(": Escape or unescape input")]),e("li",null,[a(t,{to:"/commands/esccli.html"},{default:s(()=>[u]),_:1}),o(": Escapes an array so output is valid shell code")]),e("li",null,[a(t,{to:"/commands/escurl.html"},{default:s(()=>[m]),_:1}),o(": Encode or decodes text for the URL")]),e("li",null,[a(t,{to:"/commands/get.html"},{default:s(()=>[_]),_:1}),o(": Makes a standard HTTP request and returns the result as a JSON object")]),e("li",null,[a(t,{to:"/commands/getfile.html"},{default:s(()=>[g]),_:1}),o(": Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.")]),e("li",null,[a(t,{to:"/commands/post.html"},{default:s(()=>[f]),_:1}),o(": HTTP POST request with a JSON-parsable return")])])])}const E=c(h,[["render",x],["__file","eschtml.html.vue"]]);export{E as default};