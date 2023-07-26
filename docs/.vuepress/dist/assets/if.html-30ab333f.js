import{_ as d}from"./plugin-vue_export-helper-c27b6911.js";import{r as a,o as s,c as l,d as e,b as n,w as i,e as o,f as c}from"./app-45f7c304.js";const r={},h=c(`<h1 id="if" tabindex="-1"><a class="header-anchor" href="#if" aria-hidden="true">#</a> <code>if</code></h1><blockquote><p>Conditional statement to execute different blocks of code depending on the result of the condition</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Conditional control flow</p><p><code>if</code> can be utilized both as a method as well as a standalone function. As a method, the conditional state is derived from the calling function (eg if the previous function succeeds then the condition is <code>true</code>).</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><h3 id="function-if" tabindex="-1"><a class="header-anchor" href="#function-if" aria-hidden="true">#</a> Function <code>if</code>:</h3><pre><code>if { code-block } then {
    # true
} else {
    # false
}
</code></pre><h3 id="method-if" tabindex="-1"><a class="header-anchor" href="#method-if" aria-hidden="true">#</a> Method <code>if</code>:</h3><pre><code>command -&gt; if {
    # true
} else {
    # false
}
</code></pre><h3 id="negative-function-if" tabindex="-1"><a class="header-anchor" href="#negative-function-if" aria-hidden="true">#</a> Negative Function <code>if</code>:</h3><pre><code>!if { code-block } then {
    # false
}
</code></pre><h3 id="negative-method-if" tabindex="-1"><a class="header-anchor" href="#negative-method-if" aria-hidden="true">#</a> Negative Method <code>if</code>:</h3><pre><code>command -&gt; !if {
    # false
}
</code></pre><h3 id="please-note" tabindex="-1"><a class="header-anchor" href="#please-note" aria-hidden="true">#</a> Please Note:</h3><p>the <code>then</code> and <code>else</code> statements are optional. So the first usage could also be written as:</p><pre><code>if { code-block } {
    # true
} {
    # false
}
</code></pre><p>However the practice of omitting those statements isn&#39;t recommended beyond writing short one liners in the interactive command prompt.</p><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Check if a file exists:</p><pre><code>if { g somefile.txt } then {
    out &quot;File exists&quot;
}
</code></pre><p>...or does not exist (both ways are valid):</p><pre><code>!if { g somefile.txt } then {
    out &quot;File does not exist&quot;
}

if { g somefile.txt } else {
    out &quot;File does not exist&quot;
}
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p>The conditional block can contain entire pipelines - even multiple lines of code let alone a single pipeline - as well as solitary commands as demonstrated in the examples above. However the conditional block does not output STDOUT nor STDERR to the rest of the pipeline so you don&#39;t have to worry about redirecting the output streams to <code>null</code>.</p><p>If you require output from the conditional blocks STDOUT then you will need to use either a Murex named pipe to redirect the output, or test or debug flags (depending on your use case) if you only need to occasionally inspect the conditionals output.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>if</code></li><li><code>!if</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,29),u=e("code",null,"!",-1),f=e("code",null,"and",-1),p=e("code",null,"true",-1),m=e("code",null,"false",-1),_=e("code",null,"catch",-1),b=e("code",null,"try",-1),x=e("code",null,"trypipe",-1),g=e("code",null,"debug",-1),y=e("code",null,"false",-1),v=e("code",null,"false",-1),k=e("code",null,"or",-1),w=e("code",null,"true",-1),q=e("code",null,"false",-1),R=e("code",null,"switch",-1),N=e("code",null,"test",-1),S=e("code",null,"true",-1),T=e("code",null,"true",-1),D=e("code",null,"try",-1),C=e("code",null,"trypipe",-1);function F(B,H){const t=a("RouterLink");return s(),l("div",null,[h,e("ul",null,[e("li",null,[n(t,{to:"/commands/not.html"},{default:i(()=>[u,o(" (not)")]),_:1}),o(": Reads the STDIN and exit number from previous process and not's it's condition")]),e("li",null,[n(t,{to:"/commands/and.html"},{default:i(()=>[f]),_:1}),o(": Returns "),p,o(" or "),m,o(" depending on whether multiple conditions are met")]),e("li",null,[n(t,{to:"/commands/catch.html"},{default:i(()=>[_]),_:1}),o(": Handles the exception code raised by "),b,o(" or "),x]),e("li",null,[n(t,{to:"/commands/debug.html"},{default:i(()=>[g]),_:1}),o(": Debugging information")]),e("li",null,[n(t,{to:"/commands/false.html"},{default:i(()=>[y]),_:1}),o(": Returns a "),v,o(" value")]),e("li",null,[n(t,{to:"/commands/or.html"},{default:i(()=>[k]),_:1}),o(": Returns "),w,o(" or "),q,o(" depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.")]),e("li",null,[n(t,{to:"/commands/switch.html"},{default:i(()=>[R]),_:1}),o(": Blocks of cascading conditionals")]),e("li",null,[n(t,{to:"/commands/test.html"},{default:i(()=>[N]),_:1}),o(": Murex's test framework - define tests, run tests and debug shell scripts")]),e("li",null,[n(t,{to:"/commands/true.html"},{default:i(()=>[S]),_:1}),o(": Returns a "),T,o(" value")]),e("li",null,[n(t,{to:"/commands/try.html"},{default:i(()=>[D]),_:1}),o(": Handles errors inside a block of code")]),e("li",null,[n(t,{to:"/commands/trypipe.html"},{default:i(()=>[C]),_:1}),o(": Checks state of each function in a pipeline and exits block on error")])])])}const E=d(r,[["render",F],["__file","if.html.vue"]]);export{E as default};
