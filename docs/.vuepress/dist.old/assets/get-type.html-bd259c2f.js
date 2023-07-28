import{_ as l}from"./plugin-vue_export-helper-c27b6911.js";import{r as d,o as i,c as s,d as e,b as o,w as n,e as t,f as r}from"./app-45f7c304.js";const p={},c=r(`<h1 id="get-type" tabindex="-1"><a class="header-anchor" href="#get-type" aria-hidden="true">#</a> <code>get-type</code></h1><blockquote><p>Returns the data-type of a variable or pipe</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>get-type</code> returns the Murex data-type of a variable or pipe without reading the data from it.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>get-type: \\$variable -&gt; \`&lt;stdout&gt;\`

get-type: stdin -&gt; \`&lt;stdout&gt;\`

get-type: pipe -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Get the data-type of a variable</p><pre><code>» set: json example={[1,2,3]}
» get-type: \\$example
json
</code></pre><blockquote><p>Please note that you will need to escape the dollar sign. If you don&#39;t the value of the variable will be passed to <code>get-type</code> rather than the name.</p></blockquote><p>Get the data-type of a functions STDIN</p><pre><code>» function: example { get-type stdin }
» tout: json {[1,2,3]} -&gt; example
json
</code></pre><p>Get the data-type of a Murex named pipe</p><pre><code>» pipe: example
» tout: &lt;example&gt; json {[1,2,3]}
» get-type: example
» !pipe: example
json
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,15),u=e("code",null,"debug",-1),h=e("code",null,"function",-1),m=e("code",null,"pipe",-1),f=e("code",null,"runtime",-1),g=e("code",null,"set",-1),_=e("code",null,"tout",-1);function x(b,y){const a=d("RouterLink");return i(),s("div",null,[c,e("ul",null,[e("li",null,[o(a,{to:"/user-guide/reserved-vars.html"},{default:n(()=>[t("Reserved Variables")]),_:1}),t(": Special variables reserved by Murex")]),e("li",null,[o(a,{to:"/user-guide/scoping.html"},{default:n(()=>[t("Variable and Config Scoping")]),_:1}),t(": How scoping works within Murex")]),e("li",null,[o(a,{to:"/commands/debug.html"},{default:n(()=>[u]),_:1}),t(": Debugging information")]),e("li",null,[o(a,{to:"/commands/function.html"},{default:n(()=>[h]),_:1}),t(": Define a function block")]),e("li",null,[o(a,{to:"/commands/pipe.html"},{default:n(()=>[m]),_:1}),t(": Manage Murex named pipes")]),e("li",null,[o(a,{to:"/commands/runtime.html"},{default:n(()=>[f]),_:1}),t(": Returns runtime information on the internal state of Murex")]),e("li",null,[o(a,{to:"/commands/set.html"},{default:n(()=>[g]),_:1}),t(": Define a local variable and set it's value")]),e("li",null,[o(a,{to:"/commands/tout.html"},{default:n(()=>[_]),_:1}),t(": Print a string to the STDOUT and set it's data-type")])])])}const w=l(p,[["render",x],["__file","get-type.html.vue"]]);export{w as default};
