import{_ as n}from"./plugin-vue_export-helper-c27b6911.js";import{r as u,o as d,c as s,d as e,b as a,w as r,e as t,f as l}from"./app-45f7c304.js";const c={},i=l(`<h1 id="alter" tabindex="-1"><a class="header-anchor" href="#alter" aria-hidden="true">#</a> <code>alter</code></h1><blockquote><p>Change a value within a structured data-type and pass that change along the pipeline without altering the original source input</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>alter</code> a value within a structured data-type.</p><p>The path separater is defined by the first character in the path. For example <code>/path/to/key</code>, <code>,path,to,key</code>, <code>|path|to|key</code> and <code>#path#to#key</code> are all valid however you should remember to quote or escape any special characters (tokens) used by the shell (such as pipe, <code>|</code>, and hash, <code>#</code>).</p><p>The <em>value</em> must always be supplied as JSON however</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>\`&lt;stdin&gt;\` -&gt; alter: [ -m | --merge | -s | --sum ] /path value -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» config: -&gt; [ shell ] -&gt; [ prompt ] -&gt; alter: /Value moo
{
    &quot;Data-Type&quot;: &quot;block&quot;,
    &quot;Default&quot;: &quot;{ out &#39;murex » &#39; }&quot;,
    &quot;Description&quot;: &quot;Interactive shell prompt.&quot;,
    &quot;Value&quot;: &quot;moo&quot;
}
</code></pre><p><code>alter</code> also accepts JSON as a parameter for adding structured data:</p><pre><code>config: -&gt; [ shell ] -&gt; [ prompt ] -&gt; alter: /Example { &quot;Foo&quot;: &quot;Bar&quot; }
{
    &quot;Data-Type&quot;: &quot;block&quot;,
    &quot;Default&quot;: &quot;{ out &#39;murex » &#39; }&quot;,
    &quot;Description&quot;: &quot;Interactive shell prompt.&quot;,
    &quot;Example&quot;: {
        &quot;Foo&quot;: &quot;Bar&quot;
    },
    &quot;Value&quot;: &quot;{ out &#39;murex » &#39; }&quot;
}
</code></pre><p>However it is also data type aware so if they key you&#39;re updating holds a string (for example) then the JSON data a will be stored as a string:</p><pre><code>» config: -&gt; [ shell ] -&gt; [ prompt ] -&gt; alter: /Value { &quot;Foo&quot;: &quot;Bar&quot; }
{
    &quot;Data-Type&quot;: &quot;block&quot;,
    &quot;Default&quot;: &quot;{ out &#39;murex » &#39; }&quot;,
    &quot;Description&quot;: &quot;Interactive shell prompt.&quot;,
    &quot;Value&quot;: &quot;{ \\&quot;Foo\\&quot;: \\&quot;Bar\\&quot; }&quot;
}
</code></pre><p>Numbers will also follow the same transparent conversion treatment:</p><pre><code>» tout: json { &quot;one&quot;: 1, &quot;two&quot;: 2 } -&gt; alter: /two &quot;3&quot;
{
    &quot;one&quot;: 1,
    &quot;two&quot;: 3
}
</code></pre><blockquote><p>Please note: <code>alter</code> is not changing the value held inside <code>config</code> but instead took the STDOUT from <code>config</code>, altered a value and then passed that new complete structure through it&#39;s STDOUT.</p><p>If you require modifying a structure inside Murex config (such as http headers) then you can use <code>config alter</code>. Read the config docs for reference.</p></blockquote><h3 id="m-merge" tabindex="-1"><a class="header-anchor" href="#m-merge" aria-hidden="true">#</a> -m / --merge</h3><p>Thus far all the examples have be changing existing keys. However you can also alter a structure by appending to an array or a merging two maps together. You do this with the <code>--merge</code> (or <code>-m</code>) flag.</p><pre><code>» out: a\\nb\\nc -&gt; alter: --merge / ([ &quot;d&quot;, &quot;e&quot;, &quot;f&quot; ])
a
b
c
d
e
f
</code></pre><h3 id="s-sum" tabindex="-1"><a class="header-anchor" href="#s-sum" aria-hidden="true">#</a> -s / --sum</h3><p>This behaves similarly to <code>--merge</code> where structures are blended together. However where a map exists with two keys the same and the values are numeric, those values are added together.</p><pre><code>» tout json { &quot;a&quot;: 1, &quot;b&quot;: 2 } -&gt; alter --sum / { &quot;b&quot;: 3, &quot;c&quot;: 4 }
{
    &quot;a&quot;: 1,
    &quot;b&quot;: 5,
    &quot;c&quot;: 4
}
</code></pre><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2><ul><li><code>--merge</code> Merge data structures rather than overwrite</li><li><code>--sum</code> Sum values in a map, merge items in an array</li><li><code>-m</code> Alias for \`--merge</li><li><code>-s</code> Alias for \`--sum</li></ul><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="path" tabindex="-1"><a class="header-anchor" href="#path" aria-hidden="true">#</a> Path</h3><p>The path parameter can take any character as node separators. The separator is assigned via the first character in the path. For example</p><pre><code>config -&gt; alter: .shell.prompt.Value moo
config -&gt; alter: &gt;shell&gt;prompt&gt;Value moo
</code></pre><p>Just make sure you quote or escape any characters used as shell tokens. eg</p><pre><code>config -&gt; alter: &#39;#shell#prompt#Value&#39; moo
config -&gt; alter: &#39; shell prompt Value&#39; moo
</code></pre><h3 id="supported-data-types" tabindex="-1"><a class="header-anchor" href="#supported-data-types" aria-hidden="true">#</a> Supported data-types</h3><p>The <em>value</em> field must always be supplied as JSON however the <em>STDIN</em> struct can be any data-type supported by murex.</p><p>You can check what data-types are available via the <code>runtime</code> command:</p><pre><code>runtime --marshallers
</code></pre><p>Marshallers are enabled at compile time from the <code>builtins/data-types</code> directory.</p><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,37),h=e("code",null,"[[",-1),p=e("code",null,"[",-1),m=e("code",null,"append",-1),q=e("code",null,"cast",-1),g=e("code",null,"config",-1),f=e("code",null,"format",-1),y=e("code",null,"prepend",-1),b=e("code",null,"runtime",-1);function _(x,v){const o=u("RouterLink");return d(),s("div",null,[i,e("ul",null,[e("li",null,[a(o,{to:"/commands/element.html"},{default:r(()=>[h,t(" (element)")]),_:1}),t(": Outputs an element from a nested structure")]),e("li",null,[a(o,{to:"/commands/index2.html"},{default:r(()=>[p,t(" (index)")]),_:1}),t(": Outputs an element from an array, map or table")]),e("li",null,[a(o,{to:"/commands/append.html"},{default:r(()=>[m]),_:1}),t(": Add data to the end of an array")]),e("li",null,[a(o,{to:"/commands/cast.html"},{default:r(()=>[q]),_:1}),t(": Alters the data type of the previous function without altering it's output")]),e("li",null,[a(o,{to:"/commands/config.html"},{default:r(()=>[g]),_:1}),t(": Query or define Murex runtime settings")]),e("li",null,[a(o,{to:"/commands/format.html"},{default:r(()=>[f]),_:1}),t(": Reformat one data-type into another data-type")]),e("li",null,[a(o,{to:"/commands/prepend.html"},{default:r(()=>[y]),_:1}),t(": Add data to the start of an array")]),e("li",null,[a(o,{to:"/commands/runtime.html"},{default:r(()=>[b]),_:1}),t(": Returns runtime information on the internal state of Murex")])])])}const T=n(c,[["render",_],["__file","alter.html.vue"]]);export{T as default};
