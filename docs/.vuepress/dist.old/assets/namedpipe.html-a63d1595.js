import{_ as d}from"./plugin-vue_export-helper-c27b6911.js";import{r as i,o as r,c as s,d as e,b as o,w as n,e as a,f as p}from"./app-45f7c304.js";const l={},c=p(`<h1 id="read-named-pipe" tabindex="-1"><a class="header-anchor" href="#read-named-pipe" aria-hidden="true">#</a> <code>&lt;&gt;</code> / <code>read-named-pipe</code></h1><blockquote><p>Reads from a Murex named pipe</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Sometimes you will need to start a command line with a Murex named pipe, eg</p><pre><code>» &lt;namedpipe&gt; -&gt; match: foobar
</code></pre><blockquote><p>See the documentation on <code>pipe</code> for more details about Murex named pipes.</p></blockquote><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><p>Read from pipe</p><pre><code>&lt;namedpipe&gt; -&gt; \`&lt;stdout&gt;\`
</code></pre><p>Write to pipe</p><pre><code>\`&lt;stdin&gt;\` -&gt; &lt;namedpipe&gt;
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>The follow two examples function the same</p><pre><code>» pipe: example
» bg { &lt;example&gt; -&gt; match: 2 }
» a: &lt;example&gt; [1..3]
2
» !pipe: example
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
</code></pre><h3 id="namespaces-and-usage-in-modules-and-packages" tabindex="-1"><a class="header-anchor" href="#namespaces-and-usage-in-modules-and-packages" aria-hidden="true">#</a> Namespaces and usage in modules and packages</h3><p>Pipes created via <code>pipe</code> are created in the global namespace. This allows pipes to be used across different functions easily however it does pose a risk with name clashes where Murex named pipes are used heavily. Thus is it recommended that pipes created in modules should be prefixed with the name of its package.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>(murex named pipe)</code></li><li><code>&lt;&gt;</code></li><li><code>read-named-pipe</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,32),h=e("code",null,"<stdin>",-1),u=e("code",null,"a",-1),m=e("code",null,"bg",-1),g=e("code",null,"ja",-1),f=e("code",null,"pipe",-1),x=e("code",null,"runtime",-1);function b(w,_){const t=i("RouterLink");return r(),s("div",null,[c,e("ul",null,[e("li",null,[o(t,{to:"/commands/stdin.html"},{default:n(()=>[h]),_:1}),a(": Read the STDIN belonging to the parent code block")]),e("li",null,[o(t,{to:"/commands/a.html"},{default:n(()=>[u,a(" (mkarray)")]),_:1}),a(": A sophisticated yet simple way to build an array or list")]),e("li",null,[o(t,{to:"/commands/bg.html"},{default:n(()=>[m]),_:1}),a(": Run processes in the background")]),e("li",null,[o(t,{to:"/commands/ja.html"},{default:n(()=>[g,a(" (mkarray)")]),_:1}),a(": A sophisticated yet simply way to build a JSON array")]),e("li",null,[o(t,{to:"/commands/pipe.html"},{default:n(()=>[f]),_:1}),a(": Manage Murex named pipes")]),e("li",null,[o(t,{to:"/commands/runtime.html"},{default:n(()=>[x]),_:1}),a(": Returns runtime information on the internal state of Murex")])])])}const q=d(l,[["render",b],["__file","namedpipe.html.vue"]]);export{q as default};
