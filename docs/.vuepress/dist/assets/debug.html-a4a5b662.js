import{_ as a}from"./plugin-vue_export-helper-c27b6911.js";import{r as d,o as r,c as s,d as e,b as o,w as n,e as u,f as i}from"./app-45f7c304.js";const q={},l=i(`<h1 id="debug" tabindex="-1"><a class="header-anchor" href="#debug" aria-hidden="true">#</a> <code>debug</code></h1><blockquote><p>Debugging information</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>debug</code> has two modes: as a function and as a method.</p><h3 id="debug-method" tabindex="-1"><a class="header-anchor" href="#debug-method" aria-hidden="true">#</a> Debug Method</h3><p>This usage will return debug information about the previous function ran.</p><h3 id="debug-function" tabindex="-1"><a class="header-anchor" href="#debug-function" aria-hidden="true">#</a> Debug Function:</h3><p>This will enable or disable debugging mode.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>\`&lt;stdin&gt;\` -&gt; debug -&gt; \`&lt;stdout&gt;\`

debug: boolean -&gt; \`&lt;stdout&gt;\`

debug -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Return debugging information on the previous function:</p><pre><code>» echo: &quot;hello, world!&quot; -&gt; debug
{
    &quot;DataType&quot;: {
        &quot;Go&quot;: &quot;[]string&quot;,
        &quot;Murex&quot;: &quot;str&quot;
    },
    &quot;Process&quot;: {
        &quot;Context&quot;: {
            &quot;Context&quot;: 0
        },
        &quot;Stdin&quot;: {},
        &quot;Stdout&quot;: {},
        &quot;Stderr&quot;: {},
        &quot;Parameters&quot;: {
            &quot;Params&quot;: [
                &quot;hello, world!&quot;
            ],
            &quot;Tokens&quot;: [
                [
                    {
                        &quot;Type&quot;: 0,
                        &quot;Key&quot;: &quot;&quot;
                    }
                ],
                [
                    {
                        &quot;Type&quot;: 1,
                        &quot;Key&quot;: &quot;hello, world!&quot;
                    }
                ],
                [
                    {
                        &quot;Type&quot;: 0,
                        &quot;Key&quot;: &quot;&quot;
                    }
                ]
            ]
        },
        &quot;ExitNum&quot;: 0,
        &quot;Name&quot;: &quot;echo&quot;,
        &quot;Id&quot;: 3750,
        &quot;Exec&quot;: {
            &quot;Pid&quot;: 0,
            &quot;Cmd&quot;: null,
            &quot;PipeR&quot;: null,
            &quot;PipeW&quot;: null
        },
        &quot;PromptGoProc&quot;: 1,
        &quot;Path&quot;: &quot;&quot;,
        &quot;IsMethod&quot;: false,
        &quot;IsNot&quot;: false,
        &quot;NamedPipeOut&quot;: &quot;out&quot;,
        &quot;NamedPipeErr&quot;: &quot;err&quot;,
        &quot;NamedPipeTest&quot;: &quot;&quot;,
        &quot;State&quot;: 7,
        &quot;IsBackground&quot;: false,
        &quot;LineNumber&quot;: 1,
        &quot;ColNumber&quot;: 1,
        &quot;RunMode&quot;: 0,
        &quot;Config&quot;: {},
        &quot;Tests&quot;: {
            &quot;Results&quot;: null
        },
        &quot;Variables&quot;: {},
        &quot;FidTree&quot;: [
            0,
            3750
        ],
        &quot;CreationTime&quot;: &quot;2019-01-20T00:00:52.167127131Z&quot;,
        &quot;StartTime&quot;: &quot;2019-01-20T00:00:52.167776212Z&quot;
    }
}
</code></pre><p>Enable or disable debug mode:</p><pre><code>» debug: on
true

» debug: off
false
</code></pre><p>Output whether debug mode is enabled or disabled:</p><pre><code>» debug
false
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p>When enabling or disabling debug mode, because the parameter is a murex boolean type, it means you can use other boolean terms. eg</p><pre><code># enable debugging
» debug 1
» debug on
» debug yes
» debug true

# disable debugging
» debug 0
» debug off
» debug no
» debug false
</code></pre><p>It is also worth noting that the debugging information needs to be written into the Go source code rather than in Murex&#39;s shell scripting language. If you require debugging those processes then please use Murex&#39;s <code>test</code> framework</p><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,22),h=e("code",null,"runtime",-1),c=e("code",null,"test",-1);function g(b,p){const t=d("RouterLink");return r(),s("div",null,[l,e("ul",null,[e("li",null,[o(t,{to:"/commands/runtime.html"},{default:n(()=>[h]),_:1}),u(": Returns runtime information on the internal state of Murex")]),e("li",null,[o(t,{to:"/commands/test.html"},{default:n(()=>[c]),_:1}),u(": Murex's test framework - define tests, run tests and debug shell scripts")])])])}const x=a(q,[["render",g],["__file","debug.html.vue"]]);export{x as default};
