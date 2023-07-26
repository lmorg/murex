import{_ as i}from"./plugin-vue_export-helper-c27b6911.js";import{r as d,o as s,c as r,d as e,b as o,w as n,e as t,f as l}from"./app-45f7c304.js";const c={},p=l(`<h1 id="stdin" tabindex="-1"><a class="header-anchor" href="#stdin" aria-hidden="true">#</a> <code>&lt;stdin&gt;</code></h1><blockquote><p>Read the STDIN belonging to the parent code block</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>This is used inside functions and other code blocks to pass that block&#39;s STDIN down a pipeline</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>\`&lt;stdin&gt;\` -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>When writing more complex scripts, you cannot always invoke your read as the first command in a code block. For example a simple pipeline might be:</p><pre><code>» function: example { -&gt; match: 2 }
</code></pre><p>But this only works if <code>-&gt;</code> is the very first command. The following would fail:</p><pre><code># Incorrect code
function: example {
    out: &quot;only match 2&quot;
    -&gt; match 2
}
</code></pre><p>This is where <code>&lt;stdin&gt;</code> comes to our rescue:</p><pre><code>function: example {
    out: &quot;only match 2&quot;
    \`&lt;stdin&gt;\` -&gt; match 2
}
</code></pre><p>This could also be written as:</p><pre><code>function: example { out: &quot;only match 2&quot;; \`&lt;stdin&gt;\` -&gt; match 2 }
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p><code>&lt;stdin&gt;</code> makes use of a feature called <strong>named pipes</strong>, which are a way of piping data between processes without chaining them together as a single command pipeline (eg commands delimited with <code>|</code>, <code>-&gt;</code>, <code>=&gt;</code>, <code>?</code> tokens).</p><h3 id="what-are-murex-named-pipes" tabindex="-1"><a class="header-anchor" href="#what-are-murex-named-pipes" aria-hidden="true">#</a> What are Murex named pipes?</h3><p>In POSIX, there is a concept of STDIN, STDOUT and STDERR, these are FIFO files while are &quot;piped&quot; from one executable to another. ie STDOUT for application &#39;A&#39; would be the same file as STDIN for application &#39;B&#39; when A is piped to B: <code>A | B</code>. Murex adds a another layer around this to enable support for passing data types and builtins which are agnostic to the data serialization format traversing the pipeline. While this does add overhead the advantage is this new wrapper can be used as a primitive for channelling any data from one point to another.</p><p>Murex named pipes are where these pipes are created in a global store, decoupled from any executing functions, named and can then be used to pass data along asynchronously.</p><p>For example</p><pre><code>pipe: example

bg {
    &lt;example&gt; -&gt; match: Hello
}

out: &quot;foobar&quot;        -&gt; &lt;example&gt;
out: &quot;Hello, world!&quot; -&gt; &lt;example&gt;
out: &quot;foobar&quot;        -&gt; &lt;example&gt;

!pipe: example
</code></pre><p>This returns <code>Hello, world!</code> because <code>out</code> is writing to the <strong>example</strong> named pipe and <code>match</code> is also reading from it in the background (<code>bg</code>).</p><p>Named pipes can also be inlined into the command parameters with <code>&lt;&gt;</code> tags</p><pre><code>pipe: example

bg {
    &lt;example&gt; -&gt; match: Hello
}

out: &lt;example&gt; &quot;foobar&quot;
out: &lt;example&gt; &quot;Hello, world!&quot;
out: &lt;example&gt; &quot;foobar&quot;

!pipe: example
</code></pre><blockquote><p>Please note this is also how <code>test</code> works.</p></blockquote><p>Murex named pipes can also represent network sockets, files on a disk or any other read and/or write endpoint. Custom builtins can also be written in Golang to support different abstractions so your Murex code can work with those read or write endpoints transparently.</p><p>To see the different supported types run</p><pre><code>runtime --pipes
</code></pre><h3 id="namespaces-and-usage-in-modules-and-packages" tabindex="-1"><a class="header-anchor" href="#namespaces-and-usage-in-modules-and-packages" aria-hidden="true">#</a> Namespaces and usage in modules and packages</h3><p>Pipes created via <code>pipe</code> are created in the global namespace. This allows pipes to be used across different functions easily however it does pose a risk with name clashes where Murex named pipes are used heavily. Thus is it recommended that pipes created in modules should be prefixed with the name of its package.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>&lt;stdin&gt;</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,34),h=e("code",null,"<>",-1),u=e("code",null,"read-named-pipe",-1),m=e("code",null,"function",-1),g=e("code",null,"match",-1),f=e("code",null,"out",-1),x=e("code",null,"pipe",-1),b=e("code",null,"runtime",-1);function w(_,k){const a=d("RouterLink");return s(),r("div",null,[p,e("ul",null,[e("li",null,[o(a,{to:"/user-guide/pipeline.html"},{default:n(()=>[t("Pipeline")]),_:1}),t(': Overview of what a "pipeline" is')]),e("li",null,[o(a,{to:"/commands/namedpipe.html"},{default:n(()=>[h,t(" / "),u]),_:1}),t(": Reads from a Murex named pipe")]),e("li",null,[o(a,{to:"/commands/function.html"},{default:n(()=>[m]),_:1}),t(": Define a function block")]),e("li",null,[o(a,{to:"/commands/match.html"},{default:n(()=>[g]),_:1}),t(": Match an exact value in an array")]),e("li",null,[o(a,{to:"/commands/out.html"},{default:n(()=>[f]),_:1}),t(": Print a string to the STDOUT with a trailing new line character")]),e("li",null,[o(a,{to:"/commands/pipe.html"},{default:n(()=>[x]),_:1}),t(": Manage Murex named pipes")]),e("li",null,[o(a,{to:"/commands/runtime.html"},{default:n(()=>[b]),_:1}),t(": Returns runtime information on the internal state of Murex")])])])}const T=i(c,[["render",w],["__file","stdin.html.vue"]]);export{T as default};
