import{_ as r}from"./plugin-vue_export-helper-c27b6911.js";import{r as s,o as i,c,d as t,b as n,w as a,e,f as d}from"./app-45f7c304.js";const l={},u=d(`<h1 id="code-block-parsing-user-guide" tabindex="-1"><a class="header-anchor" href="#code-block-parsing-user-guide" aria-hidden="true">#</a> Code Block Parsing - User Guide</h1><blockquote><p>Overview of how code blocks are parsed</p></blockquote><p>The murex parser creates ASTs ahead of interpreting each block of code. However the AST is only generated for a block at a time. Take this sample code:</p><pre><code>function example {
    # An example function
    if { $ENVVAR } then {
        out: &#39;foobar&#39;
    }
    out: &#39;Finished!&#39;
}
</code></pre><p>When that code is run <code>function</code> is executed with the parameters <code>example</code> and <code>{ ... }</code> but the contents of <code>{ ... }</code> isn&#39;t converted into ASTs until someone calls <code>example</code> elsewhere in the shell.</p><p>When <code>example</code> (the Murex function defined above) is executed the parser will then generate AST of the commands inside said function but not any blocks that are associated with those functions. eg the AST would look something like this:</p><pre><code>[
    {
        &quot;Command&quot;: &quot;if&quot;,
        &quot;Parameters&quot;: [
            &quot;{ $ENVVAR }&quot;,
            &quot;then&quot;,
            &quot;{\\n        out: &#39;foobar&#39;\\n    }&quot;
        ]
    },
    {
        &quot;Command&quot;: &quot;out&quot;,
        &quot;Parameters&quot;: [
            &quot;Finished!&quot;
        ]
    }
]
</code></pre><blockquote><p>Please note this is a mock JSON structure rather than a representation of the actual AST that would be created. Parameters are stored differently to allow infixing of variables; and there also needs to be data shared about how pipelining (eg STDOUT et al) is chained. What is being captured above is only the command name and parameters.</p></blockquote><p>So when <code>if</code> executes, the conditional (the first parameter) is then parsed and turned into ASTs and executed. Then the last parameter (the <strong>then</strong> block) is parsed and turned into ASTs, if the first conditional is true.</p><p>This sequence of parsing is defined within the <code>if</code> builtin rather than Murex&#39;s parser. That means any code blocks are parsed only when a builtin specifically requests that they are executed.</p><p>With murex, there&#39;s no distinction between text and code. It&#39;s up to commands to determine if they want to execute a parameter as code or not (eg a curly brace block might be JSON).</p><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,12),h=t("code",null,"{",-1),p=t("code",null,"}",-1);function f(m,b){const o=s("RouterLink");return i(),c("div",null,[u,t("ul",null,[t("li",null,[n(o,{to:"/user-guide/ansi.html"},{default:a(()=>[e("ANSI Constants")]),_:1}),e(": Infixed constants that return ANSI escape sequences")]),t("li",null,[n(o,{to:"/parser/curly-brace.html"},{default:a(()=>[e("Curly Brace ("),h,e(", "),p,e(") Tokens")]),_:1}),e(": Initiates or terminates a code block")]),t("li",null,[n(o,{to:"/user-guide/pipeline.html"},{default:a(()=>[e("Pipeline")]),_:1}),e(': Overview of what a "pipeline" is')]),t("li",null,[n(o,{to:"/user-guide/schedulers.html"},{default:a(()=>[e("Schedulers")]),_:1}),e(": Overview of the different schedulers (or 'run modes') in Murex")])])])}const x=r(l,[["render",f],["__file","code-block.html.vue"]]);export{x as default};
