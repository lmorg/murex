import{_ as s}from"./plugin-vue_export-helper-c27b6911.js";import{r as i,o as r,c as d,d as e,b as o,w as n,e as t,f as c}from"./app-45f7c304.js";const h={},l=c(`<h1 id="bexists" tabindex="-1"><a class="header-anchor" href="#bexists" aria-hidden="true">#</a> <code>bexists</code></h1><blockquote><p>Check which builtins exist</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>bexists</code> takes an array of parameters and returns which commands have been compiled into Murex. The &#39;b&#39; in <code>bexists</code> stands for &#39;builtins&#39;</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>bexists command... -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» bexists: qr gzip runtime config
{
    &quot;Installed&quot;: [
        &quot;runtime&quot;,
        &quot;config&quot;
    ],
    &quot;Missing&quot;: [
        &quot;qr&quot;,
        &quot;gzip&quot;
    ]
}
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p>This builtin dates back to the start of Murex when all of the builtins were considered optional. This was intended to be a way for scripts to determine which builtins were compiled. Since then <code>runtime</code> has absorbed and centralized a number of similar commands which have since been deprecated. The same fate might also happen to <code>bexists</code> however it is in use by a few modules and for that reason alone it has been spared from the axe.</p><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,11),u=e("code",null,"fexec",-1),m=e("code",null,"runtime",-1);function p(f,b){const a=i("RouterLink");return r(),d("div",null,[l,e("ul",null,[e("li",null,[o(a,{to:"/user-guide/modules.html"},{default:n(()=>[t("Modules and Packages")]),_:1}),t(": An introduction to Murex modules and packages")]),e("li",null,[o(a,{to:"/commands/fexec.html"},{default:n(()=>[u]),_:1}),t(": Execute a command or function, bypassing the usual order of precedence.")]),e("li",null,[o(a,{to:"/commands/runtime.html"},{default:n(()=>[m]),_:1}),t(": Returns runtime information on the internal state of Murex")])])])}const q=s(h,[["render",p],["__file","bexists.html.vue"]]);export{q as default};
