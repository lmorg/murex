import{_ as r}from"./plugin-vue_export-helper-c27b6911.js";import{r as i,o as d,c as l,d as t,b as a,w as n,e,f as p}from"./app-45f7c304.js";const c={},s=p(`<h1 id="generic-pipe-token-parser-reference" tabindex="-1"><a class="header-anchor" href="#generic-pipe-token-parser-reference" aria-hidden="true">#</a> Generic Pipe (<code>=&gt;</code>) Token - Parser Reference</h1><blockquote><p>Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>This token behaves much like the <code>|</code> pipe would except it injects <code>format generic</code> into the pipeline. The purpose of a formatted pipe is to support piping out to external commands which don&#39;t support Murex data types. For example they might expect arrays as lists rather than JSON objects).</p><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» ja: [Mon..Wed] =&gt; cat
Mon
Tue
Wed
</code></pre><p>The above is literally the same as typing:</p><pre><code>» ja: [Mon..Wed] -&gt; format generic -&gt; cat
Mon
Tue
Wed
</code></pre><p>To demonstrate how the previous pipeline might look without a formatted pipe:</p><pre><code>» ja: [Mon..Wed] -&gt; cat
[&quot;Mon&quot;,&quot;Tue&quot;,&quot;Wed&quot;]

» ja: [Mon..Wed] | cat
[&quot;Mon&quot;,&quot;Tue&quot;,&quot;Wed&quot;]

» ja: [Mon..Wed]
[
    &quot;Mon&quot;,
    &quot;Tue&quot;,
    &quot;Wed&quot;
]
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,11),h=t("code",null,"->",-1),u=t("code",null,"|",-1),m=t("code",null,"?",-1),f=t("code",null,"<>",-1),_=t("code",null,"read-named-pipe",-1),T=t("code",null,"format",-1),q=t("code",null,"ja",-1);function g(x,k){const o=i("RouterLink");return d(),l("div",null,[s,t("ul",null,[t("li",null,[a(o,{to:"/parser/pipe-arrow.html"},{default:n(()=>[e("Arrow Pipe ("),h,e(") Token")]),_:1}),e(": Pipes STDOUT from the left hand command to STDIN of the right hand command")]),t("li",null,[a(o,{to:"/parser/pipe-posix.html"},{default:n(()=>[e("POSIX Pipe ("),u,e(") Token")]),_:1}),e(": Pipes STDOUT from the left hand command to STDIN of the right hand command")]),t("li",null,[a(o,{to:"/user-guide/pipeline.html"},{default:n(()=>[e("Pipeline")]),_:1}),e(': Overview of what a "pipeline" is')]),t("li",null,[a(o,{to:"/parser/pipe-err.html"},{default:n(()=>[e("STDERR Pipe ("),m,e(") Token")]),_:1}),e(": Pipes STDERR from the left hand command to STDIN of the right hand command")]),t("li",null,[a(o,{to:"/commands/namedpipe.html"},{default:n(()=>[f,e(" / "),_]),_:1}),e(": Reads from a Murex named pipe")]),t("li",null,[a(o,{to:"/commands/format.html"},{default:n(()=>[T]),_:1}),e(": Reformat one data-type into another data-type")]),t("li",null,[a(o,{to:"/commands/ja.html"},{default:n(()=>[q,e(" (mkarray)")]),_:1}),e(": A sophisticated yet simply way to build a JSON array")])])])}const M=r(c,[["render",g],["__file","pipe-generic.html.vue"]]);export{M as default};
