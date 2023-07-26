import{_ as n}from"./plugin-vue_export-helper-c27b6911.js";import{r as d,o as r,c as i,d as e,b as c,w as a,e as t,f as s}from"./app-45f7c304.js";const l={},h=s(`<h1 id="catch" tabindex="-1"><a class="header-anchor" href="#catch" aria-hidden="true">#</a> <code>catch</code></h1><blockquote><p>Handles the exception code raised by <code>try</code> or <code>trypipe</code></p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>catch</code> is designed to be used in conjunction with <code>try</code> and <code>trypipe</code> as it handles the exceptions raised by the aforementioned.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>[ try | trypipe ] { code-block } -&gt; \`&lt;stdout&gt;\`

catch { code-block } -&gt; \`&lt;stdout&gt;\`

!catch { code-block } -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>try {
    out: &quot;Hello, World!&quot; -&gt; grep: &quot;non-existent string&quot;
    out: &quot;This command will be ignored&quot;
}

catch {
    out: &quot;An error was caught&quot;
}

!catch {
    out: &quot;No errors were raised&quot;
}
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p><code>catch</code> can be used with a bang prefix to check for a lack of errors.</p><p><code>catch</code> forwards on the STDIN and exit number of the calling function.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>catch</code></li><li><code>!catch</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,14),u=e("code",null,"if",-1),p=e("code",null,"runmode",-1),f=e("code",null,"switch",-1),m=e("code",null,"try",-1),_=e("code",null,"trypipe",-1);function b(x,g){const o=d("RouterLink");return r(),i("div",null,[h,e("ul",null,[e("li",null,[c(o,{to:"/user-guide/schedulers.html"},{default:a(()=>[t("Schedulers")]),_:1}),t(": Overview of the different schedulers (or 'run modes') in Murex")]),e("li",null,[c(o,{to:"/commands/if.html"},{default:a(()=>[u]),_:1}),t(": Conditional statement to execute different blocks of code depending on the result of the condition")]),e("li",null,[c(o,{to:"/commands/runmode.html"},{default:a(()=>[p]),_:1}),t(": Alter the scheduler's behaviour at higher scoping level")]),e("li",null,[c(o,{to:"/commands/switch.html"},{default:a(()=>[f]),_:1}),t(": Blocks of cascading conditionals")]),e("li",null,[c(o,{to:"/commands/try.html"},{default:a(()=>[m]),_:1}),t(": Handles errors inside a block of code")]),e("li",null,[c(o,{to:"/commands/trypipe.html"},{default:a(()=>[_]),_:1}),t(": Checks state of each function in a pipeline and exits block on error")])])])}const q=n(l,[["render",b],["__file","catch.html.vue"]]);export{q as default};
