import{_ as n}from"./plugin-vue_export-helper-c27b6911.js";import{r,o as d,c as s,d as e,b as t,w as l,e as o,f as c}from"./app-45f7c304.js";const i={},u=c(`<h1 id="formap" tabindex="-1"><a class="header-anchor" href="#formap" aria-hidden="true">#</a> <code>formap</code></h1><blockquote><p>Iterate through a map or other collection of data</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>formap</code> is a generic tool for iterating through a map, table or other sequences of data similarly like a <code>foreach</code>. In fact <code>formap</code> can even be used on array too.</p><p>Unlike <code>foreach</code>, <code>formap</code>&#39;s default output is <code>str</code>, so each new line will be treated as a list item. This behaviour will differ if any additional flags are used with <code>foreach</code>, such as <code>--jmap</code>.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><p><code>formap</code> writes a list:</p><pre><code>\`&lt;stdin&gt;\` -&gt; foreach variable { code-block } -&gt; \`&lt;stdout&gt;\`
</code></pre><p><code>formap</code> writes to a buffered JSON map:</p><pre><code>\`&lt;stdin&gt;\` -&gt; formap --jmap key value { code-block (map key) } { code-block (map value) } -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>First of all lets assume the following dataset:</p><pre><code>set json people={
    &quot;Tom&quot;: {
        &quot;Age&quot;: 32,
        &quot;Gender&quot;: &quot;Male&quot;
    },
    &quot;Dick&quot;: {
        &quot;Age&quot;: 43,
        &quot;Gender&quot;: &quot;Male&quot;
    },
    &quot;Sally&quot;: {
        &quot;Age&quot;: 54,
        &quot;Gender&quot;: &quot;Female&quot;
    }
}
</code></pre><p>We can create human output from this:</p><pre><code>» $people -&gt; formap key value { out &quot;$key is $value[Age] years old&quot; }
Sally is 54 years old
Tom is 32 years old
Dick is 43 years old
</code></pre><blockquote><p>Please note that maps are intentionally unsorted so you cannot guarantee the order of the output produced even if the input has been superficially set in a specific order.</p></blockquote><p>With <code>--jmap</code> we can turn that structure into a new structure:</p><pre><code>» $people -&gt; formap --jmap key value { $key } { $value[Age] }
{
    &quot;Dick&quot;: &quot;43&quot;,
    &quot;Sally&quot;: &quot;54&quot;,
    &quot;Tom&quot;: &quot;32&quot;
}
</code></pre><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2><ul><li><code>--jmap</code> Write a <code>json</code> map to STDOUT instead of an array</li></ul><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p><code>formap</code> can also work against arrays and tables as well. However <code>foreach</code> is a much better tool for ordered lists and tables can look a little funky when when there are more than 2 columns. In those instances you&#39;re better off using <code>[</code> (index) to specify columns and then <code>tabulate</code> for any data transformation.</p><h3 id="meta-values" tabindex="-1"><a class="header-anchor" href="#meta-values" aria-hidden="true">#</a> Meta values</h3><p>Meta values are a JSON object stored as the variable <code>$.</code>. The meta variable will get overwritten by any other block which invokes meta values. So if you wish to persist meta values across blocks you will need to reassign <code>$.</code>, eg</p><pre><code>%[1..3] -&gt; foreach {
    meta_parent = $.
    %[7..9] -&gt; foreach {
        out &quot;$(meta_parent.i): $.i&quot;
    }
}
</code></pre><p>The following meta values are defined:</p><ul><li><code>i</code>: iteration number</li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,28),h=e("code",null,"[",-1),p=e("code",null,"break",-1),m=e("code",null,"for",-1),f=e("code",null,"foreach",-1),q=e("code",null,"json",-1),b=e("code",null,"set",-1),_=e("code",null,"tabulate",-1),g=e("code",null,"while",-1);function y(k,v){const a=r("RouterLink");return d(),s("div",null,[u,e("ul",null,[e("li",null,[t(a,{to:"/commands/index2.html"},{default:l(()=>[h,o(" (index)")]),_:1}),o(": Outputs an element from an array, map or table")]),e("li",null,[t(a,{to:"/commands/break.html"},{default:l(()=>[p]),_:1}),o(": Terminate execution of a block within your processes scope")]),e("li",null,[t(a,{to:"/commands/for.html"},{default:l(()=>[m]),_:1}),o(": A more familiar iteration loop to existing developers")]),e("li",null,[t(a,{to:"/commands/foreach.html"},{default:l(()=>[f]),_:1}),o(": Iterate through an array")]),e("li",null,[t(a,{to:"/types/json.html"},{default:l(()=>[q]),_:1}),o(": JavaScript Object Notation (JSON)")]),e("li",null,[t(a,{to:"/commands/set.html"},{default:l(()=>[b]),_:1}),o(": Define a local variable and set it's value")]),e("li",null,[t(a,{to:"/commands/tabulate.html"},{default:l(()=>[_]),_:1}),o(": Table transformation tools")]),e("li",null,[t(a,{to:"/commands/while.html"},{default:l(()=>[g]),_:1}),o(": Loop until condition false")])])])}const j=n(i,[["render",y],["__file","formap.html.vue"]]);export{j as default};
