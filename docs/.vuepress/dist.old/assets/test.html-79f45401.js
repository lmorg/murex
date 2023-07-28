import{_ as s}from"./plugin-vue_export-helper-c27b6911.js";import{r,o as d,c as i,d as e,b as a,w as n,e as t,f as c}from"./app-45f7c304.js";const l={},p=c(`<h1 id="test" tabindex="-1"><a class="header-anchor" href="#test" aria-hidden="true">#</a> <code>test</code></h1><blockquote><p>Murex&#39;s test framework - define tests, run tests and debug shell scripts</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>test</code> is used to define tests, run tests and debug Murex shell scripts.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><p>Define an inlined test</p><pre><code>test: define test-name { json-properties }
</code></pre><p>Define a state report</p><pre><code>test: state name { code block }
</code></pre><p>Define a unit test</p><pre><code>test: unit function|private|open|event test-name { json-properties }
</code></pre><p>Enable or disable boolean test states (more options available in <code>config</code>)</p><pre><code>test: config [ enable|!enable ] [ verbose|!verbose ] [ auto-report|!auto-report ]
</code></pre><p>Disable test mode</p><pre><code>!test
</code></pre><p>Execute a function with testing enabled</p><pre><code>test: run { code-block }
</code></pre><p>Execute unit test(s)</p><pre><code>test: run package/module/test-name|*
</code></pre><p>Write report</p><pre><code>test: report
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Inlined test</p><pre><code>function: hello-world {
    test: define example {
        &quot;StdoutRegex&quot;: (^Hello World$)
    }

    out: &lt;test_example&gt; &quot;Hello Earth&quot;
}

test: run { hello-world }
</code></pre><p>Unit test</p><pre><code>test: unit function aliases {
    &quot;PreBlock&quot;: ({
        alias ALIAS_UNIT_TEST=example param1 param2 param3
    }),
    &quot;StdoutRegex&quot;: &quot;([- _0-9a-zA-Z]+ =&gt; .*?\\n)+&quot;,
    &quot;StdoutType&quot;: &quot;str&quot;,
    &quot;PostBlock&quot;: ({
        !alias ALIAS_UNIT_TEST
    })
}

function: aliases {
    # Output the aliases in human readable format
    runtime: --aliases -&gt; formap: name alias {
        $name -&gt; sprintf: &quot;%10s =&gt; \${esccli @alias}\\n&quot;
    } -&gt; cast: str
}

test: run aliases
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="report" tabindex="-1"><a class="header-anchor" href="#report" aria-hidden="true">#</a> Report</h3><p><code>test: report</code> is only needed if <code>config: test auto-report</code> is set false. However <code>test: run</code> automatically enables <strong>auto-report</strong>.</p><p>When the report is generated, be it automatically or manually triggered, it flushes the table of pending reports.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>test</code></li><li><code>!test</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,33),u=e("code",null,"<>",-1),h=e("code",null,"read-named-pipe",-1),m=e("code",null,"config",-1);function f(_,b){const o=r("RouterLink");return d(),i("div",null,[p,e("ul",null,[e("li",null,[a(o,{to:"/commands/namedpipe.html"},{default:n(()=>[u,t(" / "),h]),_:1}),t(": Reads from a Murex named pipe")]),e("li",null,[a(o,{to:"/commands/config.html"},{default:n(()=>[m]),_:1}),t(": Query or define Murex runtime settings")])])])}const q=s(l,[["render",f],["__file","test.html.vue"]]);export{q as default};
