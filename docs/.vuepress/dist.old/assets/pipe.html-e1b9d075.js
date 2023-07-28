import{_ as i}from"./plugin-vue_export-helper-c27b6911.js";import{r as d,o as s,c as p,d as e,b as o,w as n,e as a,f as r}from"./app-45f7c304.js";const l={},c=r(`<h1 id="pipe" tabindex="-1"><a class="header-anchor" href="#pipe" aria-hidden="true">#</a> <code>pipe</code></h1><blockquote><p>Manage Murex named pipes</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>pipe</code> creates and destroys Murex named pipes.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><p>Create pipe</p><pre><code>pipe: name [ pipe-type ]
</code></pre><p>Destroy pipe</p><pre><code>!pipe: name
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Create a standard pipe:</p><pre><code>pipe: example
</code></pre><p>Delete a pipe:</p><pre><code>!pipe: example
</code></pre><p>Create a TCP pipe (deleting a pipe is the same regardless of the type of pipe):</p><pre><code>pipe example --tcp-dial google.com:80
bg { &lt;example&gt; }
out: &quot;GET /&quot; -&gt; &lt;example&gt;
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="what-are-murex-named-pipes" tabindex="-1"><a class="header-anchor" href="#what-are-murex-named-pipes" aria-hidden="true">#</a> What are Murex named pipes?</h3><p>In POSIX, there is a concept of STDIN, STDOUT and STDERR, these are FIFO files while are &quot;piped&quot; from one executable to another. ie STDOUT for application &#39;A&#39; would be the same file as STDIN for application &#39;B&#39; when A is piped to B: <code>A | B</code>. Murex adds a another layer around this to enable support for passing data types and builtins which are agnostic to the data serialization format traversing the pipeline. While this does add overhead the advantage is this new wrapper can be used as a primitive for channelling any data from one point to another.</p><p>Murex named pipes are where these pipes are created in a global store, decoupled from any executing functions, named and can then be used to pass data along asynchronously.</p><p>For example</p><pre><code>pipe: example

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
</code></pre><h3 id="namespaces-and-usage-in-modules-and-packages" tabindex="-1"><a class="header-anchor" href="#namespaces-and-usage-in-modules-and-packages" aria-hidden="true">#</a> Namespaces and usage in modules and packages</h3><p>Pipes created via <code>pipe</code> are created in the global namespace. This allows pipes to be used across different functions easily however it does pose a risk with name clashes where Murex named pipes are used heavily. Thus is it recommended that pipes created in modules should be prefixed with the name of its package.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>pipe</code></li><li><code>!pipe</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,34),h=e("code",null,"<>",-1),u=e("code",null,"read-named-pipe",-1),m=e("code",null,"<>",-1),f=e("code",null,"read-named-pipe",-1),g=e("code",null,"<stdin>",-1),x=e("code",null,"bg",-1),_=e("code",null,"match",-1),b=e("code",null,"out",-1),w=e("code",null,"runtime",-1),y=e("code",null,"test",-1);function k(q,T){const t=d("RouterLink");return s(),p("div",null,[c,e("ul",null,[e("li",null,[o(t,{to:"/user-guide/pipeline.html"},{default:n(()=>[a("Pipeline")]),_:1}),a(': Overview of what a "pipeline" is')]),e("li",null,[o(t,{to:"/commands/namedpipe.html"},{default:n(()=>[h,a(" / "),u]),_:1}),a(": Reads from a Murex named pipe")]),e("li",null,[o(t,{to:"/commands/namedpipe.html"},{default:n(()=>[m,a(" / "),f]),_:1}),a(": Reads from a Murex named pipe")]),e("li",null,[o(t,{to:"/commands/stdin.html"},{default:n(()=>[g]),_:1}),a(": Read the STDIN belonging to the parent code block")]),e("li",null,[o(t,{to:"/commands/bg.html"},{default:n(()=>[x]),_:1}),a(": Run processes in the background")]),e("li",null,[o(t,{to:"/commands/match.html"},{default:n(()=>[_]),_:1}),a(": Match an exact value in an array")]),e("li",null,[o(t,{to:"/commands/out.html"},{default:n(()=>[b]),_:1}),a(": Print a string to the STDOUT with a trailing new line character")]),e("li",null,[o(t,{to:"/commands/runtime.html"},{default:n(()=>[w]),_:1}),a(": Returns runtime information on the internal state of Murex")]),e("li",null,[o(t,{to:"/commands/test.html"},{default:n(()=>[y]),_:1}),a(": Murex's test framework - define tests, run tests and debug shell scripts")])])])}const D=i(l,[["render",k],["__file","pipe.html.vue"]]);export{D as default};
