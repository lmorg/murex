import{_ as a}from"./plugin-vue_export-helper-c27b6911.js";import{r as l,o as i,c as s,d as e,b as n,w as u,e as t,f as d}from"./app-45f7c304.js";const r={},c=d(`<h1 id="runtime" tabindex="-1"><a class="header-anchor" href="#runtime" aria-hidden="true">#</a> <code>runtime</code></h1><blockquote><p>Returns runtime information on the internal state of Murex</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>runtime</code> is a tool for querying the internal state of Murex. It&#39;s output will be JSON dumps.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>runtime: flags -&gt; \`&lt;stdout&gt;\`
</code></pre><p><code>builtins</code> is an alias for <code>runtime: --builtins</code>:</p><pre><code>builtins -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>List all the builtin data-types that support WriteArray()</p><pre><code>» runtime: --writearray
[
    &quot;*&quot;,
    &quot;commonlog&quot;,
    &quot;csexp&quot;,
    &quot;hcl&quot;,
    &quot;json&quot;,
    &quot;jsonl&quot;,
    &quot;qs&quot;,
    &quot;sexp&quot;,
    &quot;str&quot;,
    &quot;toml&quot;,
    &quot;yaml&quot;
]
</code></pre><p>List all the functions</p><pre><code>» runtime: --functions -&gt; [ agent aliases ]
[
    {
        &quot;Block&quot;: &quot;\\n    # Launch ssh-agent\\n    ssh-agent -\\u003e head -n2 -\\u003e [ :0 ] -\\u003e prefix \\&quot;export \\&quot; -\\u003e source\\n    ssh-add: @{g \\u003c!null\\u003e ~/.ssh/*.key} @{g \\u003c!null\\u003e ~/.ssh/*.pem}\\n&quot;,
        &quot;FileRef&quot;: {
            &quot;Column&quot;: 1,
            &quot;Line&quot;: 149,
            &quot;Source&quot;: {
                &quot;DateTime&quot;: &quot;2019-07-07T14:06:11.05581+01:00&quot;,
                &quot;Filename&quot;: &quot;/home/lau/.murex_profile&quot;,
                &quot;Module&quot;: &quot;profile/.murex_profile&quot;
            }
        },
        &quot;Summary&quot;: &quot;Launch ssh-agent&quot;
    },
    {
        &quot;Block&quot;: &quot;\\n\\t# Output the aliases in human readable format\\n\\truntime: --aliases -\\u003e formap name alias {\\n        $name -\\u003e sprintf: \\&quot;%10s =\\u003e \${esccli @alias}\\\\n\\&quot;\\n\\t} -\\u003e cast str\\n&quot;,
        &quot;FileRef&quot;: {
            &quot;Column&quot;: 1,
            &quot;Line&quot;: 6,
            &quot;Source&quot;: {
                &quot;DateTime&quot;: &quot;2019-07-07T14:06:10.886706796+01:00&quot;,
                &quot;Filename&quot;: &quot;(builtin)&quot;,
                &quot;Module&quot;: &quot;source/builtin&quot;
            }
        },
        &quot;Summary&quot;: &quot;Output the aliases in human readable format&quot;
    }
]
</code></pre><p>To get a list of every flag supported by <code>runtime</code></p><pre><code>» runtime: --help
[
    &quot;--aliases&quot;,
    &quot;--astcache&quot;,
    &quot;--config&quot;,
    &quot;--debug&quot;,
    &quot;--events&quot;,
    &quot;--fids&quot;,
    &quot;--flags&quot;,
    &quot;--functions&quot;,
    &quot;--help&quot;,
    &quot;--indexes&quot;,
    &quot;--marshallers&quot;,
    &quot;--memstats&quot;,
    &quot;--modules&quot;,
    &quot;--named-pipes&quot;,
    &quot;--open-agents&quot;,
    &quot;--pipes&quot;,
    &quot;--privates&quot;,
    &quot;--readarray&quot;,
    &quot;--readmap&quot;,
    &quot;--sources&quot;,
    &quot;--test-results&quot;,
    &quot;--tests&quot;,
    &quot;--unmarshallers&quot;,
    &quot;--variables&quot;,
    &quot;--writearray&quot;
]
</code></pre><p>Please also note that you can supply more than one flag. However when you do use multiple flags the top level of the JSON output will be a map of the flag names. eg</p><pre><code>» runtime: --pipes --tests
{
    &quot;pipes&quot;: [
        &quot;file&quot;,
        &quot;std&quot;,
        &quot;tcp-dial&quot;,
        &quot;tcp-listen&quot;,
        &quot;udp-dial&quot;,
        &quot;udp-listen&quot;
    ],
    &quot;tests&quot;: {
        &quot;state&quot;: {},
        &quot;test&quot;: []
    }
}

» runtime: --pipes
[
    &quot;file&quot;,
    &quot;std&quot;,
    &quot;tcp-dial&quot;,
    &quot;tcp-listen&quot;,
    &quot;udp-dial&quot;,
    &quot;udp-listen&quot;
]

» runtime: --tests
{
    &quot;state&quot;: {},
    &quot;test&quot;: []
}
</code></pre><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2><ul><li><code>--aliases</code> Lists all aliases</li><li><code>--astcache</code> Lists some data about cached ASTs</li><li><code>--autocomplete</code> Lists all <code>autocomplete</code> schemas - both user defined and automatically generated one</li><li><code>--builtins</code> Lists all builtin commands, compiled into Murex</li><li><code>--config</code> Lists all properties available to \`config</li><li><code>--debug</code> Outputs the state of debug and inspect mode</li><li><code>--events</code> Lists all builtin event types and any defined events</li><li><code>--exports</code> Outputs environmental variables. For Murex variables (<code>global</code> and <code>set</code>/<code>let</code>) use \`--variables</li><li><code>--fids</code> Lists all running processes / functions</li><li><code>--functions</code> Lists all Murex global functions</li><li><code>--globals</code> Lists all global variables</li><li><code>--help</code> Outputs a list of <code>runtimes</code>&#39;s flags</li><li><code>--indexes</code> Lists all builtin data-types which are supported by index (<code>[</code>)</li><li><code>--marshallers</code> Lists all builtin data-types with marshallers (eg required for <code>format</code>)</li><li><code>--memstats</code> Outputs the running state of Go&#39;s runtime</li><li><code>--methods</code> Lists all commands with a defined STDOUT and STDIN data type. This is used to generate smarter autocompletion suggestions with \`-&gt;</li><li><code>--modules</code> Lists all installed modules</li><li><code>--named-pipes</code> Lists all named pipes defined</li><li><code>--not-indexes</code> Lists all builtin data-types which are supported by index (<code>![</code>)</li><li><code>--open-agents</code> Lists all registered <code>open</code> handlers</li><li><code>--pipes</code> Lists builtin pipes compiled into Murex. These can be then be defined as named-pipes</li><li><code>--privates</code> Lists all Murex private functions</li><li><code>--readarray</code> Lists all builtin data-types which support ReadArray()</li><li><code>--readarraywithtype</code> Lists all builtin data-types which support ReadArrayWithType()</li><li><code>--readmap</code> Lists all builtin data-types which support ReadMap()</li><li><code>--sources</code> Lists all loaded murex sources</li><li><code>--summaries</code> Outputs all the override summaries</li><li><code>--test-results</code> A dump of any unreported test results</li><li><code>--tests</code> Lists defined tests</li><li><code>--unmarshallers</code> Lists all builtin data-types with unmarshallers (eg required for <code>format</code>)</li><li><code>--variables</code> Lists all local Murex variables which doesn&#39;t include environmental nor global variables</li><li><code>--writearray</code> Lists all builtin data-types which support WriteArray()</li></ul><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="usage-in-scripts" tabindex="-1"><a class="header-anchor" href="#usage-in-scripts" aria-hidden="true">#</a> Usage in scripts</h3><p><code>runtime</code> should not be used in scripts because the output of <code>runtime</code> may be subject to change as and when the internal mechanics of Murex change. The purpose behind <code>runtime</code> is not to provide an API but rather to provide a verbose &quot;dump&quot; of the internal running state of Murex.</p><p>If you require a stable API to script against then please use the respective command line tool. For example <code>fid-list</code> instead of <code>runtime --fids</code>. Some tools will provide a human readable output when STDOUT is a TTY but output a script parsable version when STDOUT is not a terminal.</p><pre><code>» fid-list
    FID   Parent    Scope  State         Run Mode  BG   Out Pipe    Err Pipe    Command     Parameters
      0        0        0  Executing     Shell     no                           -murex
 265499        0        0  Executing     Normal    no   out         err         fid-list

» fid-list -&gt; pretty
[
    {
        &quot;FID&quot;: 0,
        &quot;Parent&quot;: 0,
        &quot;Scope&quot;: 0,
        &quot;State&quot;: &quot;Executing&quot;,
        &quot;Run Mode&quot;: &quot;Shell&quot;,
        &quot;BG&quot;: false,
        &quot;Out Pipe&quot;: &quot;&quot;,
        &quot;Err Pipe&quot;: &quot;&quot;,
        &quot;Command&quot;: &quot;-murex&quot;,
        &quot;Parameters&quot;: &quot;&quot;
    },
    {
        &quot;FID&quot;: 265540,
        &quot;Parent&quot;: 0,
        &quot;Scope&quot;: 0,
        &quot;State&quot;: &quot;Executing&quot;,
        &quot;Run Mode&quot;: &quot;Normal&quot;,
        &quot;BG&quot;: false,
        &quot;Out Pipe&quot;: &quot;out&quot;,
        &quot;Err Pipe&quot;: &quot;err&quot;,
        &quot;Command&quot;: &quot;fid-list&quot;,
        &quot;Parameters&quot;: &quot;&quot;
    },
    {
        &quot;FID&quot;: 265541,
        &quot;Parent&quot;: 0,
        &quot;Scope&quot;: 0,
        &quot;State&quot;: &quot;Executing&quot;,
        &quot;Run Mode&quot;: &quot;Normal&quot;,
        &quot;BG&quot;: false,
        &quot;Out Pipe&quot;: &quot;out&quot;,
        &quot;Err Pipe&quot;: &quot;err&quot;,
        &quot;Command&quot;: &quot;pretty&quot;,
        &quot;Parameters&quot;: &quot;&quot;
    }
]
</code></pre><h3 id="file-reference" tabindex="-1"><a class="header-anchor" href="#file-reference" aria-hidden="true">#</a> File reference</h3><p>Some of the JSON dumps produced from <code>runtime</code> will include a map called <code>FileRef</code>. This is a trace of the source file that defined it. It is used by Murex to help provide meaningful errors (eg with line and character positions) however it is also useful for manually debugging user-defined properties such as which module or script defined an <code>autocomplete</code> schema.</p><h3 id="debug-mode" tabindex="-1"><a class="header-anchor" href="#debug-mode" aria-hidden="true">#</a> Debug mode</h3><p>When <code>debug</code> is enabled garbage collection is disabled for variables and FIDs. This means the output of <code>runtime --variables</code> and <code>runtime --fids</code> will contain more than just the currently defined variables and running functions.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>runtime</code></li><li><code>builtins</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,31),q=e("code",null,"[",-1),m=e("code",null,"autocomplete",-1),h=e("code",null,"config",-1),p=e("code",null,"debug",-1),f=e("code",null,"event",-1),b=e("code",null,"export",-1),g=e("code",null,"fid-list",-1),_=e("code",null,"foreach",-1),x=e("code",null,"formap",-1),y=e("code",null,"format",-1),v=e("code",null,"function",-1),L=e("code",null,"global",-1),w=e("code",null,"let",-1),S=e("code",null,"method",-1),M=e("code",null,"open",-1),T=e("code",null,"openagent",-1),D=e("code",null,"pipe",-1),O=e("code",null,"pretty",-1),P=e("code",null,"private",-1),k=e("code",null,"set",-1),F=e("code",null,"source",-1),I=e("code",null,"test",-1);function R(E,N){const o=l("RouterLink");return i(),s("div",null,[c,e("ul",null,[e("li",null,[n(o,{to:"/commands/index2.html"},{default:u(()=>[q,t(" (index)")]),_:1}),t(": Outputs an element from an array, map or table")]),e("li",null,[n(o,{to:"/commands/autocomplete.html"},{default:u(()=>[m]),_:1}),t(": Set definitions for tab-completion in the command line")]),e("li",null,[n(o,{to:"/commands/config.html"},{default:u(()=>[h]),_:1}),t(": Query or define Murex runtime settings")]),e("li",null,[n(o,{to:"/commands/debug.html"},{default:u(()=>[p]),_:1}),t(": Debugging information")]),e("li",null,[n(o,{to:"/commands/event.html"},{default:u(()=>[f]),_:1}),t(": Event driven programming for shell scripts")]),e("li",null,[n(o,{to:"/commands/export.html"},{default:u(()=>[b]),_:1}),t(": Define an environmental variable and set it's value")]),e("li",null,[n(o,{to:"/commands/fid-list.html"},{default:u(()=>[g]),_:1}),t(": Lists all running functions within the current Murex session")]),e("li",null,[n(o,{to:"/commands/foreach.html"},{default:u(()=>[_]),_:1}),t(": Iterate through an array")]),e("li",null,[n(o,{to:"/commands/formap.html"},{default:u(()=>[x]),_:1}),t(": Iterate through a map or other collection of data")]),e("li",null,[n(o,{to:"/commands/format.html"},{default:u(()=>[y]),_:1}),t(": Reformat one data-type into another data-type")]),e("li",null,[n(o,{to:"/commands/function.html"},{default:u(()=>[v]),_:1}),t(": Define a function block")]),e("li",null,[n(o,{to:"/commands/global.html"},{default:u(()=>[L]),_:1}),t(": Define a global variable and set it's value")]),e("li",null,[n(o,{to:"/commands/let.html"},{default:u(()=>[w]),_:1}),t(": Evaluate a mathematical function and assign to variable (deprecated)")]),e("li",null,[n(o,{to:"/commands/method.html"},{default:u(()=>[S]),_:1}),t(": Define a methods supported data-types")]),e("li",null,[n(o,{to:"/commands/open.html"},{default:u(()=>[M]),_:1}),t(": Open a file with a preferred handler")]),e("li",null,[n(o,{to:"/commands/openagent.html"},{default:u(()=>[T]),_:1}),t(": Creates a handler function for `open")]),e("li",null,[n(o,{to:"/commands/pipe.html"},{default:u(()=>[D]),_:1}),t(": Manage Murex named pipes")]),e("li",null,[n(o,{to:"/commands/pretty.html"},{default:u(()=>[O]),_:1}),t(": Prettifies JSON to make it human readable")]),e("li",null,[n(o,{to:"/commands/private.html"},{default:u(()=>[P]),_:1}),t(": Define a private function block")]),e("li",null,[n(o,{to:"/commands/set.html"},{default:u(()=>[k]),_:1}),t(": Define a local variable and set it's value")]),e("li",null,[n(o,{to:"/commands/source.html"},{default:u(()=>[F]),_:1}),t(": Import Murex code from another file of code block")]),e("li",null,[n(o,{to:"/commands/test.html"},{default:u(()=>[I]),_:1}),t(": Murex's test framework - define tests, run tests and debug shell scripts")])])])}const C=a(r,[["render",R],["__file","runtime.html.vue"]]);export{C as default};
